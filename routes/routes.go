package routes

import (
	"os"
	"projeto_go/internal/controllers"
	"projeto_go/internal/middleware"

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

	api := r.Group("/api")
	jwtSecret := os.Getenv("JWT_TOKEN")
	authCtrl := controllers.NewAuthController(jwtSecret)
	SetupAuthRoutes(api, authCtrl) // 🔓 público

	// 🔒 protegidas
	protected := api.Group("/")
	protected.Use(middleware.JWTMiddleware([]byte(jwtSecret)))
	
	SetupUsuarioRoutes(protected, ctrls.Usuario)
	SetupFaturaRoutes(protected, ctrls.Fatura)
	SetupProdutoRoutes(protected, ctrls.Produto)
}
