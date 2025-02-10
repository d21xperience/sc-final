package server

import (
	"context"
	pb "sekolah/generated"
)

type TranskripNilaiService struct {
	pb.UnimplementedTranskripNilaiServiceServer
}

func NewTranskripNilaiService() *TranskripNilaiService {
	return &TranskripNilaiService{}
}

func (s *TranskripNilaiService) CreateTranskripNilai(ctx context.Context, req *pb.CreateTranskripNilaiRequest) (*pb.CreateTranskripNilaiResponse, error) {

	return &pb.CreateTranskripNilaiResponse{
		Status:  true,
		Message: "sukses",
	}, nil

}
func (s *TranskripNilaiService) GetTranskripNilai(ctx context.Context, req *pb.GetTranskripNilaiRequest) (*pb.GetTranskripNilaiResponse, error) {

	return &pb.GetTranskripNilaiResponse{

		Status:  true,
		Message: "sukses",
	}, nil

}
func (s *TranskripNilaiService) UpdateTranskripNilai(ctx context.Context, req *pb.UpdateTranskripNilaiRequest) (*pb.UpdateTranskripNilaiResponse, error) {

	return &pb.UpdateTranskripNilaiResponse{
		Status:  true,
		Message: "sukses",
	}, nil

}
func (s *TranskripNilaiService) DeleteTranskripNilai(ctx context.Context, req *pb.DeleteTranskripNilaiRequest) (*pb.DeleteTranskripNilaiResponse, error) {

	return &pb.DeleteTranskripNilaiResponse{
		Status:  true,
		Message: "sukses",
	}, nil

}
func (s *TranskripNilaiService) UploadTranskripNilai(ctx context.Context, req *pb.UploadTranskripNilaiRequest) (*pb.UploadTranskripNilaiResponse, error) {

	return &pb.UploadTranskripNilaiResponse{
		Status:  true,
		Message: "sukses",
	}, nil

}
