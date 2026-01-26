package ollama

import (
	"encoding/json"
	"net/http"
	"fmt"
	"bytes"
	"os"
	
	"github.com/muskiteer/GoCP/client/structs"
	"github.com/muskiteer/GoCP/client/functions"
)

func GetToolsResult( cmd string, messages *structs.OllamaTool) (string, error){
	OllamaUrl := os.Getenv("OLLAMA_API_URL")
	if OllamaUrl == "" {
		return "", fmt.Errorf("OLLAMA_API_URL is not set")
	}
	
	messages.Messages = append(messages.Messages, structs.OllamaMessages{
		Role:    "user",
		Content: cmd,
	})

	jsonData, err := json.Marshal(messages)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(OllamaUrl+"/api/chat", "application/json", 
	                      		bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Ollama server returned non-OK status: %d", resp.StatusCode)
	}

	var result structs.OllamaChatToolResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Message.Content, nil
}


func GetFinalResponse(response string,cmd string, messages *structs.OllamaTool) (string, error) {
	OllamaUrl := os.Getenv("OLLAMA_API_URL")
	if OllamaUrl == "" {
		return "", fmt.Errorf("OLLAMA_API_URL is not set")
	}
	toolAnw, err := functions.GetToolResponse(response)
	if err != nil || toolAnw == "" {
		return "", err
	}
messages.Messages = append(
	messages.Messages,
	structs.OllamaMessages{
		Role:    "tool",
		Content: "The following information was retrieved using an external tool. Use it ONLY to answer the user's question directly.\n\n" + toolAnw,
	},
)



	jsonData, err := json.Marshal(messages)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(OllamaUrl+"/api/chat", "application/json",
	                      		bytes.NewBuffer(jsonData))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Ollama server returned non-OK status: %d", resp.StatusCode)
	}

	var result structs.OllamaChatToolResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Message.Content, nil
}

func Embed(text string) ([]float64, error) {
	req := map[string]any{
		"model":  "nomic-embed-text",
		"prompt": text,
	}	

	body, _ := json.Marshal(req)

	resp, err := http.Post(
		"http://localhost:11434/api/embeddings",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result structs.EmbeddingResponse
	json.NewDecoder(resp.Body).Decode(&result)

	return result.Embedding, nil
}