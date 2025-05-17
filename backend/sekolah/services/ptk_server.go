package services

import (
	"context"
	"fmt"
	"log"
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/utils"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PTKServiceServer struct {
	pb.UnimplementedPTKServiceServer
	repo repositories.GenericRepository[models.TabelPTK]
}
type PTKConflictResponse struct {
	Index  int       `json:"index"`
	PtkID  uuid.UUID `json:"ptk_id"`
	Nama   string    `json:"nama"`
	Reason string    `json:"reason"`
}

func NewPTKServiceServer() *PTKServiceServer {
	repoPTK := repositories.NewPTKRepository(config.DB)
	return &PTKServiceServer{
		repo: *repoPTK,
	}
}

// **CreatePTK**
func (s *PTKServiceServer) CreatePTK(ctx context.Context, req *pb.CreatePTKRequest) (*pb.CreatePTKResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	PTKReq := req.GetPTK()
	PTKModel := utils.ConvertPBToModels(PTKReq, func(item *pb.PTK) *models.TabelPTK {
		return &models.TabelPTK{
			PtkID:             utils.StringToUUID(item.PtkId),
			Nama:              item.Nama,
			NIP:               &item.Nip,
			JenisPtkID:        item.JenisPtkId,
			JenisKelamin:      item.JenisKelamin,
			TempatLahir:       item.TempatLahir,
			TanggalLahir:      utils.TimeToPointer(item.TanggalLahir),
			NUPTK:             &item.Nuptk,
			AlamatJalan:       item.AlamatJalan,
			StatusKeaktifanID: item.StatusKeaktifanId,
			GelarDepan:        &item.GelarBelakang,
			GelarBelakang:     &item.GelarBelakang,
			NIP_NIY:           &item.NipNiy,
		}
	})
	conflicts, err := s.repo.SaveManyWithConflictCheck(ctx, schemaName, PTKModel, "PtkID", "ptk_id", 100, nil)
	// convert conflict rows ke response JSON
	if err != nil {
		return nil, status.Errorf(codes.Internal, "insert failed: %v", err)
	}

	conflictProto := repositories.ConvertConflictsToProto(conflicts, "PtkID", "Nama")

	return &pb.CreatePTKResponse{
		Message: "PTK berhasil ditambahkan",
		Status:  true,
		Conflicts: &pb.ConflictResponse{
			Message:       "Sebagian data berhasil disimpan",
			Conflicts:     conflictProto,
			TotalConflict: int32(len(conflictProto)),
		},
	}, nil
}

// **GetPTK**
func (s *PTKServiceServer) GetPTK(ctx context.Context, req *pb.GetPTKRequest) (*pb.GetPTKResponse, error) {
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	ptkId := req.GetPtkId()
	var conditions = map[string]any{
		"soft_delete": 0,
	}
	if ptkId != "" {
		conditions["ptk_id"] = ptkId
	}
	// log.Print(conditions)
	anggotaPTKModel, err := s.repo.FindWithPreloadAndJoins(ctx, schemaName, nil, nil, conditions, nil, []string{"tabel_ptk.nama"}, false)
	if err != nil {
		return nil, err
	}

	ptkList := utils.ConvertModelsToPB(anggotaPTKModel, func(item models.TabelPTK) *pb.PTK {
		return &pb.PTK{
			PtkId:             item.PtkID.String(),
			Nama:              item.Nama,
			Nuptk:             utils.SafeString(item.NUPTK),
			JenisKelamin:      item.JenisKelamin,
			TempatLahir:       item.TempatLahir,
			AlamatJalan:       item.AlamatJalan,
			Nip:               utils.SafeString(item.NIP),
			JenisPtkId:        item.JenisPtkID,
			TanggalLahir:      item.TanggalLahir.Format("2006-01-02"),
			StatusKeaktifanId: item.StatusKeaktifanID,
			GelarDepan:        utils.SafeString(item.GelarBelakang),
			GelarBelakang:     utils.SafeString(item.GelarBelakang),
			NipNiy:            utils.SafeString(item.NIP_NIY),
			
		}
	})

	return &pb.GetPTKResponse{
		PTK: ptkList,
	}, nil
}

// **UpdatePTK**
func (s *PTKServiceServer) UpdatePTK(ctx context.Context, req *pb.UpdatePTKRequest) (*pb.UpdatePTKResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	PTKReq := req.GetPTK()
	PTKModels := utils.ConvertPBToModels(PTKReq, func(item *pb.PTK) *models.TabelPTK {
		return &models.TabelPTK{
			PtkID:             utils.StringToUUID(item.PtkId),
			Nama:              item.Nama,
			NIP:               &item.Nip,
			JenisPtkID:        item.JenisPtkId,
			JenisKelamin:      item.JenisKelamin,
			TempatLahir:       item.TanggalLahir,
			TanggalLahir:      utils.TimeToPointer(item.TanggalLahir),
			NUPTK:             &item.Nuptk,
			AlamatJalan:       item.AlamatJalan,
			StatusKeaktifanID: item.StatusKeaktifanId,
			GelarDepan:        &item.GelarBelakang,
			GelarBelakang:     &item.GelarBelakang,
			NIP_NIY:           &item.NipNiy,
		}
	})
	for _, v := range PTKModels {
		err := s.repo.Update(ctx, v, schemaName, "ptk_id", v.PtkID.String())
		if err != nil {
			log.Printf("Gagal memperbarui PTK: %v", err)
			return nil, fmt.Errorf("gagal memperbarui PTK: %w", err)
		}
	}

	return &pb.UpdatePTKResponse{
		Message: "PTK berhasil diperbarui",
		Status:  true,
	}, nil
}

// // **DeletePTK**
func (s *PTKServiceServer) DeletePTK(ctx context.Context, req *pb.DeletePTKRequest) (*pb.DeletePTKResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	PTKID := req.GetPtkId()

	err = s.repo.Delete(ctx, PTKID, schemaName, "ptk_id")
	if err != nil {
		log.Printf("Gagal menghapus PTK: %v", err)
		return nil, fmt.Errorf("gagal menghapus PTK: %w", err)
	}

	return &pb.DeletePTKResponse{
		Message: "PTK berhasil dihapus",
		Status:  true,
	}, nil
}

// func ToPTKConflictResponse(conflicts []repositories.ConflictRow[models.TabelPTK]) []PTKConflictResponse {
// 	res := make([]PTKConflictResponse, 0, len(conflicts))
// 	for _, c := range conflicts {
// 		res = append(res, PTKConflictResponse{
// 			Index:  c.Index,
// 			PtkID:  c.Entity.PtkID,
// 			Nama:   c.Entity.Nama,
// 			Reason: c.Reason,
// 		})
// 	}
// 	return res
// }
