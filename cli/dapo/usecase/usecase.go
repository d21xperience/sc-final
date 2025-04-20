package usecase

import (
	"dapo/config"
	pb "dapo/generated"
	"dapo/models"
	"dapo/services"
	"dapo/utils"
	"fmt"
)

func ProcessSekolah(cfg *config.AppConfig, semesterID string) error {
	client := services.NewAPIClient(cfg.BaseURL, cfg.Token)
	if client == nil {
		return fmt.Errorf("Error client berisi nil")
	}
	data, err := services.GetOrFetch[models.SekolahResponse](
		client,
		cfg.NPSN,
		semesterID,
		cfg.OutputDir,
		"/WebService/getSekolah",
		map[string]string{
			"npsn":        cfg.NPSN,
			"semester_id": semesterID,
		},
	)
	if err != nil {
		fmt.Printf("Error fetching sekolah data: %v\n", err)
		return err
	}

	grpcClient, err := services.NewGRPCClient(cfg.GRPCHost, cfg.GRPCPort, cfg.GRPCTimeout)
	if err != nil {
		return err
	}
	defer grpcClient.Close()

	modSekolah := models.Sekolah{
		NSS:      data.Rows.NSS,
		NPSN:     data.Rows.NPSN,
		NomorFax: data.Rows.NomorFax,
		Website:  data.Rows.Website,
		Email:    data.Rows.Email,
	}
	pbSekolah := utils.ConvertModelToPB(&modSekolah, func(item *models.Sekolah) *pb.SekolahDapo {
		return &pb.SekolahDapo{
			SekolahId: item.SekolahID,
			Nama:      item.Nama,
		}
	})
	_, err = grpcClient.SendSekolahData(pbSekolah)
	return err
}

