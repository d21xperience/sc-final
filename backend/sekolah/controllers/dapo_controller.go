package controllers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
)

type DataDapo struct {
	Scheme      string
	BaseURL     string
	Port        string
	FullPathUrl string
	// Username  string
	// Password  string
	SemeterID []string
	SekolahID string
	Token     string
}

type DapodikClient struct {
	HTTPClient *http.Client
}

var setDataDapo DataDapo

func NewDapodikClient(dapoServiceimpl *http.Client) *DapodikClient {
	return &DapodikClient{HTTPClient: dapoServiceimpl}
}

// Mendapatkan aplikasi dapodik
func (ds DapodikClient) GetDapodikApp(ctx *gin.Context) {
	fmt.Println("ok")
	type reVal struct {
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Status  int      `json:"status"`
		Data    []string `json:"data"`
	}
	var re reVal //Variabel untuk mengembalikan JSON
	// Jika aplikasi Dapodik terdeteksi kembalikan JSON dengan true dan kembalikan daftar semester yang ada pada Dapodik

	// ds.SetDataDapo.BaseURL = ""
	setDataDapo.Scheme = "http"
	setDataDapo.BaseURL = ctx.Query("ip_dapodik")
	setDataDapo.Port = ctx.Query("port")
	setDataDapo.FullPathUrl = fmt.Sprintf("%s://%s:%s", setDataDapo.Scheme, setDataDapo.BaseURL, setDataDapo.Port)
	setDataDapo.SekolahID = "8a5bd957-66bc-4096-9ff1-fee096b87ba0"

	endpoint := setDataDapo.FullPathUrl
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("error pada request", err)
	}
	client := ds.HTTPClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error pada saat akses", err)
		re.Success = false
		re.Message = err.Error()
		re.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, re)

	}
	defer resp.Body.Close()
	//
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Menampilkan response
	// fmt.Println("respons", string(body))
	// Cek apakah Aplikasi sudah terhubung dengan Database Dapodik, jika tidak tampilkan error
	if strings.Contains(string(body), "Tidak terhubung dengan database") {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from:", r)
				re.Message = "Tidak terhubung dengan database."
				re.Status = http.StatusInternalServerError
				re.Success = false
				ctx.JSON(http.StatusInternalServerError, re)
			}
		}()
		panic("Tidak terhubung dengan database.")
	}
	// Parse HTML dari body response
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("Error parsing HTML: ", err)
		return
	}
	var semester []string
	findSelectOptions(doc, "semester_id", &semester)
	fmt.Println(semester)

	re.Success = true
	re.Data = semester
	re.Message = "Aplikasi Dapodik sudah terhubung"
	re.Status = http.StatusOK

	setDataDapo.SemeterID = semester

	ctx.JSON(200, gin.H{"data": re})
}

func (ds DapodikClient) LoginToDapodik(ctx *gin.Context) {
	// Langkah 1: Kirim permintaan POST untuk login
	dataDapo := &setDataDapo

	// URL endpoint yang akan direquest
	endpoint := fmt.Sprintf("%v/roleperan", dataDapo.FullPathUrl)

	// Data form yang akan dikirim
	form := url.Values{}
	form.Add("username", ctx.Query("username"))
	form.Add("password", ctx.Query("password"))
	form.Add("semester_id", ctx.Query("semester_id"))

	// Membuat request dengan method POST
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println("Error create request: ", err)
		return
	}

	// Menambahkan header Conten-Type
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Membuat client HTTP dan mengirim request
	client := ds.HTTPClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	// Membaca response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	// ctx.JSON(200, string(body))
	// fmt.Println(string(body))

	// Parse HTML dari body response
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("Error parsing HTML: ", err)
		return
	}

	// cari elemen <a> dan dapatkan href
	sessionURL := findHref(doc)
	if sessionURL == "" {
		fmt.Println("Tidak menemukan elemen <a> dengan href yang valid")
		return
	}
	// ctx.JSON(200, string(body))
	// Parsing URL href untuk mendapatkan parameter session ID
	// u, err := url.Parse(sessionURL)
	// if err != nil {
	// 	fmt.Println("error parsing URL:", err)
	// 	return
	// }

	// fmt.Println("Href URL:", sessionURL)
	// fmt.Println("Query Paramters:", u.Query())
	// login ke dapodik

	endpoint = fmt.Sprintf("%s%s", dataDapo.FullPathUrl, sessionURL)
	// Membuat jar untuk menyimpan cookie
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Error creating cookie jar:", err)
		return
	}
	// Membuat client HTTP dan mengirim request
	// Membuat http.Client dengan cookie jar
	client.Jar = jar
	req, err = http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Error ketika membuat request", err)
		return
	}
	// Kirim permintaan
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Error terjadi mengirim request", err)
		return
	}
	defer resp.Body.Close()

	// fmt.Println(client.Jar)

	// fmt.Println("Protected resource response status:", resp.Status)
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }
	ctx.JSON(200, "login telah berhasil")
}

