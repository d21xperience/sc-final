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
)

type RombelServiceServer struct {
	pb.UnimplementedKelasServiceServer
	repo              repositories.GenericRepository[models.RombonganBelajar]
	repoRombelAnggota repositories.GenericRepository[models.RombelAnggota]
	repoSemester      repositories.GenericRepository[models.Semester]
}

func NewRombelServiceServer() *RombelServiceServer {
	repoRombel := repositories.NewrombonganBelajarRepository(config.DB)
	repoRombelAnggota := repositories.NewRombelAnggotaRepository(config.DB)
	repoSemester := repositories.NewSemesterRepository(config.DB)
	return &RombelServiceServer{
		repo:              *repoRombel,
		repoRombelAnggota: *repoRombelAnggota,
		repoSemester:      *repoSemester,
	}
}

// **CreateKelas**
func (s *RombelServiceServer) CreateKelas(ctx context.Context, req *pb.CreateKelasRequest) (*pb.CreateKelasResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"SchemaName", "Kelas"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaName()
	kelas := req.Kelas
	// rombelId := uuid.New()
	kelasModel := utils.ConvertPBToModels(kelas, func(item *pb.Kelas) *models.RombonganBelajar {
		return &models.RombonganBelajar{
			RombonganBelajarId:  utils.StringToUUID(item.RombonganBelajarId),
			SekolahId:           utils.StringToUUID(item.SekolahId),
			SemesterId:          item.SemesterId,
			JurusanId:           item.JurusanId,
			PtkID:               utils.StringToUUID(item.PtkId),
			NmKelas:             item.NmKelas,
			TingkatPendidikanId: item.TingkatPendidikanId,
			JenisRombel:         item.JenisRombel,
			NamaJurusanSp:       item.NamaJurusanSp,
			KurikulumId:         item.KurikulumId,
		}
	})
	// simpan kelas ke database
	res := s.repo.SaveMany(ctx, schemaName, kelasModel,100)
	if res != nil {
		return nil, err
	}
	
	return &pb.CreateKelasResponse{
		Message: "Kelas berhasil ditambahkan",
		Status:  true,
	}, nil
}

// func (s *RombelServiceServer) CreateBanyakKelas(ctx context.Context, req *pb.CreateBanyakKelasRequest) (*pb.CreateBanyakKelasResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"SchemaName", "Kelas"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaName()
// 	kelas := req.GetKelas()

// 	rombelId := uuid.New()
// 	sekolahId, err := uuid.Parse(kelas.SekolahId)
// 	if err != nil {
// 		fmt.Println("Error parsing UUID:", err)
// 	}
// 	ptkId, err := uuid.Parse(kelas.PtkId)
// 	if err != nil {
// 		fmt.Println("Error parsing UUID:", err)
// 	}
// 	kelasModels := ConvertPBToModels(kelas, func(rom *pb.Kelas) *models.RombonganBelajar {
// 		return &models.RombonganBelajar{
// 			RombonganBelajarId:  rombelId,
// 			SekolahId:           sekolahId,
// 			SemesterId:          rom.SemesterId,
// 			JurusanId:           rom.JurusanId,
// 			PtkID:               ptkId,
// 			NmKelas:             rom.NmKelas,
// 			TingkatPendidikanId: rom.TingkatPendidikanId,
// 			JenisRombel:         rom.JenisRombel,
// 			NamaJurusanSp:       rom.NamaJurusanSp,
// 			// JurusanSpId:         rom.JurusanSpId,
// 			KurikulumId: rom.KurikulumId,
// 		}
// 	})
// 	err = s.repo.SaveMany(ctx, schemaName, kelasModels, 100)
// 	if err != nil {
// 		log.Printf("Gagal menyimpan Kelas: %v", err)
// 		return nil, fmt.Errorf("gagal menyimpan Kelas: %w", err)
// 	}

// 	return &pb.CreateBanyakKelasResponse{
// 		Message: "Kelas berhasil ditambahkan",
// 		Status:  true,
// 	}, nil
// }

