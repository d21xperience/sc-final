package services

// DownloadTemplateExcel mengembalikan file template Excel untuk import kelas
// func (s *RombelServiceServer) DownloadTemplateExcel(w http.ResponseWriter, r *http.Request) {
// 	// Path file template Excel
// 	templatePath := "templates/template_kelas.xlsx"

// 	// Set header untuk response sebagai file download
// 	w.Header().Set("Content-Disposition", "attachment; filename=template_kelas.xlsx")
// 	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

// 	// Baca file template
// 	http.ServeFile(w, r, templatePath)
// }