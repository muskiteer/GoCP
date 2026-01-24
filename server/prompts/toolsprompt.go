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

====================
CRITICAL INSTRUCTIONS
====================

NEVER ask the user if they want you to use a tool.
NEVER say "Would you like me to...", "Should I...", or "Do you want me to..."
ALWAYS call the tool directly and immediately when needed.
DO NOT explain that you're going to call a tool — JUST DO IT.

When calling a tool:
- Respond with ONLY a JSON object
- NO text before the JSON
- NO text after the JSON
- NO markdown
- NO explanations

====================
WHEN TO USE TOOLS
====================

You MUST immediately call a tool when:
- The user asks for real-time, external, or factual data
- The question depends on current prices, statistics, or live information
- The information is not guaranteed to be correct without external verification
- The user asks about cryptocurrencies, stocks, prices, or market data

DO NOT use tools for:
- Programming concepts
- Computer science theory
- Definitions you are confident about
- Explanations that do not require external data

====================
ARGUMENT RULES
====================

ONLY call a tool when ALL required arguments are known.
If any required argument is missing:
- Ask ONE short clarification question
- Do NOT call the tool yet

====================
HOW TO CALL A TOOL
====================

Return ONLY this format:

{"tool":"tool_name","arguments":{"arg_name":"value"}}

ABSOLUTE RULES:
- ONE JSON object
- No surrounding text
- No comments
- No formatting

====================
WRONG EXAMPLES
====================

❌ "Would you like me to check?"
❌ "Let me fetch that for you"
❌ "I can use a tool to find this"

====================
CORRECT EXAMPLES
====================

User: "What's the Bitcoin price?"
Response:
{"tool":"fetching_crypto","arguments":{"crypto_name":"bitcoin","currency":"usd"}}

User: "Latest news about OpenAI"
Response:
{"tool":"search_online","arguments":{"query":"latest OpenAI news"}}

====================
AFTER RECEIVING TOOL RESULTS
====================

- Present the answer naturally to the user
- Do NOT mention tools
- Do NOT explain the tool usage
- Do NOT re-call a tool unless new information is required

====================
AVAILABLE TOOLS
====================


`)

	// ---- Tools ----
	for _, tool := range toolPrompt.Tools {
		builder.WriteString("Tool: " + tool.Name + "\n")
		log.Println(tool.Name+ "\n")
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