// **GetKelas**
func (s *RombelServiceServer) GetKelas(ctx context.Context, req *pb.GetKelasRequest) (*pb.GetKelasResponse, error) {
	var err error
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "SemesterId"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaName()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}

	// Cek apakah harus mengambil semua data atau data spesifik berdasarkan SemesterId
	// kelasId := req.GetKelasId()
	semesterId := req.GetSemesterId()
	var rombelModel []models.RombonganBelajar
	var conditions map[string]any
	if req.KelasId != "" {
		conditions = map[string]any{
			"rombongan_belajar_id": req.KelasId,
		}
	} else {
		conditions = map[string]any{
			"semester_id": semesterId,
		}
	}
	joins := []string{
		"JOIN tabel_ptk ON tabel_kelas.ptk_id = tabel_ptk.ptk_id",
		fmt.Sprintf("JOIN ref.jurusan ON %s.tabel_kelas.jurusan_id = ref.jurusan.jurusan_id", schemaName),
		fmt.Sprintf("JOIN ref.kurikulum ON %s.tabel_kelas.kurikulum_id = ref.kurikulum.kurikulum_id", schemaName),
		fmt.Sprintf("JOIN ref.tingkat_pendidikan ON %s.tabel_kelas.tingkat_pendidikan_id = ref.tingkat_pendidikan.tingkat_pendidikan_id", schemaName),
	}
	preloads := []string{"PTK", "Jurusan", "Kurikulum", "TingkatPendidikan"}

	groupByColumns := []string{"tabel_kelas.rombongan_belajar_id"} // Hindari duplikasi
	rombelModel, err = s.repo.FindWithPreloadAndJoins(ctx, schemaName, joins, preloads, conditions, groupByColumns)
	if err != nil {
		return nil, err
	}

	banyakKelasList := utils.ConvertModelsToPB(rombelModel, func(kelas models.RombonganBelajar) *pb.Kelas {
		// jurusanSPId, err := utils.ConvertUUIDToStringViceVersa(kelas.JurusanSpId)
		// if err != nil {
		// 	return nil
		// }
		jmlhAnggota, err := s.repoRombelAnggota.CountRows(ctx, schemaName, "rombongan_belajar_id", kelas.RombonganBelajarId.String())
		if err != nil {
			return nil
		}
		return &pb.Kelas{
			RombonganBelajarId:  kelas.RombonganBelajarId.String(),
			SekolahId:           kelas.SekolahId.String(),
			SemesterId:          kelas.SemesterId,
			JurusanId:           kelas.JurusanId,
			PtkId:               kelas.PtkID.String(),
			NmKelas:             kelas.NmKelas,
			TingkatPendidikanId: kelas.TingkatPendidikanId,
			JenisRombel:         kelas.JenisRombel,
			NamaJurusanSp:       kelas.NamaJurusanSp,
			// JurusanSpId:         jurusanSPId.(*string),
			KurikulumId: kelas.KurikulumId,
			Ptk: &pb.PTK{
				PtkId:             kelas.PTK.PtkID.String(),
				Nama:              kelas.PTK.Nama,
				JenisKelamin:      kelas.PTK.JenisKelamin,
				JenisPtkId:        kelas.PTK.JenisPtkID,
				TempatLahir:       kelas.PTK.TempatLahir,
				TanggalLahir:      kelas.PTK.TanggalLahir.Format("2006-01-02"),
				AlamatJalan:       kelas.PTK.AlamatJalan,
				StatusKeaktifanId: kelas.PTK.StatusKeaktifanID,
				Nuptk:             utils.SafeString(kelas.PTK.NUPTK),
				Nip:               utils.SafeString(kelas.PTK.NIP),
			},
			Jurusan: &pb.Jurusan{
				JurusanId:           kelas.Jurusan.JurusanID,
				NamaJurusan:         kelas.Jurusan.NamaJurusan,
				JenjangPendidikanId: utils.PointerToUint32(utils.Uint16ToUint32Pointer(kelas.Jurusan.JenjangPendidikanID)),
				UntukSma:            uint32(kelas.Jurusan.UntukSMA),
				UntukSmk:            uint32(kelas.Jurusan.UntukSMK),
				UntukPt:             uint32(kelas.Jurusan.UntukPT),
				UntukSlb:            uint32(kelas.Jurusan.UntukSLB),
				UntukSmklb:          uint32(kelas.Jurusan.UntukSMKLB),
				JurusanInduk:        utils.SafeString(kelas.Jurusan.JurusanInduk),
				LevelBidangId:       kelas.Jurusan.LevelBidangID,
			},
			Kurikulum: &pb.Kurikulum{
				KurikulumId:         uint32(kelas.Kurikulum.KurikulumID),
				NamaKurikulum:       kelas.Kurikulum.NamaKurikulum,
				MulaiBerlaku:        kelas.Kurikulum.MulaiBerlaku.Format("2006-01-02"),
				JenjangPendidikanId: uint32(kelas.Kurikulum.JenjangPendidikanID),
				SistemSks:           uint32(kelas.Kurikulum.SistemSKS),
				JurusanId:           *kelas.Kurikulum.JurusanID,
			},
			TingkatPendidikan: &pb.TingkatPendidikan{
				TingkatPendidikanId: uint32(kelas.TingkatPendidikan.TingkatPendidikanID),
				Kode:                kelas.TingkatPendidikan.Kode,
				Nama:                kelas.TingkatPendidikan.Nama,
				JenjangPendidikanId: uint32(kelas.TingkatPendidikan.JenjangPendidikanID),
			},
			JumlahAnggota: uint32(jmlhAnggota),
		}
	})
	return &pb.GetKelasResponse{
		Kelas: banyakKelasList,
	}, nil
}

