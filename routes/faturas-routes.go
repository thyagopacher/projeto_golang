package routes

import (
	"github.com/gin-gonic/gin"
	"projeto_go/controllers"
)

func SetupFaturaRoutes(r *gin.Engine) {
	fatura := r.Group("/fatura")
	{
		fatura.GET("/pdf", controllers.GerarFaturaPDF)
	}
}