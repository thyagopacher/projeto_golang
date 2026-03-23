package routes

import (
	"github.com/gin-gonic/gin"
	"projeto_go/internal/controllers"
)

func SetupFaturaRoutes(rg *gin.RouterGroup, faturaController *controllers.FaturaController) {
	fatura := rg.Group("/fatura")
	{
		fatura.GET("/pdf", faturaController.GerarPDF)
	}
}