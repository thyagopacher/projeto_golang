package controllers

import (
	"log"
	"net/http"
	"os"
	"projeto_go/internal/auth"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	secretKey []byte // HMAC (mais simples que RSA)
}

func NewAuthController(secret string) *AuthController {
	return &AuthController{
		secretKey: []byte(secret),
	}
}

func (a *AuthController) Authenticate(c *gin.Context) {
	var body struct {
		User string `json:"user"`
		Pass string `json:"pass"`
	}

	log.Println("Validando credenciais...")
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	// 🔥 validação fake (trocar por banco depois)
	log.Println("Usuario is valid...")
	if body.User != os.Getenv("AUTH_USER") || body.Pass != os.Getenv("AUTH_PASS") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciais inválidas"})
		return
	}

	// 🔐 cria claims
	claims := auth.CustomClaims{
		Name: body.User,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "projeto_go",
		},
	}

	// 🔑 cria token
	log.Println("Gerando token...")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(a.secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao gerar token"})
		return
	}

	log.Println("Retornando JSON token...")
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
