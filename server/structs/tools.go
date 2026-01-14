package structs

import (
	"context"
)

type ToolManifest struct {
	Tools []ToolSpec `json:"tools"`
}

type ToolSpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Arguments   map[string]ToolArguments `json:"arguments"`
}	

type ToolArguments struct {
	Type     string `json:"type"`
	Description string `json:"description"`
}

type Tool struct {
    Spec     ToolSpec
    Execute  ToolExecutor
}

type ToolExecutor func(ctx context.Context, args map[string]any) (any, error)






