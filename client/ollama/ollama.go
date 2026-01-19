package ollama

import (
	"encoding/json"
	"net/http"
	"log"
	"bytes"
	"github.com/muskiteer/GoCP/client/structs"
	"github.com/muskiteer/GoCP/client/functions"
)

func GetToolsResult(user_prompt string, model string) (string, error){
	tool_prompt, err := functions.GetToolsPrompt()
	if err != nil || tool_prompt == "" {
		return "", err	
	}
	response := structs.OllamaTool{
		Model: model,
		Stream : false,
		Messages: []structs.OllamaMessages{
			{
				Role:    "system",
				Content: tool_prompt,
			},
			{
				Role:    "user",
				Content: user_prompt,
			},
		},
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	resp, err := http.Post("http://localhost:11434/api/chat", "application/json", 
	                      		bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ollama server returned non-OK status: %d", resp.StatusCode)
	}

	var result structs.OllamaChatToolResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Message.Content, nil
}