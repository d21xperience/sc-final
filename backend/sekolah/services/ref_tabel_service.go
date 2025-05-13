package services

import (
	"context"
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/utils"
)

type ReferensiServiceServer struct {
	pb.UnimplementedReferensiServiceServer
	repoBentukPendidikan  repositories.GenericRepository[models.BentukPendidikan]
	repoJenjangPendidikan repositories.GenericRepository[models.JenjangPendidikan]
	repoTingkatPendidikan repositories.GenericRepository[models.TingkatPendidikan]
	repoStatusKepemilikan repositories.GenericRepository[models.StatusKepemilikan]
	repoKurikulum         repositories.GenericRepository[models.Kurikulum]
	repoJurusan           repositories.GenericRepository[models.Jurusan]
	repoMapel             repositories.GenericRepository[models.MataPelajaran]
	repoGelarAkademik     repositories.GenericRepository[models.GelarAkademik]
}

func NewReferensiServiceServer() *ReferensiServiceServer {
	repoBentukPendidikan := repositories.NewBentukPendidikanRepository(config.DB)
	repoJenjangPendidikan := repositories.NewJenjangPendidikanRepository(config.DB)
	repoTingkatPendidikan := repositories.NewTingkatPendidikanRepository(config.DB)
	repoStatusKepemilikan := repositories.NewStatusKepemilikanRepository(config.DB)
	repoKurikulum := repositories.NewKurikulumRepository(config.DB)
	repoJurusan := repositories.NewJurusanRepository(config.DB)
	repoMapel := repositories.NewMapelRepository(config.DB)
	repoGelarAkademik := repositories.NewGelarAkademikRepository(config.DB)
	return &ReferensiServiceServer{
		repoBentukPendidikan:  *repoBentukPendidikan,
		repoJenjangPendidikan: *repoJenjangPendidikan,
		repoTingkatPendidikan: *repoTingkatPendidikan,
		repoStatusKepemilikan: *repoStatusKepemilikan,
		repoKurikulum:         *repoKurikulum,
		repoJurusan:           *repoJurusan,
		repoMapel:             *repoMapel,
		repoGelarAkademik:     *repoGelarAkademik,
	}
}

