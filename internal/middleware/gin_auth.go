package middleware

import (
	"net/http"
	"strings"

	"kumulus/internal/auth"

	"github.com/gin-gonic/gin"
)

func GinAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := auth.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		// Injeta os dados no contexto do Gin
		c.Set("auth", claims)
		c.Next()
	}
}
