package services

import (
	"context"
	"fmt"
	"sekolah/models"
	"sekolah/repositories"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type ParamTemplate struct {
	templateType string
	schemaname   string
	filePath     string
	semesterId   string
}

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

// Fungsi untuk membaca file Excel dan memproses data berdasarkan jenis
func BacaDataExcel(param *ParamTemplate) ([][]string, error) {
	f, err := excelize.OpenFile(param.filePath)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca file Excel: %w", err)
	}
	defer f.Close()

	ret, err := f.GetDocProps()
	if err != nil {
		return nil, err
	}
	// if !strings.EqualFold(ret.Keywords, param.schemaname) || !strings.EqualFold(ret.ContentStatus, param.templateType) {
	// 	return nil, err
	// }
	param.semesterId = ret.Category
	param.schemaname = ret.Keywords
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data dari sheet: %w", err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("file Excel kosong atau tidak memiliki data yang valid")
	}
	// Kembalikan data mulai dari baris kedua (tanpa header)
	return rows[1:], nil

}

// Fungsi generik untuk membaca file Excel dan memproses data berdasarkan jenis
func UploadDataSekolah[T any](param ParamTemplate) ([]T, error) {
	f, err := excelize.OpenFile(param.filePath)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca file Excel: %w", err)
	}
	defer f.Close()

	ret, err := f.GetDocProps()
	if err != nil {
		return nil, err
	}
	// ContentStatus: param.templateType,
	// Category:      param.semesterId,
	// Keywords:      param.schemaname,
	if ret.Keywords != param.schemaname || ret.ContentStatus != param.templateType {
		return nil, err
	}

	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data dari sheet: %w", err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("file Excel kosong atau tidak memiliki data yang valid")
	}

	// Parsing data berdasarkan uploadType
	switch param.templateType {
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
	case "nilai_akhir":
		if v, ok := any(parseNilaiAkhir(rows)).([]T); ok {
			return v, nil
		}
	case "ijazah":
		if v, ok := any(parseNilaiAkhir(rows)).([]T); ok {
			return v, nil
		}
	default:
		return nil, fmt.Errorf("jenis unggahan tidak dikenali: %s", param.templateType)
	}

	return nil, fmt.Errorf("gagal memproses data dengan tipe yang diberikan")
}

