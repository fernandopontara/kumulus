package main

import (
	"kumulus/internal/config"
	"kumulus/internal/handler"
	"kumulus/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Carrega as variáveis de ambiente
	config.LoadEnv()

	// Inicializa os repositórios e serviços necessários
	userRepo := config.InitializeUserRepo()                 // Certifique-se de implementar esta função
	chatService := config.InitializeChatService()           // Certifique-se de implementar esta função
	knowledgeService := config.InitializeKnowledgeService() // Certifique-se de implementar esta função

	// Inicializa os handlers com as dependências
	authHandler := handler.AuthHandler{UserRepo: userRepo}
	chatHandler := handler.ChatHandler{ChatService: chatService}
	knowledgeHandler := handler.KnowledgeHandler{Service: knowledgeService}

	// Configura o roteador do Gin
	r := gin.Default()
	r.Use(middleware.CORSMiddleware()) // Aplica globalmente o middleware de CORS

	// Rota pública para login
	r.POST("/login", authHandler.LoginHandler)

	// Grupo protegido com autenticação multitenant
	auth := r.Group("/")
	auth.Use(middleware.GinAuthMiddleware())
	{
		auth.POST("/chat", chatHandler.HandleChat)
		auth.POST("/knowledge", knowledgeHandler.HandleKnowledge)
	}

	// Inicia o servidor na porta 8080
	r.Run(":8080")
}
