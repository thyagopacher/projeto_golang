package routes

import (
	"github.com/gin-gonic/gin"
	"projeto_go/internal/controllers" // Importe seus controllers
)

func SetupProdutoRoutes(r *gin.Engine, ProdutoController *controllers.ProdutoController) {
	Produtos := r.Group("/Produtos")
	{
		Produtos.GET("/", ProdutoController.GetProdutos)        // listar
		Produtos.GET("/:id", ProdutoController.GetProdutoByID)  // buscar 1
		Produtos.POST("/", ProdutoController.CreateProduto)     // criar
		Produtos.PUT("/:id", ProdutoController.UpdateProduto)   // atualizar
		Produtos.DELETE("/:id", ProdutoController.DeleteProduto) // deletar
	}
}