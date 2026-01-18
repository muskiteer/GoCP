package registery

import (
	"github.com/muskiteer/GoCP/server/structs"
	"context"
	// "github.com/muskiteer/GoCP/tools"
	"strings"
	// // "encoding/json"
	"fmt"
)

func ToolsExec(
	ctx context.Context,
	toolsNeeded structs.ToolExecute,
	registry *Registry,
) (string, error) {

	var builder strings.Builder

	
	builder.WriteString("Result for " + toolsNeeded.ToolName + ":\n")

	result, err := registry.Execute(ctx, toolsNeeded.ToolName, toolsNeeded.Arguments)
	// result, err := registry.Execute(ctx, toolsNeeded.ToolName, toolsNeeded.Arguments)
	if err != nil {
		builder.WriteString("Error executing tool: " + err.Error() + "\n")
	}

	builder.WriteString(fmt.Sprintf("%v\n\n", result))
	builder.WriteString("If another tool is needed, respond with a tool call in JSON.\nOtherwise, provide the final answer to the user.")
	return builder.String(), nil
}



