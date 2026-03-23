package routes

import (
	"os"
	"projeto_go/internal/controllers"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	Usuario *controllers.UsuarioController
	Fatura  *controllers.FaturaController
	Produto *controllers.ProdutoController
	Home    *controllers.HomeController
}

func SetupRoutes(r *gin.Engine, ctrls *Controllers) {
	SetupHomeRoutes(r, ctrls.Home)
	SetupUsuarioRoutes(r, ctrls.Usuario)
	SetupFaturaRoutes(r, ctrls.Fatura)
	SetupProdutoRoutes(r, ctrls.Produto)

	api := r.Group("/api")
	{
		jwtSecret := os.Getenv("JWT_TOKEN")
		authCtrl := controllers.NewAuthController(jwtSecret)
		SetupAuthRoutes(api, authCtrl) // 🔓 público
	}
}
