package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"auth_service/config"
	pb "auth_service/generated"
	"auth_service/models"
	"auth_service/queue"
	"auth_service/repositories"

	"auth_service/utils"

	"google.golang.org/protobuf/proto"
)

// AuthServiceServer dengan Redis Client sebagai Dependency Injection
type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	// RedisClient    *redis.Client // Tambahkan Redis sebagai field
	repoSekolah repositories.SekolahRepository
	authService AuthService
	repoProfile repositories.GenericRepository[models.UserProfile]
	repoUser    repositories.UserRepository
	rQueue      queue.RedisEnqueue
}

func NewAuthServiceServer() *AuthServiceServer {
	repoAuth := repositories.NewUserRepository(config.DB)
	authService := NewAuthService(repoAuth)
	repoSekolah := repositories.NewSekolahRepository(config.DB)
	repoProfile := repositories.NewUserProfileRepository(config.DB)
	repoUser := repositories.NewUserRepository(config.DB)
	rQueue := queue.NewRedisEnqueue(config.RDB)
	return &AuthServiceServer{
		authService: authService,
		repoSekolah: repoSekolah,
		repoProfile: *repoProfile,
		repoUser:    repoUser,
		rQueue:      *rQueue,
	}
}

type SchoolRegistration struct {
	SchoolName string `json:"school_name"`
	AdminEmail string `json:"admin_email"`
}

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// log.Printf("Received Sekolah data request: %+v\n", req)

	username := req.GetUsername()
	email := req.GetEmail()
	password := req.Password
	var resp *models.User
	var err error
	if email != "" {
		resp, err = s.authService.Login(email, password)
		if err != nil {
			return nil, err
		}
	}
	if username != "" {
		resp, err = s.authService.Login(username, password)
		if err != nil {
			log.Printf("Error username/password salah : %v", err)
			return nil, err
		}
	}
	// Generate JWT
	token, err := utils.GenerateJWT(resp)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	user := &pb.User{
		Id:              resp.ID,
		Username:        resp.Username,
		Email:           resp.Email,
		Role:            resp.Role,
		SekolahTenantId: resp.SekolahTenantID,
	}

	// Ambil data sekolah
	sekolahModel, err := s.repoSekolah.GetSekolahByTenantId(resp.SekolahTenantID)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Token: token,
		Ok:    true,
		User:  user,
		SekolahTenant: &pb.SekolahTenant{
			NamaSekolah: sekolahModel.NamaSekolah,
			Npsn:        sekolahModel.NPSN,
		},
	}, nil
}

