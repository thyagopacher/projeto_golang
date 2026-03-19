package controllers

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

/**
 * Controller para gerar fatura em PDF
 * Endpoint: GET /fatura/pdf
 * gera um PDF simples usando wkhtmltopdf
 * @author Thyago Henrique Pacher
 */
func GerarFaturaPDF(c *gin.Context) {

	// Simulando dados (depois você pode puxar do banco)
	nome := c.Query("nome") // get parameter 
	if nome == "" {
		nome = "Cliente Teste"
	}	
	valor := c.Query("valor") // get parameter 
	if valor == "" {
		valor = "R$ 199,90"
	}

	html := `
	<html>
	<head>
		<style>
			body { font-family: Arial; }
			.container { padding: 20px; }
			h1 { color: #333; }
			.box {
				border: 1px solid #ccc;
				padding: 10px;
				margin-top: 20px;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Fatura</h1>
			<div class="box">
				<p><strong>Cliente:</strong> ` + nome + `</p>
				<p><strong>Valor:</strong> ` + valor + `</p>
			</div>
		</div>
	</body>
	</html>
	`

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Erro ao criar PDF"})
		return
	}

	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Erro ao gerar PDF"})
		return
	}

	// Headers pra download
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=fatura.pdf")

	c.Data(200, "application/pdf", pdfg.Bytes())
}