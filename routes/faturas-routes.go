package routes

import (
	"github.com/gin-gonic/gin"
	"projeto_go/internal/controllers"
)

func SetupFaturaRoutes(r *gin.Engine, faturaController *controllers.FaturaController) {
	fatura := r.Group("/fatura")
	{
		fatura.GET("/pdf", faturaController.GerarPDF)
	}
}