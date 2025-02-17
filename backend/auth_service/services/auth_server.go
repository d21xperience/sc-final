package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"auth_service/config"
	pb "auth_service/generated"
	"auth_service/models"
	"auth_service/repositories"

	"auth_service/utils"
)

// AuthServiceServer dengan Redis Client sebagai Dependency Injection
type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	// RedisClient    *redis.Client // Tambahkan Redis sebagai field
	repoSekolah repositories.SekolahRepository
	authService AuthService
	repoProfile repositories.GenericRepository[models.UserProfile]
}

func NewAuthServiceServer() *AuthServiceServer {
	repoAuth := repositories.NewUserRepository(config.DB)
	authService := NewAuthService(repoAuth)
	repoSekolah := repositories.NewSekolahRepository(config.DB)
	repoProfile := repositories.NewUserProfileRepository(config.DB)
	return &AuthServiceServer{
		authService: authService,
		repoSekolah: repoSekolah,
		repoProfile: *repoProfile,
	}
}

type SchoolRegistration struct {
	SchoolName string `json:"school_name"`
	AdminEmail string `json:"admin_email"`
}

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	username := req.GetUsername()
	password := req.Password

	resp, err := s.authService.Login(username, password)
	if err != nil {
		log.Printf("Error username/password salah : %v", err)
		return nil, err
	}
	// Generate JWT
	token, err := utils.GenerateJWT(resp)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	return &pb.LoginResponse{
		Token: token,
		Ok:    true,
		User: &pb.User{
			UserId:    resp.ID,
			Username:  resp.Username,
			Email:     resp.Email,
			Role:      resp.Role,
			SekolahId: resp.SekolahID,
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
	var query = repositories.SekolahQuery{
		Npsn: sekolah.Npsn,
	}
	// Cek apakah sekolah sudah ada
	var sekolahModel *models.Sekolah
	sekolahModel, err = s.repoSekolah.GetSekolah(query)
	if err != nil {
		if errors.Is(err, repositories.ErrRecordNotFound) {
			if user.Role == "admin" {
				// Buat sekolah baru
				sekolahModel = &models.Sekolah{
					NPSN:            sekolah.Npsn,
					NamaSekolah:     sekolah.NamaSekolah,
					SekolahIDEnkrip: sekolah.SekolahIdEnkrip,
					Kecamatan:       sekolah.Kecamatan,
					Kabupaten:       sekolah.Kabupaten,
					Propinsi:        sekolah.Propinsi,
					KodeKecamatan:   sekolah.KodeKecamatan,
					KodeKab:         sekolah.KodeKab,
					KodeProp:        sekolah.KodeProp,
					AlamatJalan:     sekolah.AlamatJalan,
					Status:          sekolah.Status,
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
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		SekolahID: sekolahModel.ID,
		Password:  user.Password,
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
		// Buat database
		// 2. Buat client untuk sekolah_service
		sekolahClient, err := NewSekolahServiceClient()
		if err != nil {
			return nil, err
		}
		// 3. Panggil sekolah_service untuk membuat schema database sekolah
		if err := sekolahClient.RegistrasiSekolah(sekolahModel); err != nil {
			return nil, err
		}
		// 4. Panggil sekolah_service untuk membuat inisialiasi data sekolah
		if err := sekolahClient.CreateSekolah(sekolahModel); err != nil {
			return nil, err
		}

	} else if userModel.Role == "siswa" {
		// Registrasi siswa
		if err := s.authService.Register(userModel); err != nil {
			log.Printf("Error registrasi siswa: %v", err)
			return nil, fmt.Errorf("gagal registrasi siswa: %w", err)
		}
	}

	// Hubungkan user dengan profil
	userProfileModel := &models.UserProfile{
		UserId: userModel.ID,
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
				UserId:   userModel.ID,
				Username: userModel.Username,
				// Email:     userModel.Email,
				Role:      userModel.Role,
				SekolahId: userModel.SekolahID,
			},
		}
	} else {
		response = &pb.RegisterResponse{
			Ok: true,
		}
	}
	log.Println("User berhasil didaftarkan")

	// // Siapkan data registrasi sekolah untuk dikirim ke Redis
	// registration := SchoolRegistration{
	// 	SchoolName: sekolah.GetNamaSekolah(), // Menggunakan metode GetNamaSekolah()
	// 	AdminEmail: user.GetEmail(),          // Menggunakan metode GetEmail()
	// }

	// // Konversi data ke JSON
	// data, err := json.Marshal(registration)
	// if err != nil {
	// 	log.Printf("Error marshalling registration data: %v", err)
	// 	return nil, fmt.Errorf("gagal memproses data registrasi")
	// }

	// // Kirim data ke Redis
	// err = s.RedisClient.Publish("school_registration", data).Err()
	// if err != nil {
	// 	log.Printf("Error publishing to Redis: %v", err)
	// 	return nil, fmt.Errorf("gagal mengirim data ke sistem antrian")
	// }

	// fmt.Println("Registration message published successfully!")

	return response, nil
}

// func (s *AuthServiceServer) GetSekolah(ctx context.Context, req *pb.GetSekolahRequest) (*pb.GetSekolahResponse, error) {
// 	var query = repository.SekolahQuery{
// 		Npsn:      req.GetNpsn(),
// 		SekolahID: int(req.GetSekolahId()),
// 	}

// 	sekolah, err := s.sekolahService.GetSekolah(query)
// 	if err != nil {
// 		log.Printf("Error fetching school data: %v", err)
// 		return nil, errors.New("failed to retrieve school data")
// 	}

// 	return &pb.GetSekolahResponse{
// 		Nama: sekolah.NamaSekolah,
// 		SekolahData: &pb.Sekolah{
// 			SekolahId:       int32(sekolah.ID),
// 			SekolahIdEnkrip: sekolah.SekolahIDEnkrip,
// 			Kecamatan:       sekolah.Kecamatan,
// 			Kabupaten:       sekolah.Kabupaten,
// 			Propinsi:        sekolah.Propinsi,
// 			KodeKecamatan:   sekolah.Kecamatan,
// 			AlamatJalan:     sekolah.AlamatJalan,
// 			KodeKab:         sekolah.KodeKab,
// 			KodeProp:        sekolah.KodeProp,
// 			NamaSekolah:     sekolah.NamaSekolah,
// 			Status:          sekolah.Status,
// 			Npsn:            sekolah.NPSN,
// 		},
// 	}, nil
// }
