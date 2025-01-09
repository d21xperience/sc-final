package routes

import (
	"gateway_service/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupGatewayRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/auth/login", func(c *gin.Context) {
		proxyRequest(c, "http://auth_service:8081/auth/login")
	})
	r.POST("/auth/register", func(c *gin.Context) {
		proxyRequest(c, "http://auth_service:8081/auth/register")
	})

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/sekolah", func(c *gin.Context) {
			proxyRequest(c, "http://sekolah_service:8083/api/sekolah")
		})
		api.POST("/smartcontract", func(c *gin.Context) {
			proxyRequest(c, "http://smartcontract_service:8082/api/verifikasi")
		})
	}
}

// Helper function to proxy requests
func proxyRequest(c *gin.Context, targetURL string) {
	req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Copy headers
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to forward request"})
		return
	}
	defer resp.Body.Close()

	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}
