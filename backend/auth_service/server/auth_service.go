package server

import (
	"context"
	"errors"
	"fmt"
	"log"

	pb "auth_service/generated"
	"auth_service/models"
	"auth_service/services"

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
	log.Println(req.Email)
	log.Println(req.GetEmail())

	token, err := s.authService.Login(req.Email, req.Password)
	if err != nil {
		log.Printf("Error publishing to Redis: %v", err)
		return nil, err
	}

	return &pb.LoginResponse{Token: token}, nil
}

func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Ambil data dari request
	user := req.GetUser()
	sekolah := req.GetSekolah()
	// userProfile := req.GetUserProfile()

	log.Printf("cek poin 1")
	log.Printf("Menerima pendaftaran: User %s, Sekolah %s\n", user.GetUsername(), sekolah.GetNpsn())

	// Cek apakah sekolah sudah ada
	sekolahModel, err := s.sekolahService.GetSekolahByNpsn(sekolah.Npsn)
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
					// CreatedAt: sekolah.Crea,
				})
				if err != nil {
					log.Printf("Gagal membuat sekolah: %v", err)
					return nil, fmt.Errorf("gagal membuat sekolah: %w", err)
				}
			} else {
				return nil, fmt.Errorf("sekolah tidak ditemukan")
			}
		} else {
			log.Printf("Server error saat mencari sekolah: %v", err)
			return nil, fmt.Errorf("server error: %w", err)
		}
	}
	log.Printf("cek point")
	// Hubungkan user dengan sekolah
	userModel := &models.User{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		SchoolID: sekolahModel.ID,
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

	log.Println("User berhasil didaftarkan")
	response := &pb.RegisterResponse{
		Message: "User registered successfully",
		UserId:  fmt.Sprintf("%d", userModel.ID),
	}

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

// ✅ GetUserProfile - Mengambil profil pengguna berdasarkan UserID
func (s *AuthServiceServer) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	profile, err := s.userProfile.GetUserProfileByID(req.UserId)
	if err != nil {
		log.Printf("Error fetching user profile: %v", err)
		return nil, errors.New("failed to retrieve user profile")
	}

	return &pb.GetUserProfileResponse{
		Name: "OK",
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

// ✅ GetSekolahByNpsn - Mengambil data sekolah berdasarkan NPSN
func (s *AuthServiceServer) GetSekolahByNpsn(ctx context.Context, req *pb.GetSekolahByNpsnRequest) (*pb.GetSekolahByNpsnResponse, error) {
	sekolah, err := s.sekolahService.GetSekolahByNpsn(req.Npsn)
	if err != nil {
		log.Printf("Error fetching school data: %v", err)
		return nil, errors.New("failed to retrieve school data")
	}

	return &pb.GetSekolahByNpsnResponse{
		Nama: "nama_sekolah",
		SekolahData: &pb.Sekolah{
			SekolahIdEnkrip: sekolah.SekolahIDEnkrip,
			Kecamatan:       sekolah.Kecamatan,
			Kabupaten:       sekolah.Kabupaten,
			Propinsi:        sekolah.Propinsi,
			KodeKecamatan:   sekolah.Kecamatan,
		},
	}, nil
}
