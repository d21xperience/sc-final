package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	pb "sc-service/generated"

	shell "github.com/ipfs/go-ipfs-api"
)

// UploadService menangani penyimpanan file yang diunggah
type UploadServiceServer struct {
	pb.UnimplementedUploadDataSekolahServiceServer
	uploadDir string
	// repoSiswa          repositories.GenericRepository[models.PesertaDidik]
	// repoSiswaPelengkap repositories.GenericRepository[models.PesertaDidikPelengkap]
	// repoKelas          repositories.GenericRepository[models.RombonganBelajar]
	// repoKelasAnggota   repositories.GenericRepository[models.RombelAnggota]
	// repoGuru           repositories.GenericRepository[models.TabelPTK]
	// repoGuruTerdaftar  repositories.GenericRepository[models.PTKTerdaftar]
}

func NewUploadServiceServer() *UploadServiceServer {
	// repoSiswa := repositories.NewSiswaRepository(config.DB)
	// repoSiswaPelengkap := repositories.NewSiswaPelengkapRepository(config.DB)
	// repoKelas := repositories.NewrombonganBelajarRepository(config.DB)
	// repoKelasAnggota := repositories.NewRombelAnggotaRepository(config.DB)
	// repoGuru := repositories.NewPTKRepository(config.DB)
	// repoGuruTerdaftar := repositories.NewPTKTerdaftarRepository(config.DB)
	// if repoSiswa == nil {
	// 	log.Fatal("❌ ERROR: Gagal menginisialisasi repoSiswa") // Debugging
	// }
	return &UploadServiceServer{
		uploadDir: "uploads",
		// repoSiswa:          *repoSiswa,
		// repoSiswaPelengkap: *repoSiswaPelengkap,
		// repoKelas:          *repoKelas,
		// repoKelasAnggota:   *repoKelasAnggota,
		// repoGuru:           *repoGuru,
		// repoGuruTerdaftar:  *repoGuruTerdaftar,
	}
}

// UploadFile menangani upload melalui gRPC
func (s *UploadServiceServer) UploadDataSekolah(ctx context.Context, req *pb.UploadDataSekolahRequest) (*pb.UploadDataSekolahResponse, error) {
	// Tentukan lokasi penyimpanan file
	filePath := filepath.Join(s.uploadDir, req.Filename)

	// Simpan file yang dikirim dalam bytes
	err := os.WriteFile(filePath, req.File, 0644)
	if err != nil {
		return nil, fmt.Errorf("gagal menyimpan file: %w", err)
	}

	// Kembalikan URL file yang diunggah
	return &pb.UploadDataSekolahResponse{
		Message: "File berhasil diunggah",
		FileUrl: "/storage/uploads/" + req.Filename,
	}, nil
}

// UploadFileHTTP menangani upload file melalui REST API dengan multipart/form-data
func (s *UploadServiceServer) UploadFileHTTP(w http.ResponseWriter, r *http.Request) {
	// Hubungkan ke daemon IPFS lokal (API port 5001)
	sh := shell.NewShell("localhost:5001")

	// Buka file ijazah
	file, err := os.Open("ijazah.pdf")
	if err != nil {
		fmt.Println("Gagal membuka file:", err)
		return
	}
	defer file.Close()

	// Upload ke IPFS
	cid, err := sh.Add(file)
	if err != nil {
		fmt.Println("Gagal upload:", err)
		return
	}

	fmt.Printf("Ijazah berhasil diupload! CID: %s\n", cid)
	fmt.Printf("Akses via gateway: https://ipfs.io/ipfs/%s\n", cid)

	// Berikan respon
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "File berhasil diproses",
		// "data":    data,
	})
}

// GetTemplate menyediakan template Excel berdasarkan jenis data
// func (s *UploadServiceServer) DownloadDataSekolah(ctx context.Context, req *pb.DownloadDataSekolahRequest) (*pb.DownloadDataSekolahResponse, error) {
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"TemplateType"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	templateType := req.GetDownloadType()
// 	templatePath := fmt.Sprintf("templates/template_%s.xlsx", templateType)
// 	var param = ParamTemplate{
// 		schemaname:   "",
// 		filePath:     "",
// 		semesterId:   "",
// 		templateType: templateType,
// 	}
// 	// Buat file template jika belum ada
// 	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
// 		err := GenerateTemplate(param, config.DB)
// 		if err != nil {
// 			return nil, fmt.Errorf("gagal membuat template %s: %w", templateType, err)
// 		}
// 	}

// 	// Baca file template
// 	data, err := os.ReadFile(templatePath)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal membaca template %s: %w", templateType, err)
// 	}

