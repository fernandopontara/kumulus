import (
	"net/http"
	"kumulus/internal/handler"
	"kumulus/internal/middleware"
	"kumulus/internal/config"
)

func main() {
	mux := http.NewServeMux()

	chatHandler := handler.ChatHandler{ /* injeções */ }

	mux.Handle("/chat", middleware.ApplyMiddlewares(http.HandlerFunc(chatHandler.Chat)))

	http.ListenAndServe(":8080", mux)
}

func NewClient() *Client {
	return &Client{
		APIKey:      config.GetEnv("OPENAI_API_KEY", ""),
		Model:       "gpt-3.5-turbo",
		Temperature: 0.7,
	}
}