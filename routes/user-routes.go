package routes

import (
	"github.com/gin-gonic/gin"
	"projeto_go/controllers" // Importe seus controllers
)

func SetupUsuarioRoutes(r *gin.Engine) {
	usuarios := r.Group("/usuarios")
	{
		usuarios.GET("/", controllers.GetUsuarios)        // listar
		usuarios.GET("/:id", controllers.GetUsuarioByID)  // buscar 1
		usuarios.POST("/", controllers.CreateUsuario)     // criar
		usuarios.PUT("/:id", controllers.UpdateUsuario)   // atualizar
		usuarios.DELETE("/:id", controllers.DeleteUsuario) // deletar
	}
}