// 	return &pb.DownloadDataSekolahResponse{
// 		Filename: fmt.Sprintf("template_%s.xlsx", templateType),
// 		File:     data,
// 	}, nil
// }

// HandleDownloadTemplate adalah handler untuk mengunduh file template .xlsx.
func (h *UploadServiceServer) DownloadTemplateHTTP(w http.ResponseWriter, r *http.Request) {
	// Ambil nama file dari query parameter
	templateType := r.URL.Query().Get("template_type")
	if templateType == "" {
		http.Error(w, "template-type is required", http.StatusBadRequest)
		return
	}
	// // Lokasi direktori template
	// var param = ParamTemplate{
	// 	schemaname:   r.FormValue("Schemaname"),
	// 	semesterId:   r.FormValue("semesterId"),
	// 	templateType: templateType,
	// }
	// templatePath := fmt.Sprintf("templates/template_%s_%s_%s.xlsx", templateType, param.schemaname, param.semesterId)
	// param.filePath = templatePath
	// // Buat file template jika belum ada
	// if _, err := os.Stat(templatePath); os.IsNotExist(err) {
	// 	err := GenerateTemplate(param, config.DB)
	// 	if err != nil {
	// 		http.Error(w, fmt.Sprintf("Gagal membuat template: %v", err), http.StatusInternalServerError)
	// 		return
	// 	}
	// }

	// // // Buka file template
	// file, err := os.Open(templatePath)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Gagal membuka file template: %v", err), http.StatusInternalServerError)
	// 	return
	// }
	// defer file.Close()

	// // // Mendapatkan informasi file untuk header
	// fileInfo, err := file.Stat()
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Gagal mendapatkan informasi file: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"; filename*=UTF-8''%s", fileInfo.Name(), url.QueryEscape(fileInfo.Name())))
	// w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileInfo.Name()))

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet") // MIME type untuk Excel
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "no-cache")

	// Kirim file
	// if _, err := io.Copy(w, file); err != nil {
	// 	http.Error(w, fmt.Sprintf("Gagal mengirim file: %v", err), http.StatusInternalServerError)
	// 	return
	// }

}

// Definisikan fungsi untuk menangani upload berdasarkan tipe data
// func processUpload[T any](
// 	ctx context.Context,
// 	param ParamTemplate,
// 	uploadFunc func(ParamTemplate) ([]T, error),
// 	saveFunc func(context.Context, *T, string) error,
// ) error {
// 	// Ambil data dari file
// 	data, err := uploadFunc(param)
// 	if err != nil {
// 		return fmt.Errorf("gagal memproses file: %v", err)
// 	}

// 	// Iterasi data dan simpan ke database
// 	for i := range data {
// 		utils.HandleNilPointers(&data[i]) // Hindari pointer nil

//			// Simpan ke database
//			if err := saveFunc(ctx, &data[i], param.schemaname); err != nil {
//				return fmt.Errorf("gagal menyimpan data: %v", err)
//			}
//		}
//		return nil
//	}
// func (s *UploadServiceServer) processUploadSiswa(
// 	ctx context.Context,
// 	param ParamTemplate,
// 	uploadFunc func(*ParamTemplate) ([][]string, error),
// ) error {
// 	// Ambil data dari file
// 	data, err := uploadFunc(&param)
// 	if err != nil {
// 		return fmt.Errorf("gagal memproses file: %v", err)
// 	}

// 	// Iterasi data dan simpan ke database
// 	for i := range data {
// 		// utils.HandleNilPointers(&data[i]) // Hindari pointer nil

// 		// Simpan ke database
// 		// database tabel_siswa
// 		pesertaDidikId := uuid.New()
// 		pelengkapSiswaId := uuid.New()
// 		anggotaRombelId := uuid.New()
// 		alamatSiswa := fmt.Sprintf("%s RT.%s RW.%s, Desa %s Kec. %s %s", data[i][8], data[i][9], data[i][10], data[i][12], data[i][13], data[i][14])

// 		// Validasi tanggal lahir

// 		var tanggalLahir time.Time
// 		tanggalLahirStr := data[i][5]
// 		if tanggalLahirStr != "" {
// 			tanggalLahir, err = time.Parse("2006-01-02", tanggalLahirStr)
// 			if err == nil {
// 				// tanggalLahir = &tanggalLahir
// 			}
// 		}

// 		// Validasi diterima_tanggal
// 		// diterimaTanggalStr := data[i][5]
// 		// if diterimaTanggalStr != "" {
// 		// 	parsedDate, err := time.Parse("2006-01-02", diterimaTanggalStr)
// 		// 	if err == nil {
// 		// 		siswa.DiterimaTgl = &parsedDate
// 		// 	}
// 		// }

