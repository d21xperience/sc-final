package controller

import (
	"net/http"
	"sc-service/config"
	"sc-service/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentController struct {
	DB *gorm.DB
}

func NewStudentController(db *gorm.DB) *StudentController {
	return &StudentController{DB: db}
}

// Get Student mengambil dari projek sekolah

// Send to Blockhain

// List daftar semua siswa yang sudah dikirim ke blockhain
func (sc *StudentController) GetStudents(c *gin.Context) {
	var students []model.Student
	config.DB.Preload("Certificates").Find(&students)
	c.JSON(http.StatusOK, students)
}

// Get a single student
func (sc *StudentController) GetStudent(c *gin.Context) {
	id := c.Param("id")
	var student model.Student
	if err := config.DB.Preload("Certificates").First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (sc *StudentController) SearchStudent(c *gin.Context) {
	query := c.Query("query") // Ambil parameter query dari URL
	var student model.Student

	// Cari berdasarkan NISN atau NIK
	if err := config.DB.Where("nisn = ? OR nis = ?", query, query).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}