func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Sekolah", "User"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	// Ambil data dari request
	user := req.GetUser()
	sekolah := req.GetSekolah()
	// Cek apakah sekolah sudah ada
	var sekolahModel *models.SekolahTenant
	sekolahModel, err = s.repoSekolah.GetSekolahByNPSN(sekolah.Npsn)
	if err != nil {
		if errors.Is(err, repositories.ErrRecordNotFound) {
			if user.Role == "admin" {
				// Buat sekolah baru
				sekolahModel = &models.SekolahTenant{
					NamaSekolah:   sekolah.NamaSekolah,
					NPSN:          sekolah.Npsn,
					EnkripID:      *sekolah.EnkripId,
					Kecamatan:     sekolah.Kecamatan,
					Kabupaten:     sekolah.Kabupaten,
					Propinsi:      sekolah.Propinsi,
					KodeKecamatan: sekolah.KodeKecamatan,
					KodeKab:       sekolah.KodeKab,
					KodeProp:      sekolah.KodeProp,
					AlamatJalan:   sekolah.AlamatJalan,
				}
				err = s.repoSekolah.CreateSekolah(sekolahModel)

				if err != nil {
					log.Printf("Gagal membuat sekolah: %v", err)
					return nil, fmt.Errorf("gagal membuat sekolah: %w", err)
				}

			} else {
				// Pendaftaran siswa
				return nil, fmt.Errorf("sekolah belum terdaftar di aplikasi")
			}
		} else {
			log.Printf("Server error saat mencari sekolah: %v", err)
			return nil, fmt.Errorf("server error: %w", err)
		}
	}
	// Hubungkan user dengan sekolah
	userModel := &models.User{
		Username:        utils.GenerateUsername(user.Email, sekolah.Npsn),
		Email:           user.Email,
		Role:            user.Role,
		SekolahTenantID: sekolahModel.ID,
		Password:        *user.Password,
	}

	// Cek jika role user adalah admin dan apakah sudah ada admin
	if userModel.Role == "admin" {
		adminExists, err := s.authService.IsAdminExists(sekolahModel.ID)
		if err != nil {
			log.Printf("Error mengecek admin: %v", err)
			return nil, fmt.Errorf("server error: %w", err)
		}
		if adminExists {
			return nil, fmt.Errorf("admin sudah ada untuk sekolah ini")
		}

		// Registrasi admin
		if err := s.authService.RegisterAdmin(userModel); err != nil {
			log.Printf("Error registrasi admin: %v", err)
			return nil, fmt.Errorf("gagal registrasi admin: %w", err)
		}
		// ==============Inisialisasi database di sekolah service==============
		// errChan := make(chan error, 1)
		// // 1. Buat client untuk sekolah_service
		// go func() {
		// 	errChan <- initSekolahService(sekolahModel)
		// }()
		if err := s.rQueue.EnqueueInitSekolahTask(*sekolahModel); err != nil {
			log.Printf("Gagal enqueue task: %v", err)
			return nil, fmt.Errorf("gagal enqueue initSekolahService: %w", err)
		}

		// ==============Inisialisasi database di sc service==============
		// Buat client untuk SC_service
		// go func() {
		// 	errChan <- initSCService(sekolahModel, int(userModel.ID))
		// }()

		// for i := 0; i < 1; i++ {
		// 	if err := <-errChan; err != nil {
		// 		return nil, err
		// 	}
		// }

	} else if userModel.Role == "siswa" {
		// Registrasi siswa
		if err := s.authService.Register(userModel); err != nil {
			log.Printf("Error registrasi siswa: %v", err)
			return nil, fmt.Errorf("gagal registrasi siswa: %w", err)
		}
	}

	// Hubungkan user dengan profil
	userProfileModel := &models.UserProfile{
		UserID: userModel.ID,
	}

	if err := s.repoProfile.Save(ctx, userProfileModel, "public"); err != nil {
		log.Printf("Error membuat user profile: %v", err)
		return nil, fmt.Errorf("server error saat membuat user profile")
	}

	var response *pb.RegisterResponse
	if userModel.Role == "admin" {
		token, err := utils.GenerateJWT(userModel)
		if err != nil {
			return nil, errors.New("failed to generate token")
		}
		response = &pb.RegisterResponse{
			Token: token,
			Ok:    true,
			User: &pb.User{
				Id:       userModel.ID,
				Username: userModel.Username,
				// Email:     userModel.Email,
				Role:            userModel.Role,
				SekolahTenantId: userModel.SekolahTenantID,
			},
			SekolahTenant: &pb.SekolahTenant{
				NamaSekolah: sekolahModel.NamaSekolah,
			},
		}
	} else {
		response = &pb.RegisterResponse{
			Ok: true,
		}
	}
	log.Println("User berhasil didaftarkan")
	return response, nil
}

