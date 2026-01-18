package prompts

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"fmt"
	"strings"
	"github.com/muskiteer/GoCP/server/structs"
)

func ToolPromptGenerator() (string, error) {
	path := filepath.Join(
		"/home/muskiteer/Desktop/GoCP/server/schema",
		"tools.json",
	)

	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read tools.json: %w", err)
	}

	var toolPrompt structs.ToolManifest
	if err := json.Unmarshal(data, &toolPrompt); err != nil {
		return "", fmt.Errorf("failed to unmarshal tools.json: %w", err)
	}

	var builder strings.Builder

	// ---- System Prompt ----
	builder.WriteString(`System Prompt:
You are an AI assistant with access to external tools.

You may call a tool ONLY if it is necessary to answer the user's request.
You MUST NOT invent tools.
You MUST ONLY use tools listed below.

--------------------
TOOL CALLING RULES
--------------------

If you decide to call a tool:
- Respond with ONLY a valid JSON object
- Do NOT include any extra text, explanation, or formatting
- The JSON must be the entire response
- Use this exact format:

{
  "tool": "<tool_name>",
  "arguments": {
    "<argument_name>": "<value>"
  }
}

If you do NOT need a tool:
- Respond normally in plain text
- Do NOT return JSON

--------------------
AFTER TOOL EXECUTION
--------------------

After a tool is executed, you will receive its result.
Use the tool result to produce the final answer to the user.
Do NOT call the same tool again unless the input changes.

--------------------
AVAILABLE TOOLS
--------------------

`)

	// ---- Tools ----
	for _, tool := range toolPrompt.Tools {
		builder.WriteString("Tool: " + tool.Name + "\n")
		builder.WriteString("Description: " + tool.Description + "\n\n")

		if len(tool.Arguments) > 0 {
			builder.WriteString("Arguments:\n")

			for argName, arg := range tool.Arguments {
				builder.WriteString(
					"- " + argName + " (" + arg.Type + ")\n" +
						"  Description: " + arg.Description + "\n\n",
				)
			}
		}

		builder.WriteString("\n")
	}

	log.Println("Tool prompt generated successfully.")
	
	return builder.String(), nil
}

