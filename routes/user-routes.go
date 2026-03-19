package routes

import (
	"github.com/gin-gonic/gin"
	"projeto_go/internal/controllers" // Importe seus controllers
)

func SetupUsuarioRoutes(r *gin.Engine, usuarioController *controllers.UsuarioController) {
	usuarios := r.Group("/usuarios")
	{
		usuarios.GET("/", usuarioController.GetUsuarios)        // listar
		usuarios.GET("/:id", usuarioController.GetUsuarioByID)  // buscar 1
		usuarios.POST("/", usuarioController.CreateUsuario)     // criar
		usuarios.PUT("/:id", usuarioController.UpdateUsuario)   // atualizar
		usuarios.DELETE("/:id", usuarioController.DeleteUsuario) // deletar
	}
}