func (s *ReferensiServiceServer) GetBentukPendidikan(ctx context.Context, req *pb.Empty) (*pb.GetBentukPendidikanResponse, error) {
	mod, err := s.repoBentukPendidikan.FindAll(ctx, "ref", 100, 0)
	if err != nil {
		return nil, err
	}
	res := utils.ConvertModelsToPB(mod, func(re *models.BentukPendidikan) *pb.BentukPendidikan {
		return &pb.BentukPendidikan{
			BentukPendidikanId:   uint32(re.BentukPendidikanID),
			Nama:                 re.Nama,
			JenjangPaud:          uint32(re.JenjangPaud),
			JenjangTk:            uint32(re.JenjangTk),
			JenjangSd:            uint32(re.JenjangSd),
			JenjangSmp:           uint32(re.JenjangSmp),
			JenjangSma:           uint32(re.JenjangSma),
			JenjangTinggi:        uint32(re.JenjangTinggi),
			DirektoratPembinaan:  *re.DirektoratPembinaan,
			Aktif:                uint32(re.Aktif),
			FormalitasPendidikan: re.FormalitasPendidikan,
		}
	})
	return &pb.GetBentukPendidikanResponse{
		BentukPendidikan: res,
	}, nil
}
func (s *ReferensiServiceServer) GetJenjang(ctx context.Context, req *pb.Empty) (*pb.GetJenjangResponse, error) {
	mod, err := s.repoJenjangPendidikan.FindAll(ctx, "ref", 100, 0)
	if err != nil {
		return nil, err
	}
	res := utils.ConvertModelsToPB(mod, func(re *models.JenjangPendidikan) *pb.Jenjang {
		return &pb.Jenjang{
			JenjangPendidikanId: uint32(re.JenjangPendidikanID),
			Nama:                re.Nama,
			JenjangLembaga:      uint32(re.JenjangLembaga),
			JenjangOrang:        uint32(re.JenjangOrang),
		}
	})
	return &pb.GetJenjangResponse{
		Jenjang: res,
	}, nil
}
func (s *ReferensiServiceServer) GetTingkatPendidikan(ctx context.Context, req *pb.GetTingkatPendidikanRequest) (*pb.GetTingkatPendidikanResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"JenjangPendidikanId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	conditions := map[string]interface{}{
		"jenjang_pendidikan_id": req.GetJenjangPendidikanId(),
	}
	mod, err := s.repoTingkatPendidikan.FindAllByConditions(ctx, "ref", conditions, 100, 0, nil)
	if err != nil {
		return nil, err
	}
	res := utils.ConvertModelsToPB(mod, func(re *models.TingkatPendidikan) *pb.TingkatPendidikan {
		return &pb.TingkatPendidikan{
			TingkatPendidikanId: int32(re.TingkatPendidikanID),
			JenjangPendidikanId: uint32(re.JenjangPendidikanID),
			Kode:                re.Kode,
			Nama:                re.Nama,
		}
	})
	return &pb.GetTingkatPendidikanResponse{
		TingkatPendidikan: res,
	}, nil
}
func (s *ReferensiServiceServer) GetStatusKepemilikan(ctx context.Context, req *pb.Empty) (*pb.GetStatusKepemilikanResponse, error) {
	mod, err := s.repoStatusKepemilikan.FindAll(ctx, "ref", 100, 0)
	if err != nil {
		return nil, err
	}
	res := utils.ConvertModelsToPB(mod, func(re *models.StatusKepemilikan) *pb.StatusKepemilikan {
		return &pb.StatusKepemilikan{
			StatusKepemilikanId: uint32(re.StatusKepemilikanID),
			Nama:                re.Nama,
		}
	})
	return &pb.GetStatusKepemilikanResponse{
		StatusKepemilikan: res,
	}, nil
}
func (s *ReferensiServiceServer) GetKurikulum(ctx context.Context, req *pb.GetKurikulumRequest) (*pb.GetKurikulumResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"jurusanId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	conditions := map[string]any{
		"jurusan_id": req.GetJurusanId(),
	}
	mod, err := s.repoKurikulum.FindAllByConditions(ctx, "ref", conditions, 1000, 0, nil)
	if err != nil {
		return nil, err
	}
	res := utils.ConvertModelsToPB(mod, func(re *models.Kurikulum) *pb.Kurikulum {
		return &pb.Kurikulum{
			KurikulumId:         uint32(re.KurikulumID),
			NamaKurikulum:       re.NamaKurikulum,
			MulaiBerlaku:        re.MulaiBerlaku.Format("2006-01-02"),
			SistemSks:           uint32(re.SistemSKS),
			TotalSks:            uint32(re.TotalSKS),
			JenjangPendidikanId: uint32(re.JenjangPendidikanID),
			JurusanId:           utils.SafeString(re.JurusanID),
		}
	})
	return &pb.GetKurikulumResponse{
		Kurikulum: res,
	}, nil
}
func (s *ReferensiServiceServer) GetJurusan(ctx context.Context, req *pb.GetJurusanRequest) (*pb.GetJurusanResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"JenjangPendidikanId", "Param"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	param := req.GetParam()
	conditions := map[string]any{
		"jenjang_pendidikan_id": req.GetJenjangPendidikanId(),
		param:                   1,
	}
	mod, err := s.repoJurusan.FindAllByConditions(ctx, "ref", conditions, 100, 0, nil)
	if err != nil {
		return nil, err
	}
	res := utils.ConvertModelsToPB(mod, func(re *models.Jurusan) *pb.Jurusan {
		return &pb.Jurusan{
			JurusanId:           re.JurusanID,
			NamaJurusan:         re.NamaJurusan,
			UntukSma:            uint32(re.UntukSMA),
			UntukSmk:            uint32(re.UntukSMK),
			UntukPt:             uint32(re.UntukPT),
			UntukSlb:            uint32(re.UntukSLB),
			UntukSmklb:          uint32(re.UntukSMKLB),
			JenjangPendidikanId: utils.PointerToUint32(utils.Uint16ToUint32Pointer(re.JenjangPendidikanID)),
			JurusanInduk:        utils.SafeString(re.JurusanInduk),
			LevelBidangId:       re.LevelBidangID,
		}
	})
	return &pb.GetJurusanResponse{
		Jurusan: res,
	}, nil
}

func (s *ReferensiServiceServer) GetMapel(ctx context.Context, req *pb.GetMapelRequest) (*pb.GetMapelResponse, error) {
	// // Daftar field yang wajib diisi
	// requiredFields := []string{"Param"}
	// // Validasi request
	// err := utils.ValidateFields(req, requiredFields)
	// if err != nil {
	// 	return nil, err
	// }
	conditions := map[string]any{}
	mapelId := req.GetMapelId()
	if mapelId != "" {
		conditions["mapel_id"] = mapelId
	}
	mod, err := s.repoMapel.FindAllByConditions(ctx, "ref", conditions, 7000, 0, nil)
	if err != nil {
		return nil, err
	}
	res := utils.ConvertModelsToPB(mod, func(item *models.MataPelajaran) *pb.Mapel {
		return &pb.Mapel{
			MataPelajaranId:     item.MataPelajaranID,
			Nama:                item.Nama,
			PilihanSekolah:      item.PilihanSekolah,
			PilihanBuku:         item.PilihanBuku,
			PilihanKepengawasan: item.PilihanKepengawasan,
			PilihanEvaluasi:     item.PilihanEvaluasi,
			JurusanId:           utils.SafeString(item.JurusanID),
		}
	})
	return &pb.GetMapelResponse{
		Mapel: res,
	}, nil
}

func (s *ReferensiServiceServer) GetGelarAkademik(ctx context.Context, req *pb.GetGelarAkademikRequest) (*pb.GetGelarAkademikResponse, error) {
	gelarAkademikModel, err := s.repoGelarAkademik.FindAll(ctx, "ref", 1000, 0)
	if err != nil {
		return nil, err
	}
	results := utils.ConvertModelsToPB(gelarAkademikModel, func(item *models.GelarAkademik) *pb.GelarAkademik {
		return &pb.GelarAkademik{
			Kode:        item.Kode,
			Nama:        item.Nama,
			PosisiGelar: item.PosisiGelar,
		}
	})
	return &pb.GetGelarAkademikResponse{
		GelarAkademik: results,
	}, nil
}
