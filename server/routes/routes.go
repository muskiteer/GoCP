package routes

import (
	"net/http"
	"github.com/muskiteer/GoCP/handler"
	
)

func SetupRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/health", handler.HealthcheckHandler)
	mux.HandleFunc("/tools/prompt", handler.ToolsPromptsHandler)
	mux.HandleFunc("/tools/execution", handler.ToolsExecutionHandler)
}

