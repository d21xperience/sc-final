package services

import (
	"fmt"
	"sekolah/models"
	"time"

	"github.com/xuri/excelize/v2"
)

// // RequestRequirement digunakan untuk menyimpan dependensi service dan protobuf
// type RequestRequirement[S any, P any] struct {
// 	service  S // Service untuk mengakses data
// 	protoBuf P // Protobuf atau struct untuk representasi data
// }

// func (r *RequestRequirement[S, P]) GetModelByID(ctx context.Context, id string, fetcher func(S, string) (*P, error)) (*P, error) {
// 	// Gunakan fetcher untuk mengambil data berdasarkan ID
// 	result, err := fetcher(r.service, id)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get model by ID: %w", err)
// 	}

// 	return result, nil
// }

func ConvertModelsToPB[T any, U any](models []*T, converter func(*T) *U) []*U {
	var pbList []*U
	for _, model := range models {
		pbList = append(pbList, converter(model))
	}
	return pbList
}
func ConvertPBToModels[T any, U any](pbs []*T, converter func(*T) *U) []*U {
	var modelList []*U
	for _, model := range pbs {
		modelList = append(modelList, converter(model))
	}
	return modelList
}

func ConvertModelToPB[T any, U any](model *T, converter func(*T) *U) *U {
	if model == nil {
		return nil
	}
	return converter(model)
}

// Fungsi generik untuk membaca file Excel dan memproses data berdasarkan jenis
func UploadDataSekolah[T any](filePath, uploadType string) ([]T, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca file Excel: %w", err)
	}
	defer f.Close()

	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data dari sheet: %w", err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("file Excel kosong atau tidak memiliki data yang valid")
	}

	// Parsing data berdasarkan uploadType
	switch uploadType {
	case "siswa":
		if v, ok := any(parseSiswa(rows)).([]T); ok {
			return v, nil
		}
	case "guru":
		if v, ok := any(parseGuru(rows)).([]T); ok {
			return v, nil
		}
	case "kelas":
		if v, ok := any(parseKelas(rows)).([]T); ok {
			return v, nil
		}
	case "nilaiAkhir":
		if v, ok := any(parseNilaiAkhir(rows)).([]T); ok {
			return v, nil
		}
	default:
		return nil, fmt.Errorf("jenis unggahan tidak dikenali: %s", uploadType)
	}

	return nil, fmt.Errorf("gagal memproses data dengan tipe yang diberikan")
}

// Fungsi parsing untuk siswa
func parseSiswa(rows [][]string) []*models.PesertaDidik {
	var siswaList []*models.PesertaDidik
	for _, row := range rows[1:] {
		if len(row) < 3 {
			continue
		}
		siswa := &models.PesertaDidik{
			PesertaDidikId:  row[0],
			Nis:             row[1],
			Nisn:            row[2],
			NmSiswa:         row[3],
			TempatLahir:     row[4],
			TanggalLahir:    row[5],
			JenisKelamin:    row[6],
			Agama:           row[7],
			AlamatSiswa:     &row[8],
			TeleponSiswa:    row[9],
			DiterimaTanggal: row[10],
			// Umur:   parseInt(row[1]), // Fungsi parseInt bisa digunakan untuk mengubah string ke int
			// Alamat: row[2],
		}
		siswaList = append(siswaList, siswa)
	}
	return siswaList
}

// Fungsi parsing untuk guru
func parseGuru(rows [][]string) []*models.TabelPTK {
	var guruList []*models.TabelPTK
	for _, row := range rows[1:] {
		if len(row) < 3 {
			continue
		}
		guru := &models.TabelPTK{
			// Nama:   row[0],
			// Mapel:  row[1],
			// Alamat: row[2],
		}
		guruList = append(guruList, guru)
	}
	return guruList
}

// Fungsi parsing untuk kelas
func parseKelas(rows [][]string) []*models.RombonganBelajar {
	var kelasList []*models.RombonganBelajar
	for _, row := range rows[1:] {
		if len(row) < 2 {
			continue
		}
		kelas := &models.RombonganBelajar{
			// NamaKelas:  row[0],
			// JumlahSiswa: parseInt(row[1]),
		}
		kelasList = append(kelasList, kelas)
	}
	return kelasList
}

// Fungsi parsing untuk kelas
func parseNilaiAkhir(rows [][]string) []*models.NilaiAkhir {
	var kelasList []*models.NilaiAkhir
	for _, row := range rows[1:] {
		if len(row) < 2 {
			continue
		}
		kelas := &models.NilaiAkhir{
			// AnggotaRombelID: (uuid.UUID).row[0],
			// NamaKelas:  row[0],
			// JumlahSiswa: parseInt(row[1]),
			// NilaiPeng: int32(parseInt(row[2])),
		}
		kelasList = append(kelasList, kelas)
	}
	return kelasList
}