// Fungsi parsing untuk siswa
func parseSiswa(rows [][]string) []models.PesertaDidik {
	var siswaList []models.PesertaDidik
	for _, row := range rows[1:] {
		if len(row) < 3 {
			continue
		}
		siswa := models.PesertaDidik{
			PesertaDidikId: row[0],
			Nis:            row[1],
			Nisn:           row[2],
			NmSiswa:        row[3],
			TempatLahir:    row[4],
			// TanggalLahir:    row[5],
			JenisKelamin: row[6],
			Agama:        row[7],
			AlamatSiswa:  &row[8],
			TeleponSiswa: row[9],
			// DiterimaTanggal: row[10],
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
func GenerateTemplate(param ParamTemplate, db *gorm.DB) error {
	f := excelize.NewFile()
	sheetName := "Template"
	f.SetSheetName("Sheet1", sheetName)

	// Tentukan header berdasarkan jenis template
	var headers []string
	templates := map[string][]string{
		"siswa": {"Nama", "NIPD", "JK", "NISN", "Tempat Lahir", "Tanggal Lahir", "NIK", "Agama", "Alamat", "RT", "RW", "Dusun", "Kelurahan", "Kecamatan", "Kode Pos", "Jenis Tinggal", "Alat Transportasi", "Telepon", "HP", "E-Mail", "SKHUN", "Penerima KPS", "No. KPS", "Nama Ayah", "Tahun Lahir Ayah", "Jenjang Pendidikan Ayah", "Pekerjaan Ayah", "Penghasilan Ayah", "NIK Ayah", "Nama Ibu", "Tahun Lahir Ibu", "Jenjang Pendidikan Ibu", "Pekerjaan Ibu", "Penghasilan Ibu", "NIK Ibu", "Nama Wali", "Tahun Lahir Wali", "Jenjang Pendidikan Wali", "Pekerjaan Wali", "Penghasilan Wali", "NIK Wali", "Rombel Saat Ini", "No Peserta Ujian Nasional", "No Seri Ijazah", "Penerima KIP", "Nomor KIP", "Nama di KIP", "Nomor KKS", "No Registrasi Akta Lahir", "Bank", "Nomor Rekening Bank", "Rekening Atas Nama", "Layak PIP (usulan dari sekolah)", "Alasan Layak PIP", "Kebutuhan Khusus", "Sekolah Asal", "Anak ke-berapa", "Lintang", "Bujur", "No KK", "Berat Badan", "Tinggi Badan", "Lingkar Kepala", "Jml. Saudara Kandung", "Jarak Rumah ke Sekolah (KM)"},

		"nilai_akhir": {"peserta_didik_id(uuid)", "nm_siswa", "mata_pelajaran_id"},
		"ijazah":      {"peserta_didik_id(uuid)", "nm_siswa", "nis", "nomor_ijazah", "tahun_lulus"},

		"kelas": {"Jenis Rombel", "Tingkat", "Nama Rombel", "Kurikulum", "Program/Kompetensi Keahlian", "Nama PTK", "NUPTK", "PTK Induk", "Kepegawaian", "Nama Matpel", "Kode Matpel", "JJM", "Jml Siswa", "Tgl SK Mengajar", "SK Mengajar", "Status di Kurikulum"},

		"guru": {"Nama", "NUPTK", "JK", "Tempat Lahir", "Tanggal Lahir", "NIP", "Status Kepegawaian", "Jenis PTK", "Agama", "Alamat Jalan", "RT", "RW", "Nama Dusun", "Desa/Kelurahan", "Kecamatan", "Kode Pos", "Telepon", "HP", "Email", "Tugas Tambahan", "SK CPNS", "Tanggal CPNS", "SK Pengangkatan", "TMT Pengangkatan", "Lembaga Pengangkatan", "Pangkat Golongan", "Sumber Gaji", "Nama Ibu Kandung", "Status Perkawinan", "Nama Suami/Istri", "NIP Suami/Istri", "Pekerjaan Suami/Istri", "TMT PNS", "Sudah Lisensi Kepala Sekolah", "Pernah Diklat Kepengawasan", "Keahlian Braille", "Keahlian Bahasa Isyarat", "NPWP", "Nama Wajib Pajak", "Kewarganegaraan", "Bank", "Nomor Rekening Bank", "Rekening Atas Nama", "NIK", "No KK", "Karpeg", "Karis/Karsu", "Lintang", "Bujur", "NUKS"},
	}
	headers, exists := templates[param.templateType]
	if !exists {
		return fmt.Errorf("template type %s tidak ditemukan", param.templateType)
	}

	for col, header := range headers {
		cell := fmt.Sprintf("%s1", getExcelColumnName(col))
		f.SetCellValue(sheetName, cell, header)
	}
	err := f.SetDocProps(&excelize.DocProperties{
		ContentStatus: param.templateType,
		Category:      param.semesterId,
		Keywords:      param.schemaname,
	})
	if err != nil {
		return fmt.Errorf("gagal membuat properties %s: %w", param.templateType, err)
	}

	// Jika templatetype adalah siswa
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// switch param.templateType {

	// case "nilai_akhir":
	// 	// err := templateNilaiAkhir(ctx, db, f, param, sheetName)
	// 	// if err != nil {
	// 	// 	return err
	// 	// }
	// 	// Tambahkan mata pelajaran
	// case "ijazah":
	// 	// err := templateIjazah(ctx, db, f, param, sheetName)
	// 	// if err != nil {
	// 	// 	return err
	// 	// }
	// 	// default:
	// 	// 	return fmt.Errorf("template type %s tidak ditemukan", param.templateType)
	// }

	// Simpan ke file
	err = f.SaveAs(param.filePath)
	if err != nil {
		return fmt.Errorf("gagal membuat template %s: %w", param.templateType, err)
	}

	return nil
}

func templateNilaiAkhir(ctx context.Context, db *gorm.DB, f *excelize.File, param ParamTemplate, sheetName string) error {
	// TODO: Implementasi template nilai akhir
	// sheetName := "Template"
	f.SetSheetName("Sheet1", sheetName)
	repoAnggotaKelas := repositories.NewRombelAnggotaRepository(db)
	joins := []string{
		"JOIN tabel_siswa ON tabel_siswa.peserta_didik_id = tabel_anggotakelas.peserta_didik_id",
		// "JOIN ref.jurusan ON tabel_kelas.jurusan_id = ref.jurusan.jurusan_id",
	}
	preloads := []string{"PesertaDidik"}
	var conditions = map[string]interface{}{
		"semester_id": param.semesterId,
	}
	groupByColumns := []string{"tabel_anggotakelas.anggota_rombel_id"} // Hindari duplikasi
	anggotaKelasModel, err := repoAnggotaKelas.FindWithPreloadAndJoins(ctx, param.schemaname, joins, preloads, conditions, groupByColumns, nil, false)
	if err != nil {
		return err
	}

	for row, pd := range anggotaKelasModel {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row+2), pd.PesertaDidik.PesertaDidikId)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row+2), pd.PesertaDidik.NmSiswa)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row+2), pd.SemesterId)
	}
	f.SetColVisible(sheetName, "A", false)
	return nil
}
func templateIjazah(ctx context.Context, db *gorm.DB, f *excelize.File, param ParamTemplate, sheetName string) error {
	// TODO: Implementasi template nilai akhir
	// sheetName := "Template"
	f.SetSheetName("Sheet1", sheetName)
	repoAnggotaKelas := repositories.NewRombelAnggotaRepository(db)
	joins := []string{
		"JOIN tabel_siswa ON tabel_siswa.peserta_didik_id = tabel_anggotakelas.peserta_didik_id",
		// "JOIN ref.jurusan ON tabel_kelas.jurusan_id = ref.jurusan.jurusan_id",
	}
	preloads := []string{"PesertaDidik"}
	var conditions = map[string]interface{}{
		"semester_id": param.semesterId,
	}
	groupByColumns := []string{"tabel_anggotakelas.anggota_rombel_id"} // Hindari duplikasi
	anggotaKelasModel, err := repoAnggotaKelas.FindWithPreloadAndJoins(ctx, param.schemaname, joins, preloads, conditions, groupByColumns, nil, false)
	if err != nil {
		return err
	}

	for row, pd := range anggotaKelasModel {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row+2), pd.PesertaDidik.PesertaDidikId)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row+2), pd.PesertaDidik.NmSiswa)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row+2), pd.SemesterId)
	}
	f.SetColVisible(sheetName, "A", false)
	return nil
}

