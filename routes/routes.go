package routes

import (
	"github.com/gin-gonic/gin"
	"projeto_go/internal/controllers"
)
	
type Controllers struct {
	Usuario *controllers.UsuarioController
	Fatura  *controllers.FaturaController
	Produto  *controllers.ProdutoController 
	Home  *controllers.HomeController
}

func SetupRoutes(r *gin.Engine, ctrls *Controllers) {
	SetupHomeRoutes(r, ctrls.Home)
	SetupUsuarioRoutes(r, ctrls.Usuario)
	SetupFaturaRoutes(r, ctrls.Fatura)
	SetupProdutoRoutes(r, ctrls.Produto)
}