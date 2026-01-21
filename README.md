# GoCP - Go Copilot with Tool Calling

A Go-based AI assistant that integrates with Ollama LLMs and provides intelligent tool calling capabilities.

## Features

- 🤖 **Interactive Chat Interface** - Terminal-based chat with Ollama models
- 🔧 **Automatic Tool Calling** - AI automatically calls tools without asking permission
- 📚 **Wikipedia Integration** - Fetch information from Wikipedia
- 💰 **Cryptocurrency Data** - Get real-time crypto prices
- 🔄 **Conversation Memory** - Maintains context throughout the session
- 🎯 **Smart Detection** - Automatically determines when external data is needed

## Architecture

```
GoCP/
├── server/              # Backend server
│   ├── handler/        # HTTP request handlers
│   ├── prompts/        # System prompt generation
│   ├── registery/      # Tool registry and execution
│   ├── routes/         # API routes
│   ├── schema/         # Tool definitions (tools.json)
│   ├── structs/        # Data structures
│   └── tools/          # Tool implementations
└── client/             # CLI client
    ├── functions/      # Tool detection and parsing
    ├── internals/      # Chat session and model selection
    ├── ollama/         # Ollama API integration
    └── structs/        # Client data structures
```

## Prerequisites

- Go 1.21 or higher
- Ollama installed and running
- A compatible Ollama model (llama3.1:8b, qwen2.5:7b, etc.)

## Installation

### 1. Install Ollama

```bash
# Linux
curl -fsSL https://ollama.com/install.sh | sh

# Start Ollama service
ollama serve
```

### 2. Pull a Compatible Model

```bash
# Recommended models for tool calling
ollama pull llama3.1:8b
# OR
ollama pull qwen2.5:7b
```

### 3. Clone and Setup

```bash
git clone <repository-url>
cd GoCP

# Install server dependencies
cd server
go mod download

# Install client dependencies
cd ../client
go mod download
```

## Configuration

### Server Setup

Create `.env` in the `server/` directory (optional):

```env
PORT=8080
```

### Client Setup

Create `.env` in the `client/` directory:

```env
OLLAMA_API_URL=http://localhost:11434
SERVER_URL=http://localhost:8080
```

## Running the Application

### Start the Server

```bash
cd server
go run main.go
```

You should see:
```
2026/01/21 14:29:22 Server is running on :8080
```

### Start the Client

In a new terminal:

```bash
cd client
go run main.go
```

The client will:
1. Check Ollama connection
2. Let you select a model
3. Start the chat session

## Usage Examples

### Basic Chat

```
GoCP> hello
Hello! How can I assist you today?
```

### Wikipedia Tool (Automatic)

```
GoCP> tell me about the book Lord of the Mysteries
[Tool automatically called: fetching_wikipedia]

Lord of the Mysteries is a Chinese web novel written by Cuttlefish That Loves Diving...
```

### Cryptocurrency Tool (Automatic)

```
GoCP> what's the current bitcoin price?
[Tool automatically called: fetching_crypto]

The current Bitcoin price is $43,250 USD.
```

### General Knowledge

```
GoCP> who invented Linux?
[Tool automatically called: fetching_wikipedia]

Linux was created by Linus Torvalds, a Finnish computer science student...
```

## Available Tools

### 1. Wikipedia Tool (`fetching_wikipedia`)

**Purpose**: Fetch information from Wikipedia

**When Used**: 
- User asks "what is", "who is", "tell me about"
- Questions about books, movies, people, places, events
- Any topic requiring factual information

**Arguments**:
- `query` (string, required): Search term for Wikipedia

### 2. Cryptocurrency Tool (`fetching_crypto`)

**Purpose**: Get real-time cryptocurrency prices

**When Used**:
- Questions about crypto prices
- Market data requests

**Arguments**:
- `crypto_name` (string, required): Name of cryptocurrency (e.g., "bitcoin")
- `currency` (string, required): Target currency (e.g., "usd")

## Adding New Tools

### 1. Define Tool Schema

Add to `server/schema/tools.json`:

```json
{
  "name": "your_tool_name",
  "description": "What your tool does",
  "arguments": {
    "param1": {
      "type": "string",
      "description": "Parameter description",
      "required": true
    }
  }
}
```

### 2. Implement Tool Function

Create `server/tools/your_tool.go`:

```go
package tools

import (
    "context"
    "fmt"
)

func YourToolName(ctx context.Context, arguments map[string]any) (string, error) {
    // Extract parameters
    param1, ok := arguments["param1"].(string)
    if !ok {
        return "", fmt.Errorf("param1 is required")
    }
    
    // Your tool logic here
    result := "Your result"
    
    return result, nil
}
```

### 3. Register the Tool

In `server/main.go`:

```go
registry.Register("your_tool_name", tools.YourToolName)
```

## API Endpoints

### Server Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/health` | GET | Health check |
| `/tools/prompts` | GET | Get system prompts |
| `/tools/execution` | POST | Execute a tool |

### Example Tool Execution Request

```bash
curl -X POST http://localhost:8080/tools/execute \
  -H "Content-Type: application/json" \
  -d '{
    "tool_name": "fetching_wikipedia",
    "arguments": {
      "query": "Artificial Intelligence"
    }
  }'
```


## Troubleshooting

### Tool Not Being Called

1. **Check Model Compatibility**: Use llama3.1:8b or qwen2.5:7b
2. **Verify System Prompt**: Check server logs for prompt generation
3. **Enable Debug Mode**: Set `DEBUG=true` in client/.env
4. **Check Server Logs**: Look for tool execution requests

### Connection Errors

```bash
# Verify Ollama is running
curl http://localhost:11434/api/tags

# Verify server is running
curl http://localhost:8080/health
```

### Model Not Responding

```bash
# Check available models
ollama list

# Pull a fresh model
ollama pull llama3.1:8b
```

## Performance Tips

1. **Model Selection**: Larger models (8B+) follow instructions better
2. **System Prompt**: Keep prompts clear and directive
3. **Tool Descriptions**: Make tool descriptions specific
4. **Conversation Context**: The system maintains full conversation history

## Development

### Project Structure

- **Server**: Handles tool execution and system prompts
- **Client**: Manages user interaction and Ollama communication
- **Tools**: Individual tool implementations
- **Registry**: Dynamic tool registration system

### Building for Production

```bash
# Build server
cd server
go build -o gocp-server main.go

# Build client
cd client
go build -o gocp-client main.go
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add your tool/feature
4. Test thoroughly
5. Submit a pull request

## Acknowledgments

- Ollama for the LLM backend
- CoinGecko API for cryptocurrency data
- Wikipedia API for knowledge retrieval

## Support

For issues and questions:
- Open an issue on GitHub
- Check existing documentation