package routes

import (
	"projeto_go/internal/controllers" // Importe seus controllers

	"github.com/gin-gonic/gin"
)

func SetupProdutoRoutes(rg *gin.RouterGroup, ProdutoController *controllers.ProdutoController) {
	Produtos := rg.Group("/produtos")
	{
		Produtos.GET("/", ProdutoController.GetProdutos)         // listar
		Produtos.GET("/:id", ProdutoController.GetProdutoByID)   // buscar 1
		Produtos.POST("/", ProdutoController.CreateProduto)      // criar
		Produtos.PUT("/:id", ProdutoController.UpdateProduto)    // atualizar
		Produtos.DELETE("/:id", ProdutoController.DeleteProduto) // deletar
	}
}
