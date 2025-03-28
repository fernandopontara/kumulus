package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Client struct {
	APIKey      string
	Model       string
	Temperature float32
}

func NewClient() *Client {
	return &Client{
		APIKey:      os.Getenv("OPENAI_API_KEY"),
		Model:       "gpt-3.5-turbo", // ou gpt-4 se preferir
		Temperature: 0.7,
	}
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model       string        `json:"model"`
	Messages    []ChatMessage `json:"messages"`
	Temperature float32       `json:"temperature"`
}

type ChatResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

func (c *Client) EnviarPrompt(prompt string) (string, error) {
	body := ChatRequest{
		Model:       c.Model,
		Temperature: c.Temperature,
		Messages: []ChatMessage{
			{Role: "system", Content: "Você é um assistente de vendas treinado pela empresa."},
			{Role: "user", Content: prompt},
		},
	}

	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("OpenAI erro %d: %s", resp.StatusCode, string(data))
	}

	var resposta ChatResponse
	err = json.Unmarshal(data, &resposta)
	if err != nil {
		return "", err
	}

	if len(resposta.Choices) == 0 {
		return "", fmt.Errorf("sem resposta do modelo")
	}

	return resposta.Choices[0].Message.Content, nil
}
