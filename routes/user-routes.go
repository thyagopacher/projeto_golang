package routes

import (
	"projeto_go/internal/controllers" // Importe seus controllers

	"github.com/gin-gonic/gin"
)

func SetupUsuarioRoutes(rg *gin.RouterGroup, usuarioController *controllers.UsuarioController) {
	usuarios := rg.Group("/usuarios")
	{
		usuarios.GET("/", usuarioController.GetUsuarios)         // listar
		usuarios.GET("/:id", usuarioController.GetUsuarioByID)   // buscar 1
		usuarios.POST("/", usuarioController.CreateUsuario)      // criar
		usuarios.PUT("/:id", usuarioController.UpdateUsuario)    // atualizar
		usuarios.DELETE("/:id", usuarioController.DeleteUsuario) // deletar
	}
}
