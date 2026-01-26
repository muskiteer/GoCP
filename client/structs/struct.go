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

type ToolExecute struct {
	ToolName string  `json:"tool"`
	Arguments map[string]any `json:"arguments"`
}

type EmbeddingResponse struct {
	Embedding []float64 `json:"embedding"`
}

type ChunkEmbedding struct {
	Text      string
	Vector    []float64
}

var MemoryStore []ChunkEmbedding


