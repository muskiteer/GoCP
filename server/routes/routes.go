package routes

import (
	"net/http"
	"context"
	
	"github.com/muskiteer/GoCP/server/handler"
	"github.com/muskiteer/GoCP/server/registery"
	
)

func SetupRoutes(mux *http.ServeMux, tools_prompt string, registry *registery.Registry, ctx context.Context) {

	mux.HandleFunc("/chat/init", handler.StartSession)

	mux.HandleFunc("/health", handler.HealthcheckHandler)
	
	mux.HandleFunc("/tools/prompt", func(w http.ResponseWriter, r *http.Request) {
		handler.ToolsPromptsHandler(w, r, tools_prompt)
	})
	
	mux.HandleFunc("/tools/execution", func(w http.ResponseWriter, r *http.Request) {
		handler.ToolsExecutionHandler(w, r, *registry, ctx)
	})
}

