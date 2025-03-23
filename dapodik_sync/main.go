// package main

// import (
// 	"log"

// 	"gioui.org/app"
// 	"gioui.org/layout"
// 	"gioui.org/op"
// 	"gioui.org/widget"
// 	"gioui.org/widget/material"
// )

// // Token autentikasi
// const bearerToken = "evwRIGsdX3bzBa8"

// func main() {
// 	var err error
// 	// Load konfigurasi
// 	cfg := config.LoadConfig()
// 	// Inisialisasi database
// 	config.InitDatabase(cfg)
// 	// Auto migrate tabel jika belum ada
// 	config.DB.AutoMigrate(&models.Semester{})
// 	config.DB.AutoMigrate(&models.User{})
// 	// konek ke grpc server
// 	config.ConnectGRPCServer(cfg)
// 	// // Load konfigurasi dari .env
// 	// err = godotenv.Load()
// 	// if err != nil {
// 	// 	log.Fatalf("Gagal memuat file .env: %v", err)
// 	// }

// 	// cfg.BaseURL := os.Getenv("LOCAL_HOST")
// 	// baseURL := os.Getenv("BASE_URL")
// 	// npsn := os.Getenv("NPSN")

// 	semesters, err := scraper.ScrapeSemesters(cfg.BaseURL)
// 	if err != nil {
// 		log.Fatalf("Gagal mengambil data semester: %v", err)
// 	}

// 	semesterID := os.Getenv("SEMESTER_ID")
// 	// Tampilkan hasil
// 	fmt.Println("Daftar Semester yang Ditemukan:")

// 	// Definisi URL dengan parameter dinamis
// 	urls := map[string]string{
// 		"sekolah": fmt.Sprintf("%s/getSekolah?npsn=%s", cfg.BaseURL, npsn),
// 		// "semester": fmt.Sprintf("%s/getRombonganBelajar?npsn=%s", cfg.BaseURL, npsn),
// 		"mapel": fmt.Sprintf("%s/getMataPelajaran?npsn=%s&semester_id=%s", cfg.BaseURL, npsn, semesterID),
// 		"gtk":   fmt.Sprintf("%s/getGtk?npsn=%s&semester_id=%s", cfg.BaseURL, npsn, semesterID),
// 		// "siswa":  fmt.Sprintf("%s/getPesertaDidik?npsn=%s&semester_id=%s", cfg.BaseURL, npsn, semesterID),
// 		"siswa":  "",
// 		"rombel": fmt.Sprintf("%s/getRombonganBelajar?npsn=%s&semester_id=%s", cfg.BaseURL, npsn, semesterID),
// 	}
// 	var jsonData []byte
// 	var data map[string]interface{}
// 	for key, url := range urls {
// 		// data, err := scraper.FetchData(url, bearerToken)
// 		// if err != nil {
// 		// 	log.Printf("Gagal mengambil data %s: %v", key, err)
// 		// 	continue
// 		// }
// 		switch key {
// 		case "siswa":
// 			for _, semester := range semesters {
// 				fmt.Println(semester)
// 				urls["siswa"] = fmt.Sprintf("%s/getPesertaDidik?npsn=%s&semester_id=%s", cfg.BaseURL, npsn, semester)
// 				data, err = scraper.FetchData(url, bearerToken)
// 				if err != nil {
// 					log.Printf("Gagal mengambil data %s: %v", key, err)
// 					continue
// 				}
// 				jsonData, err = json.MarshalIndent(data, "", "  ") // Memformat JSON agar lebih terbaca
// 				if err != nil {
// 					log.Printf("Gagal mengonversi data %s ke JSON: %v", key, err)
// 					continue
// 				}
// 			}
// 		case "rombel":
// 			for _, semester := range semesters {
// 				fmt.Println(semester)
// 				urls["rombel"] = fmt.Sprintf("%s/getRombonganBelajar?npsn=%s&semester_id=%s", cfg.BaseURL, npsn, semesterID)

// 			}

// 		}
// 		// Menampilkan data hasil parsing ke JSON
// 		jsonData, err = json.MarshalIndent(data, "", "  ") // Memformat JSON agar lebih terbaca
// 		if err != nil {
// 			log.Printf("Gagal mengonversi data %s ke JSON: %v", key, err)
// 			continue
// 		}

//			fmt.Printf("\n[DEBUG] Data JSON dari %s:\n%s\n", key, string(jsonData))
//		}
//	}
// func main() {
// 	go func() {
// 		window := new(app.Window)
// 		err := run(window)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		os.Exit(0)
// 	}()
// 	app.Main()
// }

// func run(window *app.Window) error {
// 	theme := material.NewTheme()
// 	var ops op.Ops
// 	for {
// 		switch e := window.Event().(type) {
// 		case app.DestroyEvent:
// 			return e.Err
// 		case app.FrameEvent:
// 			// This graphics context is used for managing the rendering state.
// 			gtx := app.NewContext(&ops, e)

// 			// Define an large label with an appropriate text:
// 			title := material.H1(theme, "Hello, Gio")

// 			// Change the color of the label.
// 			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
// 			title.Color = maroon

// 			// Change the position of the label.
// 			title.Alignment = text.Middle

// 			// Draw the label to the graphics context.
// 			title.Layout(gtx)

//				// Pass the drawing operations to the GPU.
//				e.Frame(gtx.Ops)
//			}
//		}
//	}
package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := new(app.Window) // Membuat window baru
		w.Option(app.Title("Dapodik_sync"))
		w.Option(app.MaxSize(unit.Dp(400), unit.Dp(600)))
		w.Option(app.MinSize(unit.Dp(400), unit.Dp(600)))
		

		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	var ops op.Ops
	th := material.NewTheme()
	// Widgets
	signInBtn := new(widget.Clickable)
	registerLink := new(widget.Clickable)
	forgotPasswordCheckbox := new(widget.Bool)
	usernameField := new(widget.Editor)
	passwordField := new(widget.Editor)

	for {
		switch e := w.Event().(type) { 
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceEvenly,
			}.Layout(gtx,
				// Judul "Sign In"
				layout.Rigid(material.H1(th, "Sign In").Layout),

				// Tautan "Don't have an account? Register"
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis: layout.Horizontal,
					}.Layout(gtx, layout.Rigid(material.Body1(th, "Don't have an account?").Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Clickable(gtx, registerLink, func(gtx layout.Context) layout.Dimensions {
								return material.Body1(th, "Register").Layout(gtx)
							})
						}),
					)
				}),

				// Field "Username or Email"
				layout.Rigid(material.Editor(th, usernameField, "Username or Email").Layout),

				// Field "Password"
				layout.Rigid(material.Editor(th, passwordField, "Password").Layout),

				// Checkbox "Forgot Password?"
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis: layout.Horizontal,
					}.Layout(gtx,
						layout.Rigid(material.CheckBox(th, forgotPasswordCheckbox, "Forgot Password?").Layout),
					)
				}),

				// Tombol "ACCESS MY ACCOUNT"
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, signInBtn, "ACCESS MY ACCOUNT").Layout(gtx)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
