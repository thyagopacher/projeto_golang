package middleware

import (
	"crypto/rsa"
	"projeto_go/internal/auth"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(verifyKey *rsa.PublicKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token não informado"})
			return
		}

		// Espera formato: Bearer TOKEN
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token inválido"})
			return
		}

		tokenString := parts[1]

		token, err := jwt.ParseWithClaims(tokenString, &auth.CustomClaims{}, func(token *jwt.Token) (any, error) {
			return verifyKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token inválido ou expirado"})
			return
		}

		claims, ok := token.Claims.(*auth.CustomClaims)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Erro ao ler claims"})
			return
		}

		// salva no contexto (igual session)
		c.Set("user", claims.Name)

		c.Next()
	}
}
