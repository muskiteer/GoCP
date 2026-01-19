package functions

import (
	"encoding/json"
	"net/http"
	"os"
)

func GetToolsPrompt() (string, error) {
	Server := os.Getenv("SERVER_URL")
	resp, err := http.Get(Server + "/tools/prompt")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", err
	}
	var result map[string]string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	tools_prompt, exists := result["prompt_tools"]
	if !exists {
		return "", err
	}
	return tools_prompt, nil	
}