package routes

import "github.com/gin-gonic/gin"
	
func SetupRoutes(r *gin.Engine) {
	SetupHealthRoutes(r)
	SetupUsuarioRoutes(r)
}