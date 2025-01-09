package controller

import (
	"net/http"
	"sc-service/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CertificateController struct {
	DB *gorm.DB
}

func NewCertificateController(db *gorm.DB) *CertificateController {
	return &CertificateController{DB: db}
}

// Get all certificates
func (cc *CertificateController) GetCertificates(c *gin.Context) {
	var certificates []model.Certificate
	cc.DB.Find(&certificates)
	c.JSON(http.StatusOK, certificates)
}

// Get a single certificate
func (cc *CertificateController) GetCertificate(c *gin.Context) {
	id := c.Param("id")
	var certificate model.Certificate
	if err := cc.DB.First(&certificate, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Certificate not found"})
		return
	}
	c.JSON(http.StatusOK, certificate)
}

// Create a new certificate
func (cc *CertificateController) CreateCertificate(c *gin.Context) {
	var certificate model.Certificate
	if err := c.ShouldBindJSON(&certificate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cc.DB.Create(&certificate)
	c.JSON(http.StatusCreated, certificate)
}

// Update an existing certificate
func (cc *CertificateController) UpdateCertificate(c *gin.Context) {
	id := c.Param("id")
	var certificate model.Certificate
	if err := cc.DB.First(&certificate, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Certificate not found"})
		return
	}

	if err := c.ShouldBindJSON(&certificate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cc.DB.Save(&certificate)
	c.JSON(http.StatusOK, certificate)
}

// Delete a certificate
func (cc *CertificateController) DeleteCertificate(c *gin.Context) {
	id := c.Param("id")
	if err := cc.DB.Delete(&model.Certificate{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete certificate"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Certificate deleted"})
}
