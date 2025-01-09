package controller

import (
	"net/http"
	"os/exec"
	"sc-service/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	shell "github.com/ipfs/go-ipfs-api"
	"gorm.io/gorm"
)

type IPFSController struct {
	DB *gorm.DB
}

func NewIPFSController(db *gorm.DB) *IPFSController {
	return &IPFSController{DB: db}
}

// get metadata
func (ac *IPFSController) GetMetadata(c *gin.Context) {
	var metadata []model.IPFSMetadata
	if err := ac.DB.Find(&metadata).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, metadata)
}

// post metadata
func (ac *IPFSController) PostMetadata(c *gin.Context) {
	var metadata model.IPFSMetadata
	if err := c.ShouldBindJSON(&metadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ac.DB.Create(&metadata).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, metadata)
}

// Get node
func (ac *IPFSController) GetNode(c *gin.Context) {
	var nodes []model.IPFSNode
	if err := ac.DB.Preload("Files").Find(&nodes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nodes)
}

// Post node
func (ac *IPFSController) PostNode(c *gin.Context) {
	var node model.IPFSNode
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ac.DB.Create(&node).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, node)
}
func (ac *IPFSController) UploadIjazah(c *gin.Context) {
	// Ambil file dari form
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	// Buka file untuk dibaca
	fileContent, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}
	defer fileContent.Close()

	// Inisialisasi koneksi ke IPFS
	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(fileContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload to IPFS"})
		return
	}

	// Simpan metadata file ke PostgreSQL
	db := c.MustGet("db").(*gorm.DB)
	newMetadata := model.IPFSMetadata{
		IPFSCID:    cid,
		Size:       int(file.Size),
		FileName:   file.Filename,
		FileType:   http.DetectContentType([]byte(file.Filename)),
		UploadedAt: time.Now(),
	}

	if err := db.Create(&newMetadata).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save metadata"})
		return
	}

	// Kirimkan CID sebagai response
	c.JSON(http.StatusOK, gin.H{"cid": cid})

}

// ------------------------------------------------
// IPFS
// ------------------------------------------------

func runCommand(command string, args ...string) (string, error) {
	// Eksekusi perintah sistem
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// Handler untuk mendapatkan status IPFS
func (ac *IPFSController) GetIPFSStatusHandler(c *gin.Context) {
	// Jalankan perintah `ipfs id`
	output, err := runCommand("ipfs", "12D3KooWFmB4rMCXfHnypUoy3KUUPGAcs7J8GPymxSMFVaUgaMFa")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve IPFS status", "details": err.Error()})
		return
	}

	// Kirimkan output sebagai response
	c.JSON(200, gin.H{"status": "success", "data": strings.TrimSpace(output)})
}

// Handler untuk mendapatkan informasi swarm peers
func (ac *IPFSController) GetSwarmPeersHandler(c *gin.Context) {
	// Jalankan perintah `ipfs swarm peers`
	output, err := runCommand("ipfs", "swarm", "peers")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve swarm peers", "details": err.Error()})
		return
	}

	// Kirimkan output sebagai response
	peers := strings.Split(strings.TrimSpace(output), "\n")
	c.JSON(200, gin.H{"status": "success", "peers": peers})
}

// Handler untuk mendapatkan informasi penyimpanan
func (ac *IPFSController) GetRepoStatsHandler(c *gin.Context) {
	// Jalankan perintah `ipfs repo stat`
	output, err := runCommand("ipfs", "repo", "stat")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve repo stats", "details": err.Error()})
		return
	}

	// Kirimkan output sebagai response
	c.JSON(200, gin.H{"status": "success", "data": strings.TrimSpace(output)})
}