// **UpdateKelas**
func (s *RombelServiceServer) UpdateKelas(ctx context.Context, req *pb.UpdateKelasRequest) (*pb.UpdateKelasResponse, error) {
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"SchemaName", "Kelas"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaName()
	Kelas := req.Kelas
	// rombelId := uuid.New()
	sekolahId, err := uuid.Parse(Kelas.SekolahId)
	if err != nil {
		fmt.Println("Error parsing UUID:", err)
	}
	ptkId, err := uuid.Parse(Kelas.PtkId)
	if err != nil {
		fmt.Println("Error parsing UUID:", err)
	}
	rombelId, err := uuid.Parse(Kelas.RombonganBelajarId)
	if err != nil {
		fmt.Println("Error parsing UUID:", err)
	}
	KelasModel := &models.RombonganBelajar{
		RombonganBelajarId:  rombelId,
		NmKelas:             Kelas.NmKelas,
		SekolahId:           sekolahId,
		SemesterId:          Kelas.SemesterId,
		JurusanId:           Kelas.JurusanId,
		TingkatPendidikanId: Kelas.TingkatPendidikanId,
		PtkID:               ptkId,
		JenisRombel:         Kelas.JenisRombel,
		NamaJurusanSp:       Kelas.NamaJurusanSp,
		JurusanSpId:         &uuid.Nil,
		KurikulumId:         Kelas.KurikulumId,
		// RombonganBelajarId:  kelas.RombonganBelajarId,
	}
	err = s.repo.Update(ctx, KelasModel, schemaName, "rombongan_belajar_id", Kelas.RombonganBelajarId)
	if err != nil {
		log.Printf("Gagal memperbaharui Kelas: %v", err)
		return nil, fmt.Errorf("gagal memperbaharui Kelas: %w", err)
	}

	// cek apakah anggota kelas berisi nilai
	anggotaKelas := req.GetAnggotaKelas()
	if len(anggotaKelas) > 0 {
		// var modelAnggotaKelas []models.RombelAnggota // Inisialisasi slice kosong

		for _, v := range anggotaKelas {
			newUUID := uuid.New()
			ced := models.RombelAnggota{ // Tidak perlu pakai pointer di sini
				AnggotaRombelId:    utils.StringToUUID(newUUID.String()),
				RombonganBelajarId: rombelId,
				PesertaDidikId:     utils.StringToUUID(v.PesertaDidikId),
				SemesterId:         KelasModel.SemesterId,
			}
			modelAnggota, err := s.repoRombelAnggota.FindByID(ctx, ced.PesertaDidikId.String(), schemaName, "peserta_didik_id")
			if err != nil {
				return nil, err
			}
			if modelAnggota == nil {
				results := s.repoRombelAnggota.Save(ctx, &ced, schemaName)
				if results != nil {
					return nil, err
				}
			} else {
				results := s.repoRombelAnggota.Update(ctx, &ced, schemaName, "peserta_didik_id", ced.PesertaDidikId.String())
				if results != nil {
					return nil, err
				}
			}
			// modelAnggotaKelas = append(modelAnggotaKelas, ced) // Tambahkan ke slice
			// modelAnggota, err := s.repoRombelAnggota.Update(ctx, ced,schemaName,"rombel_id",ced.AnggotaRombelId)
		}
	}
	return &pb.UpdateKelasResponse{
		Message: "Kelas berhasil diperbarui",
		Status:  true,
	}, nil
}

