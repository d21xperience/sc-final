package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	pb "auth_service/generated"
	"auth_service/models"
	"auth_service/repository"
	"auth_service/services"
	"auth_service/utils"

	"github.com/go-redis/redis"
)

// AuthServiceServer dengan Redis Client sebagai Dependency Injection
type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	RedisClient    *redis.Client // Tambahkan Redis sebagai field
	sekolahService services.SekolahService
	authService    services.AuthService
	userProfile    services.UserProfileService
}

// Constructor untuk AuthServiceServer dengan Redis
func NewAuthServiceServer(redisClient *redis.Client) *AuthServiceServer {
	return &AuthServiceServer{RedisClient: redisClient}
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
	// Ambil data dari request
	user := req.GetUser()
	sekolah := req.GetSekolah()
	var query = repository.SekolahQuery{
		Npsn: sekolah.Npsn,
	}
	// Cek apakah sekolah sudah ada
	sekolahModel, err := s.sekolahService.GetSekolah(query)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			if user.Role == "admin" {
				// Buat sekolah baru
				sekolahModel, err = s.sekolahService.CreateSekolah(&models.Sekolah{
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
				})
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
		SekolahID: int32(sekolahModel.ID),
		Password: user.Password,
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

	if err := s.userProfile.CreateUserProfile(userProfileModel); err != nil {
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
				UserId:    userModel.ID,
				Username:  userModel.Username,
				Email:     userModel.Email,
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

// GetUserProfile - Mengambil profil pengguna berdasarkan UserID
func (s *AuthServiceServer) GetUserPofile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	profile, err := s.userProfile.GetUserProfileByID(req.UserId)
	if err != nil {
		log.Printf("Error fetching user profile: %v", err)
		return nil, errors.New("failed to retrieve user profile")
	}
	return &pb.GetUserProfileResponse{
		UserProfile: &pb.UserProfile{
			UserId:   profile.UserID,
			Nama:     profile.Nama,
			Jk:       profile.JK,
			Phone:    profile.Phone,
			TptLahir: profile.TptLahir,
			// TglLahir: profile.TglLahir.Format(),
			AlamatJalan: profile.AlamatJalan,
			KotaKab:     profile.KotaKab,
			Prov:        profile.Prov,
			KodePos:     profile.KodePos,
			NamaAyah:    profile.NamaAyah,
			NamaIbu:     profile.NamaIbu,
		},
	}, nil
}

// UpdateUserProfile - Memperbarui profil pengguna berdasarkan UserID
func (s *AuthServiceServer) UpdateUserPofile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received UpdateUserProfile request: %+v\n", req)

	// Cek apakah req atau req.UserProfile kosong
	if req == nil {
		log.Println("Request is nil")
		return nil, errors.New("invalid request: request is nil")
	}

	if req.UserProfile == nil {
		log.Println("UserProfile is nil in request")
		return nil, errors.New("invalid request: user profile is nil")
	}

	profile, err := s.userProfile.GetUserProfileByID(int64(req.GetUserId()))
	if err != nil {
		log.Printf("Error fetching user profile: %v", err)
		return nil, errors.New("user profile not found")
	}
	// profi := req.GetUserProfile()
	// log.Println(profi)
	// Perbarui data profil berdasarkan input
	profile.Nama = req.UserProfile.Nama
	profile.JK = req.UserProfile.Jk
	profile.Phone = req.UserProfile.Phone
	profile.TptLahir = req.UserProfile.TptLahir
	profile.AlamatJalan = req.UserProfile.AlamatJalan
	profile.KotaKab = req.UserProfile.KotaKab
	profile.Prov = req.UserProfile.Prov
	profile.KodePos = req.UserProfile.KodePos
	profile.NamaAyah = req.UserProfile.NamaAyah
	profile.NamaIbu = req.UserProfile.NamaIbu

	// Simpan perubahan ke database
	err = s.userProfile.UpdateUserProfile(profile)
	if err != nil {
		log.Printf("Error updating user profile: %v", err)
		return nil, errors.New("failed to update user profile")
	}

	return &pb.UpdateUserProfileResponse{
		Message: "Updated",
	}, nil
}

// UploadUserPhotoProfile - Mengunggah foto profil pengguna
func (s *AuthServiceServer) UploadUserPhotoProfile(stream pb.AuthService_UploadUserPhotoProfileServer) error {
	var userID int32
	var filePath string
	var file *os.File
	// var err error

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Jika selesai menerima file, kirim respons sukses
			return stream.SendAndClose(&pb.UploadUserPhotoResponse{
				Status:  "OK",
				FileUrl: filePath,
			})
		}
		if err != nil {
			log.Printf("Error receiving file chunk: %v", err)
			return err
		}

		// Ambil User ID dari request pertama
		if userID != 0 {
			userID = req.UserId
			filePath = fmt.Sprintf("uploads/%d_profile.jpg", userID)
			// Simpan ke database
			curentProfile, err := s.userProfile.GetUserProfileByID(int64(userID))
			if err != nil {
				log.Printf("Error getting user profile: %v", err)
			}
			err = s.userProfile.UpdateUserProfile(&models.UserProfile{
				ID:             curentProfile.ID,
				ProfilePicture: filePath,
			})
			if err != nil {
				log.Printf("Error updating user profile: %v", err)
			}
			// Buat file baru
			file, err = os.Create(filePath)
			if err != nil {
				log.Printf("Error creating file: %v", err)
				return err
			}
			defer file.Close()
		}

		// Tulis chunk ke file
		_, err = file.Write(req.FileChunk)
		if err != nil {
			log.Printf("Error writing file: %v", err)
			return err
		}
	}
}

func (s *AuthServiceServer) GetSekolah(ctx context.Context, req *pb.GetSekolahRequest) (*pb.GetSekolahResponse, error) {
	var query = repository.SekolahQuery{
		Npsn:      req.GetNpsn(),
		SekolahID: int(req.GetSekolahId()),
	}

	sekolah, err := s.sekolahService.GetSekolah(query)
	if err != nil {
		log.Printf("Error fetching school data: %v", err)
		return nil, errors.New("failed to retrieve school data")
	}

	return &pb.GetSekolahResponse{
		Nama: sekolah.NamaSekolah,
		SekolahData: &pb.Sekolah{
			SekolahId:       int32(sekolah.ID),
			SekolahIdEnkrip: sekolah.SekolahIDEnkrip,
			Kecamatan:       sekolah.Kecamatan,
			Kabupaten:       sekolah.Kabupaten,
			Propinsi:        sekolah.Propinsi,
			KodeKecamatan:   sekolah.Kecamatan,
			AlamatJalan:     sekolah.AlamatJalan,
			KodeKab:         sekolah.KodeKab,
			KodeProp:        sekolah.KodeProp,
			NamaSekolah:     sekolah.NamaSekolah,
			Status:          sekolah.Status,
			Npsn:            sekolah.NPSN,
		},
	}, nil
}