// GenerateTemplate membuat template XLSX untuk berbagai jenis data
func GenerateTemplate(templateType, filePath string) error {
	f := excelize.NewFile()
	sheetName := "Template"
	f.SetSheetName("Sheet1", sheetName)

	// Tentukan header berdasarkan jenis template
	var headers []string

	switch templateType {
	case "siswa":
		headers = []string{"peserta_didik_id", "nis", "nisn", "nm_siswa", "tempat_lahir", "tanggal_lahir", "jenis_kelamin", "agama", "alamat_siswa", "telepon_siswa", "diterima_tanggal", "nm_ayah", "nm_ibu", "pekerjaan_ayah", "pekerjaan_ibu", "nm_wali"}
	case "siswa_pelengkap":
		headers = []string{"pelengkap_siswa_id", "peserta_didik_id", "status_dalam_kel", "anak_ke", "sekolah_asal", "diterima_kelas", "alamat_ortu", "telepon_ortu", "alamat_wali", "telepon_wali", "foto_siswa"}
	case "nilai_akhir":
		headers = []string{"NIS", "NamaSiswa", "MataPelajaran", "NilaiAkhir"}
	case "guru":
		headers = []string{"ptk_id", "nama", "nip", "jenis_ptk_id", "jenis_kelamin", "tempat_lahir", "tanggal_lahir", "nuptk", "alamat_jalan", "status_keaktifan_id"}
	case "kelas":
		headers = []string{"rombongan_belajar_id", "sekolah_id", "semester_id", "jurusan_id", "ptk_id", "nm_kelas", "tingkat_pendidikan_id", "jenis_rombel", "nama_jurusan_sp", "jurusan_sp_id", "kurikulum_id"}
	case "ijazah":
		headers = []string{"NIS", "NamaSiswa", "NomorIjazah", "TahunLulus"}
	default:
		return fmt.Errorf("jenis template tidak dikenali: %s", templateType)
	}

	// Tulis header ke Excel
	for i, header := range headers {
		col := string(rune('A'+i)) + "1"
		f.SetCellValue(sheetName, col, header)
	}
	f.SetCellValue(sheetName, "Y1", "Timestamp")
	f.SetCellValue(sheetName, "Y2", time.Now().Format("2006-01-02 15:04:05")) // Isi dengan waktu saat ini

	f.SetCellValue(sheetName, "Z1", "UserID")
	f.SetCellValue(sheetName, "Z2", "123456") // Bisa diisi dengan ID pengguna yang mengunduh template

	f.SetColVisible(sheetName, "Y", false) // Sembunyikan kolom Timestamp
	f.SetColVisible(sheetName, "Z", false) // Sembunyikan kolom UserID

	sampleData := []interface{}{"12345", "987654321", "Budi Santoso", "Jakarta", "2005-08-10", "Laki-laki", "Islam"}
	for i, data := range sampleData {
		col := string(rune('A'+i)) + "2"
		f.SetCellValue(sheetName, col, data)
	}

	// Buat validasi dropdown untuk JenisKelamin
	dv := excelize.NewDataValidation(true)
	dv.Sqref = "F2:F1000" // Rentang sel yang divalidasi
	dv.SetDropList([]string{"L", "P"})

	if err := f.AddDataValidation(sheetName, dv); err != nil {
		return fmt.Errorf("gagal menambahkan validasi: %w", err)
	}

	// Validasi NilaiAkhir hanya angka 0-100
	dvNilai := excelize.NewDataValidation(true)
	dvNilai.Sqref = "D2:D100"
	// dvNilai.SetWholeNumber(0, 100)

	if err := f.AddDataValidation(sheetName, dvNilai); err != nil {
		return fmt.Errorf("gagal menambahkan validasi angka: %w", err)
	}

	// err := f.ProtectSheet(sheetName, &excelize.SheetProtectionOptions{
	// 	FormatCells:      false, // Tidak bisa ubah format
	// 	FormatColumns:    false, // Tidak bisa ubah kolom
	// 	FormatRows:       false, // Tidak bisa ubah baris
	// 	InsertRows:       true,  // Bisa menambah baris baru
	// 	InsertColumns:    false, // Tidak bisa menambah kolom baru
	// 	InsertHyperlinks: false,
	// 	DeleteRows:       false, // Tidak bisa hapus baris
	// 	DeleteColumns:    false, // Tidak bisa hapus kolom
	// 	// SelectLockedCells:   false,
	// 	// SelectUnlockedCells: true,
	// })
	// if err != nil {
	// 	return err
	// }

	f.NewSheet("Panduan")
	f.SetCellValue("Panduan", "A1", "Panduan Pengisian Template")
	f.SetCellValue("Panduan", "A2", "1. Isi semua kolom sesuai dengan contoh yang diberikan.")
	f.SetCellValue("Panduan", "A3", "2. Gunakan dropdown untuk memilih data yang tersedia.")
	f.SetCellValue("Panduan", "A4", "3. Pastikan semua data terisi sebelum mengunggah ke sistem.")

	// Simpan ke file
	err := f.SaveAs(filePath)
	if err != nil {
		return fmt.Errorf("gagal membuat template %s: %w", templateType, err)
	}

	return nil
}