func ProcessPtk(cfg *config.AppConfig, semesterID string) error {
	client := services.NewAPIClient(cfg.BaseURL, cfg.Token)
	data, err := services.GetOrFetch[models.PTKTerdaftarResponse](
		client,
		cfg.NPSN,
		semesterID,
		cfg.OutputDir,
		"/WebService/getGtk",
		map[string]string{
			"npsn":        cfg.NPSN,
			"semester_id": semesterID,
		},
	)
	if err != nil {
		fmt.Printf("Error fetching sekolah data: %v\n", err)
		return err
	}

	grpcClient, err := services.NewGRPCClient(cfg.GRPCHost, cfg.GRPCPort, cfg.GRPCTimeout)
	if err != nil {
		return err
	}
	if grpcClient == nil {
		return fmt.Errorf("Error client berisi nil")
	}
	defer grpcClient.Close()

	var modPtk []models.PTKTerdaftar
	modPtk = append(modPtk, data.Rows...)

	pbPtkTerdaftar := utils.ConvertModelsToPB(utils.SliceToPointer(modPtk), func(item *models.PTKTerdaftar) *pb.PTKTerdaftar {
		return &pb.PTKTerdaftar{
			PtkTerdaftarId: item.PTKTerdaftarID,
			PtkId:          item.PTKID,
			TahunAjaranId:  item.TahunAjaranID,
			Ptk: &pb.PTK{
				PtkId:        item.PTKID,
				Nama:         item.Nama,
				JenisPtkId:   int32(item.JenisPTKID),
				JenisKelamin: item.JenisKelamin,
				TempatLahir:  item.TempatLahir,
				TanggalLahir: item.TanggalLahir,
				Nuptk:        item.NUPTK,
				Nip:          item.NIP,
			},
		}
	})
	_, err = grpcClient.SendPtkTerdaftarData(pbPtkTerdaftar)
	return err
}
func ProcessSiswa(cfg *config.AppConfig, semesterID string) error {
	client := services.NewAPIClient(cfg.BaseURL, cfg.Token)
	data, err := services.GetOrFetch[models.RegistrasiResponse](
		client,
		cfg.NPSN,
		semesterID,
		cfg.OutputDir,
		"/WebService/getPesertaDidik",
		map[string]string{
			"npsn":        cfg.NPSN,
			"semester_id": semesterID,
		},
	)
	if err != nil {
		fmt.Printf("Error fetching sekolah data: %v\n", err)
		return err
	}

	grpcClient, err := services.NewGRPCClient(cfg.GRPCHost, cfg.GRPCPort, cfg.GRPCTimeout)
	if err != nil {
		return err
	}
	if grpcClient == nil {
		return fmt.Errorf("Error client berisi nil")
	}
	defer grpcClient.Close()

	var modSiswa []models.PesertaDidik
	modSiswa = append(modSiswa, data.Rows...)

	pbSiswa := utils.ConvertModelsToPB(utils.SliceToPointer(modSiswa), func(item *models.PesertaDidik) *pb.Siswa {
		return &pb.Siswa{
			PesertaDidikId:  item.PesertaDidikID,
			Nis:             item.Nipd,
			Nisn:            item.Nisn,
			NmSiswa:         item.Nama,
			TempatLahir:     item.TempatLahir,
			TanggalLahir:    item.TanggalLahir,
			JenisKelamin:    item.JenisKelamin,
			Agama:           item.AgamaIDStr,
			AlamatSiswa:     item.AlamatJalan,
			TeleponSiswa:    item.NomorTeleponSeluler,
			NmAyah:          item.NamaAyah,
			NmIbu:           item.NamaIbu,
			PekerjaanAyah:   item.PekerjaanAyahStr,
			PekerjaanIbu:    item.PekerjaanIbuStr,
			DiterimaTanggal: item.TanggalMasukSekolah,
			Nik:             item.Nik,
		}
	})
	_, err = grpcClient.SendStudentData(pbSiswa)
	return err
}
func ProcessRombel(cfg *config.AppConfig, semesterID string) error {
	client := services.NewAPIClient(cfg.BaseURL, cfg.Token)
	data, err := services.GetOrFetch[models.RombelResponse](
		client,
		cfg.NPSN,
		semesterID,
		cfg.OutputDir,
		"/WebService/getRombonganBelajar",
		map[string]string{
			"npsn":        cfg.NPSN,
			"semester_id": semesterID,
		},
	)
	if err != nil {
		fmt.Printf("Error fetching sekolah data: %v\n", err)
		return err
	}

	grpcClient, err := services.NewGRPCClient(cfg.GRPCHost, cfg.GRPCPort, cfg.GRPCTimeout)
	if err != nil {
		return err
	}
	if grpcClient == nil {
		return fmt.Errorf("Error client berisi nil")
	}
	defer grpcClient.Close()

	var modRombel []models.RombonganBelajar
	modRombel = append(modRombel, data.Rows...)

	pbRombel := utils.ConvertModelsToPB(utils.SliceToPointer(modRombel), func(item *models.RombonganBelajar) *pb.Kelas {
		rombonganBelajarId := item.RombonganBelajarID
		semesterId := item.SemesterID
		return &pb.Kelas{
			RombonganBelajarId:  item.RombonganBelajarID,
			SekolahId:           "8a5bd957-66bc-4096-9ff1-fee096b87ba0",
			SemesterId:          item.SemesterID,
			JurusanId:           item.JurusanID,
			PtkId:               item.PTKID,
			NmKelas:             item.Nama,
			TingkatPendidikanId: int32(utils.StringToInt(item.TingkatPendidikanID)),
			NamaJurusanSp:       item.JurusanIDStr,
			KurikulumId:         int32(item.KurikulumID),
			JenisRombel:         int32(utils.StringToInt(item.JenisRombel)),
			JurusanSpId:         &item.JurusanIDStr,
			// JenisRombel:         item.JenisRombel,
			AnggotaKelas: utils.ConvertModelsToPB(utils.SliceToPointer(item.AnggotaRombel), func(item *models.AnggotaRombel) *pb.AnggotaKelas {
				return &pb.AnggotaKelas{
					AnggotaRombelId:    item.AnggotaRombelID,
					PesertaDidikId:     item.PesertaDidikID,
					RombonganBelajarId: rombonganBelajarId,
					SemesterId:         semesterId,
				}
			}),
			Pembelajaran: utils.ConvertModelsToPB(utils.SliceToPointer(item.Pembelajaran), func(item *models.Pembelajaran) *pb.Pembelajaran {
				return &pb.Pembelajaran{
					PembelajaranId:     item.PembelajaranID,
					RombonganBelajarId: rombonganBelajarId,
					MataPelajaranId:    int32(item.MataPelajaranID),
					SemesterId:         semesterId,
					PtkTerdaftarId:     item.PTKTerdaftarID,
					StatusDiKurikulum:  int32(utils.StringToInt(item.StatusDiKurikulum)),
					NamaMataPelajaran:  item.NamaMataPelajaran,
					IndukPembelajaran:  utils.SafeString(item.IndukPembelajaranID),
				}
			}),
		}
	})
	_, err = grpcClient.SendRombel(pbRombel)
	return err
}