// 		err := s.repoSiswa.Save(ctx, &models.PesertaDidik{
// 			PesertaDidikId: pesertaDidikId.String(),
// 			NmSiswa:        data[i][0],
// 			Nis:            data[i][1],
// 			JenisKelamin:   data[i][2],
// 			Nisn:           data[i][3],
// 			TempatLahir:    data[i][4],
// 			TanggalLahir:   &tanggalLahir,
// 			Agama:          data[i][7],
// 			AlamatSiswa:    &alamatSiswa,
// 			TeleponSiswa:   data[i][18],
// 			// DiterimaTanggal: data[i][1],
// 			NmAyah:        data[i][23],
// 			PekerjaanAyah: data[i][26],
// 			NmIbu:         data[i][29],
// 			PekerjaanIbu:  data[i][32],
// 			NmWali:        &data[i][39],
// 			PekerjaanWali: &data[i][41],
// 			// DiterimaTanggal: ,
// 		}, param.schemaname)
// 		if err != nil {
// 			return err
// 		}
// 		err = s.repoSiswaPelengkap.Save(ctx, &models.PesertaDidikPelengkap{
// 			PelengkapSiswaId: pelengkapSiswaId.String(),
// 			PesertaDidikId:   func(s string) *string { return &s }(pesertaDidikId.String()),
// 			SekolahAsal:      data[i][55],
// 			AnakKe:           &data[i][56],

// 			// FotoSiswa: ,
// 		}, param.schemaname)
// 		if err != nil {
// 			return err
// 		}

// 		// database tabel_anggotakelas
// 		// Jika kelas kosong abaikan
// 		if data[i][41] != "" {
// 			err = s.repoKelasAnggota.Save(ctx, &models.RombelAnggota{
// 				AnggotaRombelId: utils.StringToUUID(anggotaRombelId.String()),
// 				PesertaDidikId:  utils.StringToUUID(pesertaDidikId.String()),
// 				// SemesterId: ,
// 				// SemesterId: ,
// 				RombonganBelajar: models.RombonganBelajar{
// 					NmKelas: data[i][41],
// 				},
// 			}, param.schemaname)
// 			if err != nil {
// 				return err
// 			}
// 			// if err := saveFunc(ctx, &data[i], param.schemaname); err != nil {
// 			// 	return fmt.Errorf("gagal menyimpan data: %v", err)
// 			// }
// 		}

// 	}
// 	return nil
// }

// func (s *UploadServiceServer) processUploadGuru(
// 	ctx context.Context,
// 	param ParamTemplate,
// 	uploadFunc func(*ParamTemplate) ([][]string, error),
// ) error {
// 	// Ambil data dari file
// 	data, err := uploadFunc(&param)
// 	if err != nil {
// 		return fmt.Errorf("gagal memproses file: %v", err)
// 	}

// 	// Iterasi data dan simpan ke database
// 	for i := range data {
// 		// utils.HandleNilPointers(&data[i]) // Hindari pointer nil

// 		// Simpan ke database
// 		// database tabel_siswa
// 		ptkId := uuid.New()
// 		ptkTerdaftarId := uuid.New()
// 		// anggotaRombelId := uuid.New()
// 		alamatPtk := fmt.Sprintf("%s RT.%s RW.%s, Desa %s Kec. %s %s", data[i][10], data[i][11], data[i][12], data[i][14], data[i][15], data[i][16])

// 		// Validasi tanggal lahir

// 		var tanggalLahir time.Time
// 		tanggalLahirStr := data[i][5]
// 		if tanggalLahirStr != "" {
// 			tanggalLahir, err = time.Parse("2006-01-02", tanggalLahirStr)
// 			if err == nil {
// 				// tanggalLahir = &tanggalLahir
// 			}
// 		}

// 		err := s.repoGuru.Save(ctx, &models.TabelPTK{
// 			PtkID:        ptkId,
// 			Nama:         data[i][1],
// 			NUPTK:        &data[i][2],
// 			JenisKelamin: data[i][3],
// 			TempatLahir:  data[i][4],
// 			TanggalLahir: &tanggalLahir,
// 			NIP:          &data[i][6],
// 			AlamatJalan:  alamatPtk,

// 			// JenisPtkID:        data[i][0],
// 			// StatusKeaktifanID: data[i][0],
// 			// DiterimaTanggal: ,

// 		}, param.schemaname)
// 		if err != nil {
// 			return err
// 		}
// 		err = s.repoGuruTerdaftar.Save(ctx, &models.PTKTerdaftar{
// 			PtkTerdaftarId: ptkTerdaftarId,
// 			PtkID:          ptkId,
// 			TahunAjaranId:  param.semesterId[:4],
// 			// FotoSiswa: ,
// 		}, param.schemaname)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
