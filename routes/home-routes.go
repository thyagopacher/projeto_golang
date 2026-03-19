package routes

import (
	"github.com/gin-gonic/gin"
	"projeto_go/internal/controllers" // Importe seus controllers
)

func SetupHomeRoutes(r *gin.Engine, homeController *controllers.HomeController) {
	r.GET("/", homeController.GetHome)
}