# GoCP — Go Copilot with RAG & Tool Calling

GoCP is a Go-based AI copilot that works with Ollama LLMs, supporting automatic tool calling and RAG (Retrieval-Augmented Generation) for querying PDFs and live data — all from the terminal.

## ✨ Features

🤖 Terminal-based AI chat with Ollama

🔧 Automatic tool calling (Wikipedia, web search, crypto prices)

📄 RAG support for querying PDF documents

🔍 DuckDuckGo web search

💰 Real-time crypto prices via CoinGecko

🧠 Context pruning & conversation memory

🎯 Interactive model selection at startup

## 🏗️ Architecture
```
GoCP/
├── server/   # Tool execution & APIs (port 8080)
└── client/   # CLI chat client with RAG
```

**Server:** Tool registry, execution engine, APIs

**Client:** Chat UI, Ollama interaction, RAG pipeline

## 📋 Prerequisites

- Go 1.21+
- Ollama (running locally)

**Recommended Models:**
```bash
ollama pull llama3.1:8b
ollama pull qwen2.5:7b
ollama pull nomic-embed-text   # Required for RAG
```

## 🚀 Installation
```bash
git clone <repo-url>
cd GoCP

# Build Server
cd server
go mod tidy
go build -o gocp-server

# Build Client
cd ../client
go mod tidy
go build -o gocp-client
```

## ⚙️ Configuration

**server/.env**
```env
PORT=8080
```

**client/.env**
```env
OLLAMA_API_URL=http://localhost:11434
SERVER_URL=http://localhost:8080
```

## ▶️ Running GoCP

**Terminal 1 — Server**
```bash
cd server
./gocp-server
```

**Terminal 2 — Client**
```bash
cd client
./gocp-client
```

You'll be prompted to select a model and can start chatting immediately.

## 💡 Usage Examples

**Ask Questions (Auto Tool Calling)**
```


**Crypto Prices**
```
GoCP> What's the Bitcoin price?
[🔧 Tool Called: Crypto]
```

**Web Search**
```
GoCP> Tell me about Alan Turing
[🔧 Tool Called: web search]
```

GoCP> What happened in tech today?
[🔧 Tool Called: web search]

```

**PDF RAG**
```
GoCP> rag it
[📄 RAG Context Retrieved]
```

**To Exit**
```
GoCP> exit

```

## 🛠️ Available Tools

| Tool | Purpose |
|------|----------|
| fetching_wikipedia | Encyclopedic knowledge |
| fetching_crypto | Live crypto prices |
| fetching_online | Web search |
| RAG | Semantic search over PDFs |

Tools are automatically selected by the model.

## ➕ Adding a New Tool (Quick Overview)

1. Define schema in `server/schema/tools.json`
2. Implement logic in `server/tool_internals/`
3. Add wrapper in `server/tools/`
4. Register in `server/registery/`
5. Restart server → tool becomes available

## 🐛 Troubleshooting

**Ollama not responding**
```bash
ollama serve
curl http://localhost:11434/api/tags
```

**Server not reachable**
```bash
curl http://localhost:8080/health
```

**RAG not working**
```bash
ollama pull nomic-embed-text
```

## 🧠 Performance Tips

- Use 8B models for best tool-calling
- Ensure at least 8–16GB RAM
- RAG context is auto-pruned for efficiency

## 🤝 Contributing

1. Fork the repo
2. Create a feature branch
3. Add code + tests
4. Open a PR 🚀

## 📄 License

MIT 

---

**Built with ❤️ using Go & Ollama**

*Last updated: Jan 28, 2026*