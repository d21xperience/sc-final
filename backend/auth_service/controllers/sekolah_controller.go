package controllers

// import (
// 	"auth_service/services"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type SekolahController struct {
// 	SekolahService services.SekolahService
// }

// func NewSekolahController(ss services.SekolahService) *SekolahController {
// 	return &SekolahController{SekolahService: ss}
// }

// func (ss SekolahController) GetSekolahByNpsn(ctx *gin.Context) {
// 	npsn := ctx.Param("npsn")
// 	sekolah, err := ss.SekolahService.GetSekolahByNpsn(npsn)
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, sekolah)
// }
