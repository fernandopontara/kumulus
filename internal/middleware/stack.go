package middleware

import "net/http"

// Aplica todos os middlewares em ordem
func ApplyMiddlewares(handler http.Handler) http.Handler {
	handler = CORSMiddleware(handler)
	handler = LoggingMiddleware(handler)
	handler = AuthMiddleware(handler)
	return handler
}
