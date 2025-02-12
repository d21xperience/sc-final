package services

import (
	"fmt"
	"sekolah/models"

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
			PesertaDidikID:  row[0],
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
		headers = []string{"NIS", "NISN", "NamaSiswa", "TempatLahir", "TanggalLahir", "JenisKelamin", "Agama"}
	case "nilai_akhir":
		headers = []string{"NIS", "NamaSiswa", "MataPelajaran", "NilaiAkhir"}
	case "guru":
		headers = []string{"NIP", "NamaGuru", "MataPelajaran", "Telepon"}
	case "kelas":
		headers = []string{"IDKelas", "NamaKelas", "WaliKelas"}
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

	// Simpan ke file
	err := f.SaveAs(filePath)
	if err != nil {
		return fmt.Errorf("gagal membuat template %s: %w", templateType, err)
	}

	return nil
}