func getExcelColumnName(col int) string {
	result := ""
	for col >= 0 {
		result = string(rune('A'+(col%26))) + result
		col = col/26 - 1
	}
	return result
}

// ============================
// s header ke Excel
// for i, header := range headers {
// 	col := string(rune('A'+i)) + "1"
// 	f.SetCellValue(sheetName, col, header)
// }
// f.SetCellValue(sheetName, "Y1", "Timestamp")
// f.SetCellValue(sheetName, "Y2", time.Now().Format("2006-01-02 15:04:05")) // Isi dengan waktu saat ini

// f.SetCellValue(sheetName, "Z1", "UserID")
// f.SetCellValue(sheetName, "Z2", "123456") // Bisa diisi dengan ID pengguna yang mengunduh template

// f.SetColVisible(sheetName, "Y", false) // Sembunyikan kolom Timestamp
// f.SetColVisible(sheetName, "Z", false) // Sembunyikan kolom UserID

// sampleData := []interface{}{"12345", "987654321", "Budi Santoso", "Jakarta", "2005-08-10", "Laki-laki", "Islam"}
// for i, data := range sampleData {
// 	col := string(rune('A'+i)) + "2"
// 	f.SetCellValue(sheetName, col, data)
// }

// // Buat validasi dropdown untuk JenisKelamin
// dv := excelize.NewDataValidation(true)
// dv.Sqref = "F2:F1000" // Rentang sel yang divalidasi
// dv.SetDropList([]string{"L", "P"})

// if err := f.AddDataValidation(sheetName, dv); err != nil {
// 	return fmt.Errorf("gagal menambahkan validasi: %w", err)
// }

// // Validasi NilaiAkhir hanya angka 0-100
// dvNilai := excelize.NewDataValidation(true)
// dvNilai.Sqref = "D2:D100"
// // dvNilai.SetWholeNumber(0, 100)

// if err := f.AddDataValidation(sheetName, dvNilai); err != nil {
// 	return fmt.Errorf("gagal menambahkan validasi angka: %w", err)
// }

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

// f.NewSheet("Panduan")
// f.SetCellValue("Panduan", "A1", "Panduan Pengisian Template")
// f.SetCellValue("Panduan", "A2", "1. Isi semua kolom sesuai dengan contoh yang diberikan.")
// f.SetCellValue("Panduan", "A3", "2. Gunakan dropdown untuk memilih data yang tersedia.")
// f.SetCellValue("Panduan", "A4", "3. Pastikan semua data terisi sebelum mengunggah ke sistem.")

// ============================

// func parseData[T any](rows [][]string, parser func([]string) *T) []*T {
// 	var dataList []*T

// 	// Pastikan ada lebih dari satu baris (karena baris pertama biasanya header)
// 	if len(rows) < 2 {
// 		return dataList
// 	}

// 	for _, row := range rows[1:] { // Lewati header
// 		if len(row) == 0 {
// 			continue
// 		}

// 		item := parser(row)
// 		if item != nil {
// 			dataList = append(dataList, item)
// 		}
// 	}

// 	return dataList
// }
