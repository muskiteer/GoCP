package structs

import (

)

type ToolExecute struct {
	ToolName string  `json:"tool"`
	Arguments map[string]any `json:"arguments"`
}
