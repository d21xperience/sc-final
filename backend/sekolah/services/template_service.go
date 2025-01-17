package services

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

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
