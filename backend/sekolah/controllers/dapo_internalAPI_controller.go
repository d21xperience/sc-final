package controllers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Mendapatkan WebService Dapodik
func (ds DapodikClient) GetWSDapodik(ctx *gin.Context) {
	// baseURL := fmt.Sprintf("%v/%s", setDataDapo.FullPathUrl, "rest/WsAplikasi")
	// endpoint := baseURL
	endpoint := "http://localhost:5774/rest/WsAplikasi?sekolah_id=8a5bd957-66bc-4096-9ff1-fee096b87ba0"
	fmt.Println(endpoint)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("error pada request", err)
	}
	client := ds.HTTPClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error pada saat akses", err)

		ctx.JSON(http.StatusInternalServerError, err)

	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	ctx.Data(resp.StatusCode, "application/json", body)

	// var rs []response.WsApp
	// er := util.BytesToStructSlice(body, &rs)
	// if er != nil {
	// 	log.Fatalf("Error: %v", err)
	// }
	// var wsApp []response.WsApp
	// for _, v := range rs {
	// 	ws := response.WsApp{
	// 		SekolahID: rs.sekolah_id,
	// 		Nama: v.nama,

	// 	}
	// // 	wsApp = append(wsApp, ws)
	// // }

	// re := response.WsReval{
	// 	Success: true,
	// 	Status:  resp.StatusCode,
	// 	Data:    rs,
	// }
	// ctx.JSON(200, re)
}

// Mendapatkan aplikasi dapodik
func (ds DapodikClient) CreateWSDapodik(ctx *gin.Context) {
	// sekolahID := ctx.Query("sekolah_id")
	fmt.Println(setDataDapo.BaseURL)
}
