package structs

import (

)

type PromptRequest struct {
	Prompt string `json:"prompt"`
}

type ToolExecute struct {
	ToolName string  `json:"tool"`
	Arguments map[string]any `json:"arguments"`
}
