package middleware

import "github.com/gin-gonic/gin"

func ApplyAll() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		CORSMiddleware(),
		GinAuthMiddleware(),
	}
}
