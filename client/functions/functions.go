package functions

import (
	"bytes"
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/muskiteer/GoCP/client/structs"
)

func GetToolsPrompt() (string, error) {
	Server := os.Getenv("SERVER_URL")
	if Server == "" {
		return "", fmt.Errorf("SERVER_URL is not set")
	}
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

func IsToolCall(input string) bool {
	var obj map[string]any
	if json.Unmarshal([]byte(input), &obj) != nil {
		return false
	}
	_, ok1 := obj["tool"]
	_, ok2 := obj["arguments"]
	return ok1 && ok2
}


func GetToolResponse(tools string) (string, error) {
	server := os.Getenv("SERVER_URL")
	if server == "" {
		return "", fmt.Errorf("SERVER_URL is not set")
	}

	var requestBody structs.ToolExecute
	if err := json.Unmarshal([]byte(tools), &requestBody); err != nil {
		return "", fmt.Errorf("invalid tool JSON: %w", err)
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(
		server+"/tools/execution",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf(
			"tool server error (%d): %s",
			resp.StatusCode,
			string(body),
		)
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	toolResponse, exists := result["tools_response"]
	if !exists {
		return "", fmt.Errorf("missing 'tools_response' in server reply")
	}

	return toolResponse, nil
}