// **DeleteKelas**
func (s *RombelServiceServer) DeleteKelas(ctx context.Context, req *pb.DeleteKelasRequest) (*pb.DeleteKelasResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "KelasId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaName()
	KelasID := req.GetKelasId()

	err = s.repo.Delete(ctx, KelasID, schemaName, "rombongan_belajar_id")
	if err != nil {
		log.Printf("Gagal menghapus Kelas: %v", err)
		return nil, fmt.Errorf("gagal menghapus Kelas: %w", err)
	}

	return &pb.DeleteKelasResponse{
		Message: "Kelas berhasil dihapus",
		Status:  true,
	}, nil
}

// UploadKelas mengunggah data Kelas dari file Excel
// func (s *RombelServiceServer) UploadKelas(ctx context.Context, req *pb.UploadKelasRequest) (*pb.UploadKelasResponse, error) {
// 	schemaName := req.GetSchemaName()
// 	fileData := req.GetFile() // File dalam bentuk byte array

// 	// Simpan file ke sementara
// 	tempFile := "/tmp/uploaded_Kelas.xlsx"
// 	err := saveFile(tempFile, fileData)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal menyimpan file sementara: %w", err)
// 	}

// 	// Baca file Excel
// 	f, err := excelize.OpenFile(tempFile)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal membaca file Excel: %w", err)
// 	}
// 	defer f.Close()

// 	// Ambil semua data dari sheet pertama
// 	rows, err := f.GetRows(f.GetSheetName(0))
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal mengambil data dari sheet: %w", err)
// 	}

// 	// Pastikan ada data
// 	if len(rows) < 2 {
// 		return nil, fmt.Errorf("file Excel kosong atau tidak memiliki data yang valid")
// 	}

// 	// Validasi header
// 	expectedHeaders := []string{"NIS", "NISN", "NamaKelas", "TempatLahir", "TanggalLahir", "JenisKelamin", "Agama"}
// 	for i, expected := range expectedHeaders {
// 		if rows[0][i] != expected {
// 			return nil, fmt.Errorf("format kolom tidak sesuai, kolom '%s' seharusnya ada di posisi %d", expected, i+1)
// 		}
// 	}

// 	var KelasList []*models.RombonganBelajar

// 	// Mulai dari baris kedua karena baris pertama adalah header
// 	for _, row := range rows[1:] {
// 		if len(row) < len(expectedHeaders) {
// 			log.Println("Skipping row due to insufficient data:", row)
// 			continue
// 		}

// 		// Konversi data sesuai dengan model
// 		namaKelas := row[2]
// 		nis := row[0]
// 		nisn := row[1]
// 		tempatLahir := row[3]
// 		tanggalLahir := row[4]
// 		jenisKelamin := row[5]
// 		agama := row[6]

// 		// Validasi data
// 		if nis == "" || namaKelas == "" || nisn == "" {
// 			log.Println("Skipping row due to missing required fields:", row)
// 			continue
// 		}

// 		// Konversi angka
// 		nisInt, err := strconv.Atoi(nis)
// 		if err != nil {
// 			log.Printf("Format NIS tidak valid: %s", nis)
// 			continue
// 		}

// 		nisnInt, err := strconv.Atoi(nisn)
// 		if err != nil {
// 			log.Printf("Format NISN tidak valid: %s", nisn)
// 			continue
// 		}

// 		// Masukkan ke dalam list
// 		Kelas := &models.RombonganBelajar{
// 			NIS:          strconv.Itoa(nisInt),
// 			NISN:         strconv.Itoa(nisnInt),
// 			NamaKelas:    namaKelas,
// 			TempatLahir:  tempatLahir,
// 			TanggalLahir: tanggalLahir,
// 			JenisKelamin: jenisKelamin,
// 			Agama:        agama,
// 		}
// 		KelasList = append(KelasList, Kelas)
// 	}

// 	// Simpan ke database
// 	err = s.repo.BatchSave(ctx, KelasList, schemaName)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal menyimpan data Kelas ke database: %w", err)
// 	}

// 	return &pb.UploadKelasResponse{
// 		Message: "Kelas berhasil diunggah",
// 		Total:   int32(len(KelasList)),
// 		Status:  true,
// 	}, nil
// }
