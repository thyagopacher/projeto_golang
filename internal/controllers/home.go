package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "API rodando",
		"timestamp": time.Now().Format(time.RFC3339),
		"service": "projeto-go",
		"version": "1.0.0",
	})
}