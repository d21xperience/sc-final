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
				PtkId: item.PTKID,
				Nama:  item.Nama,
				// JenisPtkId:   int32(utils.ParseInt(item.JenisPTKID)),
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
		}
	})
	_, err = grpcClient.SendStudentData(pbSiswa)
	return err
}
