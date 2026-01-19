package structs

import (
	
)

type TagsResponse struct {
	Models []struct {
		Name string `json:"name"`
	} `json:"models"`
}

type OllamaTool struct {
	Model        string `json:"model"`
	Stream 	bool   `json:"stream"`
	Messages []OllamaMessages `json:"messages"`
}

type OllamaMessages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OllamaChatToolResponse struct {
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	Done bool `json:"done"`
}


