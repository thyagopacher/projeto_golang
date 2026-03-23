package routes

import (
	"projeto_go/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(rg *gin.RouterGroup, ctrl *controllers.AuthController) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", ctrl.Authenticate)
	}
}
