package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"projeto_go/internal/services"
)

type FaturaController struct {
	service *services.FaturaService
}

func NewFaturaController(service *services.FaturaService) *FaturaController {
	return &FaturaController{
		service: service,
	}
}


/**
* GET /fatura/pdf
*/
func (fc *FaturaController) GerarPDF(c *gin.Context) {
	// Pegar parâmetros da query string
	nome := c.Query("nome")
	valor := c.Query("valor")

	// Chamar service para gerar PDF
	pdfBytes, err := fc.service.GerarPdf(nome, valor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao gerar PDF",
			"details": err.Error(),
		})
		return
	}

	// Headers para download
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=fatura.pdf")
	c.Header("Content-Length", string(len(pdfBytes)))

	// Enviar PDF
	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}