func (ds DapodikClient) GetPesertaDidik(ctx *gin.Context) {
	// Langkah 1: Kirim permintaan POST untuk login
	// dataDapo := &setDataDapo

	// URL endpoint yang akan direquest
	// endpoint := fmt.Sprintf("%v/WebService/getPesertaDidik", dataDapo.FullPathUrl)
	endpoint := "http://localhost:5774/WebService/getPesertaDidik"
	// Membuat URL dengan parameter query menggunakan package net/url
	u, err := url.Parse(endpoint)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// Menambahkan query parameters
	q := u.Query()
	q.Set("npsn", "20254180")     // Tambahkan parameter semester
	q.Set("semester_id", "20232") // Tambahkan parameter kelas
	u.RawQuery = q.Encode()       // Encode query menjadi string yang sesuai

	// req, err := http.NewRequest("GET", endpoint, nil)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Println("error request", err.Error())
		// return
	}
	bearerToken := "mkegY2Ps2W7b4HM"
	req.Header.Add("Authorization", "Bearer "+bearerToken)
	client := ds.HTTPClient
	// client.Jar = ds.HTTPClient.Jar
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error di respon", err)
		// return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	// err = json.NewDecoder(resp.Body).Decode()
	ctx.JSON(200, string(body))
	// fmt.Println(endpoint)
}

func (ds DapodikClient) GetSekolah(ctx *gin.Context) {

}
func (ds DapodikClient) GetGTK(ctx *gin.Context)              {}
func (ds DapodikClient) GetRombonganBelajar(ctx *gin.Context) {}
func (ds DapodikClient) GetPengguna(ctx *gin.Context)         {}

// func (ds DapodikClient) getBearer() {

// 	endpoint := "http://localhost:5774/rest/WsAplikasi"
// 	sekolahID := "8a5bd957-66bc-4096-9ff1-fee096b87ba0"
// 	// Request dengan parameter
// 	params := url.Values{}
// 	params.Add("sekolah_id", sekolahID)
// 	urlWithParams := fmt.Sprintf("%s?%s", endpoint, params.Encode())

// 	req, err := http.NewRequest("GET", urlWithParams, nil)
// 	if err != nil {
// 		fmt.Println("error sending request", err)
// 		return ""
// 	}
// 	client := ds.HTTPClient
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("error", err)
// 		return ""
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return ""
// 	}

// 	fmt.Println("Response:", string(body))
// 	// fmt.Sprintln()

// 	// // Memeriksa apakah respons berupa JSON
// 	// if resp.Header.Get("Content-Type") != "application/json" {
// 	// 	fmt.Println("Error: Expected JSON response, but got:", resp.Header.Get("Content-Type"))
// 	// 	return ""
// 	// }

// 	// // // Membaca body response
// 	// body, err := io.ReadAll(resp.Body)
// 	// if err != nil {
// 	// 	fmt.Println("Error reading response body:", err)
// 	// 	return ""
// 	// }

// 	// // Decode JSON response ke dalam slice dari PesertaDidik
// 	// // var pesertaDidik []PesertaDidik
// 	// var tes []webAPI
// 	// err = json.Unmarshal(body, &tes)
// 	// if err != nil {
// 	// 	fmt.Println("Error parsing JSON:", err)
// 	// 	return ""
// 	// }
// 	// fmt.Println(tes)
// 	return endpoint
// }

// // ---------------------------------------------------------
// // private function
func findHref(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				return attr.Val
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		href := findHref(c)
		if href != "" {
			return href
		}
	}
	return ""
}

// Fungsi untuk menemukan elemen <select> dengan name tertentu dan mengambil nilai <option>
func findSelectOptions(n *html.Node, nameAttr string, semesters *[]string) {
	if n.Type == html.ElementNode && n.Data == "select" {
		// Cari atribut "name" pada elemen <select>
		for _, attr := range n.Attr {
			if attr.Key == "name" && attr.Val == nameAttr {
				// Temukan semua <option> di dalam elemen <select>
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.ElementNode && c.Data == "option" {
						// Cari atribut "value" pada elemen <option>
						for _, optionAttr := range c.Attr {
							if optionAttr.Key == "value" {
								// Tambahkan value ke list semester
								*semesters = append(*semesters, optionAttr.Val)
							}
						}
					}
				}
				return
			}
		}
	}

	// Lanjutkan traversal node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findSelectOptions(c, nameAttr, semesters)
	}
}
