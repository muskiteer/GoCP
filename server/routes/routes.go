package routes

import (
	"net/http"
	"github.com/muskiteer/GoCP/handler"
	
)

func SetupRoutes(mux *http.ServeMux, tools_prompt string) {

	mux.HandleFunc("/chat/init", handler.StartSession)
	mux.HandleFunc("/health", handler.HealthcheckHandler)
	mux.HandleFunc("/tools/prompt", func(w http.ResponseWriter, r *http.Request) {
		handler.ToolsPromptsHandler(w, r, tools_prompt)
	})
	mux.HandleFunc("/tools/execution", handler.ToolsExecutionHandler)
}

