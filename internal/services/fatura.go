package services

import (
	"log"
	"strings"
	"time"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type FaturaService struct {
	// Sem repositório, é um serviço simples
}

func NewFaturaService() *FaturaService {
	return &FaturaService{}
}

/**
* Gera PDF da fatura
* Recebe nome e valor como parâmetros e retorna os bytes do PDF
*/
func (s *FaturaService) GerarPdf(nome string, valor string) ([]byte, error) {
	// Valores padrão se não fornecidos
	if nome == "" {
		nome = "Cliente Teste"
	}
	if valor == "" {
		valor = "R$ 199,90"
	}

	html := s.montarHTML(nome, valor)

	// Criar gerador de PDF
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Println("Erro ao criar gerador PDF:", err)
		return nil, err
	}

	// Configurar página
	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))
	pdfg.AddPage(page)

	// Gerar PDF
	err = pdfg.Create()
	if err != nil {
		log.Println("Erro ao gerar PDF:", err)
		return nil, err
	}

	return pdfg.Bytes(), nil
}

/**
* Monta o HTML da fatura
*/
func (s *FaturaService) montarHTML(nome string, valor string) string {
	return `
	<html>
	<head>
		<meta charset="UTF-8">
		<style>
			body { font-family: Arial, sans-serif; }
			.container { 
				padding: 30px;
				max-width: 800px;
				margin: 0 auto;
			}
			h1 { 
				color: #333;
				border-bottom: 2px solid #333;
				padding-bottom: 10px;
			}
			.box {
				border: 1px solid #ccc;
				padding: 20px;
				margin-top: 20px;
				border-radius: 5px;
				background-color: #f9f9f9;
			}
			.label {
				font-weight: bold;
				color: #555;
			}
			.valor {
				font-size: 24px;
				color: #4CAF50;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Fatura</h1>
			<div class="box">
				<p><span class="label">Cliente:</span> ` + nome + `</p>
				<p><span class="label">Valor:</span> <span class="valor">` + valor + `</span></p>
				<p><span class="label">Data:</span> ` + time.Now().Format("02/01/2006 15:04") + `</p>
			</div>
		</div>
	</body>
	</html>
	`
}