// Digunakan untuk membuat user baru dengan role siswa
func (s *AuthServiceServer) CreateUsers(ctx context.Context, req *pb.CreateUsersRequest) (*pb.CreateUsersResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Users"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	users := req.GetUsers()
	for _, v := range users {
		username := utils.GenerateUsername(v.Sekolah.NamaSekolah, v.UserProfile.Nama)
		pass, err := utils.GeneratePassword(4)
		if err != nil {
			return nil, err
		}
		newUser := models.User{
			Username:        username,
			Email:           v.Email,
			Password:        pass,
			SekolahTenantID: v.User.SekolahTenantId,
			Role:            v.User.Role,
		}
		err = s.authService.Register(&newUser)
		if err != nil {
			return nil, err
		}
		// Hubungkan dengan user profile
		userId := newUser.ID
		userProfile := models.UserProfile{
			UserID:   userId,
			Nama:     v.UserProfile.Nama,
			JK:       v.UserProfile.Jk,
			Phone:    &v.UserProfile.Phone,
			TptLahir: &v.UserProfile.TptLahir,
			// TglLahir:  v.UserProfile.TglLahir,
			AlamatJalan: &v.UserProfile.AlamatJalan,
			KotaKab:     &v.UserProfile.KotaKab,
		}
		err = s.repoProfile.Save(ctx, &userProfile, "public")
		if err != nil {
			return nil, err
		}
	}

	return &pb.CreateUsersResponse{
		Message: "ok",
	}, nil
}
func (s *AuthServiceServer) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	return &pb.ResetPasswordResponse{
		Message: "password berhasil direset",
	}, nil
}
func (s *AuthServiceServer) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"SekolahId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}

	var roleSiswa bool
	if req.Role == "" {
		roleSiswa = true
	}

	var usersModel []models.User
	if roleSiswa {
		usersModel, err = s.repoUser.GetUsers("siswa", req.GetSekolahTenantId())
		if err != nil {
			return nil, err
		}

	} // tambahkan else jika ingin menampilkan role admin
	res := utils.ConvertModelsToPB(usersModel, func(model models.User) *pb.User {
		return &pb.User{
			Username:  model.Username,
			Password:  model.InitialPassword,
			LastLogin: proto.String(utils.TimeToString(*model.LastLogin, "2006-01=12")),
		}
	})

	return &pb.GetUsersResponse{
		Users: res,
	}, nil

}

// func (s *AuthServiceServer) GenerateStudentUsername(ctx context.Context, req *pb.GenerateStudentUsernameRequest) (*pb.GenerateStudentUsernameResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"SekolahId", "PesertaDidikId"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}

// }
// ✔ Digunakan
func (s *AuthServiceServer) GetSekolah(ctx context.Context, req *pb.GetSekolahRequest) (*pb.GetSekolahResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Npsn"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}

	sekolah, err := s.repoSekolah.GetSekolahByNPSN(req.GetNpsn())
	if err != nil {
		if errors.Is(err, repositories.ErrRecordNotFound) {
			log.Printf("Error fetching school data: %v", err)
			return nil, errors.New("failed to retrieve school data")
		}
		return nil, err
	}

	return &pb.GetSekolahResponse{
		Nama: sekolah.NamaSekolah,
		SekolahData: &pb.SekolahTenant{
			Id:            &sekolah.ID,
			EnkripId:      &sekolah.EnkripID,
			Kecamatan:     sekolah.Kecamatan,
			Kabupaten:     sekolah.Kabupaten,
			Propinsi:      sekolah.Propinsi,
			KodeKecamatan: sekolah.Kecamatan,
			AlamatJalan:   sekolah.AlamatJalan,
			KodeKab:       sekolah.KodeKab,
			KodeProp:      sekolah.KodeProp,
			NamaSekolah:   sekolah.NamaSekolah,
			Npsn:          sekolah.NPSN,
		},
	}, nil
}

//	func registerAdmin(authService AuthService, userModel UserModel) error {
//		if err := authService.RegisterAdmin(userModel); err != nil {
//			log.Printf("Error registrasi admin: %v", err)
//			return fmt.Errorf("gagal registrasi admin: %w", err)
//		}
//		return nil
//	}

func initSCService(sekolahModel *models.SekolahTenant, userID int) error {
	scClient, err := NewSCServiceClient()
	if err != nil {
		return err
	}

	if err := scClient.RegistrasiSekolahTenant(sekolahModel, int32(userID)); err != nil {
		return err
	}

	return nil
}
