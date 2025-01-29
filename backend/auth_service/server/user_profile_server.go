package server

import (
	pb "auth_service/generated"
	"auth_service/services"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserProfileServiceServer struct {
	pb.UnimplementedUserProfileServiceServer
	userProfile services.UserProfileService
}

// GetUserProfile - Mengambil profil pengguna berdasarkan UserID
func (s *UserProfileServiceServer) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	userID := req.GetUserId()
	profile, err := s.userProfile.GetUserProfileByID(int64(userID))
	if err != nil {
		log.Printf("Error fetching user profile: %v", err)
		return nil, errors.New("failed to retrieve user profile")
	}

	return &pb.GetUserProfileResponse{
		UserProfile: &pb.UserProfile{
			UserId:   profile.UserId,
			Nama:     profile.Nama,
			Jk:       profile.JK,
			Phone:    profile.Phone,
			TptLahir: profile.TptLahir,
			// TglLahir: profile.TglLahir.Format("2007-12-21",),
			AlamatJalan: profile.AlamatJalan,
			KotaKab:     profile.KotaKab,
			Prov:        profile.Prov,
			KodePos:     profile.KodePos,
			NamaAyah:    profile.NamaAyah,
			NamaIbu:     profile.NamaIbu,
			PhotoUrl:    profile.PhotoUrl,
		},
	}, nil
}

func (s *UserProfileServiceServer) GetUserProfilePhoto(ctx context.Context, req *pb.GetUserProfilePhotoRequest) (*pb.GetUserProfilePhotoResponse, error) {
	// Path ke file foto profil pengguna
	filePath := fmt.Sprintf("photos/%d.jpg", req.GetUserId())

	// Baca file foto
	photoBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Photo not found")
	}

	return &pb.GetUserProfilePhotoResponse{
		Photo:       photoBytes,
		ContentType: "image/jpeg", // Ganti sesuai tipe file
	}, nil
}

// UpdateUserProfile - Memperbarui profil pengguna berdasarkan UserID
func (s *UserProfileServiceServer) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponse, error) {
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

// Implementasi UploadFile
func (s *UserProfileServiceServer) UploadUserPhotoProfile(stream pb.UserProfileService_UploadUserPhotoProfileServer) error {
	var filename string
	// var userId int64
	fileBuffer := []byte{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Menyimpan data ke file setelah menerima semua chunk
			err := os.WriteFile(filename, fileBuffer, 0644)
			if err != nil {
				return err
			}
			return stream.SendAndClose(&pb.UploadUserPhotoProfileResponse{
				Message: "File uploaded successfully!",
				Status:  true,
			})
		}
		if err != nil {
			log.Printf("Error receiving data: %v", err)
			return err
		}
		// Log data yang diterima
		log.Printf("Received chunk: %s, filename: %s", string(req.GetChunk()), req.GetFilename())
		// Mengumpulkan nama file dan chunk data
		filename = req.GetFilename()
		fileBuffer = append(fileBuffer, req.GetChunk()...)
		// userId = int64(req.GetUserId())
	}

	// // Lokasi file foto profil
	// photoPath := fmt.Sprintf("./uploads/profile_photos/%v.jpg", userId)

	// // Periksa apakah file ada
	// if _, err := os.Stat(photoPath); os.IsNotExist(err) {
	// 	// return nil, status.Errorf(codes.NotFound, "Photo for user_id %d not found", userID)
	// 	log.Printf("status%v: ", err)
	// }

	// Buat URL public (asumsikan server berjalan di localhost:8080)
	// photoURL := fmt.Sprintf("http://localhost:8080/static/profile_photos/%d.jpg", userID)

}

// Implementasi DownloadFile
func (s *UserProfileServiceServer) DownloadFile(ctx context.Context, req *pb.DownloadUserPhotoProfileRequest, stream pb.UserProfileService_DownloadUserPhotoProfileServer) error {
	filename := req.GetFilename()

	// Membuka file yang diminta
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Membaca file dalam bentuk chunk dan mengirimkan ke client
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Mengirimkan chunk ke client
		stream.Send(&pb.DownloadUserPhotoProfileResponse{
			Chunk: buffer[:n],
		})
	}

	return nil
}
