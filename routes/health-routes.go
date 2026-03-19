package routes

import (
	"github.com/gin-gonic/gin"
	"projeto_go/controllers" // Importe seus controllers
)

func SetupHealthRoutes(r *gin.Engine) {
	r.GET("/", controllers.HealthCheck)
}