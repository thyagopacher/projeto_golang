package controllers

import (
	"strconv"
	"projeto_go/internal/models"
	"github.com/gin-gonic/gin"
	"projeto_go/internal/services"
)

type ProdutoController struct {
	service *services.ProdutoService
}

func NewProdutoController(service *services.ProdutoService) *ProdutoController {
	return &ProdutoController{
		service: service,
	}
}

/**
* GET /Produtos
*/
func (uc *ProdutoController) GetProdutos(c *gin.Context) {
    produtos, err := uc.service.GetProdutos()
    if err != nil {
        c.JSON(500, gin.H{
            "error": "Erro ao buscar produtos",
            "details": err.Error(),
        })
        return
    }
    
    c.JSON(200, produtos)
}

/**
* GET /Produtos/:id
*/
func (uc *ProdutoController) GetProdutoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	Produto, err := uc.service.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Produto)
}

/**
* POST /Produtos
*/
func (uc *ProdutoController) CreateProduto(c *gin.Context) {
	var input models.Produto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "JSON inválido"})
		return
	}

	Produto, err := uc.service.CreateProduto(input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, Produto)
}

/**
* PUT /Produtos/:id
*/
func (uc *ProdutoController) UpdateProduto(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input models.Produto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "JSON inválido"})
		return
	}

	Produto, err := uc.service.UpdateProduto(id, input)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Produto)
}

/**
* DELETE /Produtos/:id
*/
func (uc *ProdutoController) DeleteProduto(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	status, err := uc.service.DeleteProduto(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Produto removido", "success": status})
}