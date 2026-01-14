package routes

import (
	"net/http"

	tools "github.com/muskiteer/GoCP/tools/fetching_crypto"
)

func SetupRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/health", healthcheckHandler)
	mux.HandleFunc("/tools_Prompts", toolsPromptsHandler)
	mux.HandleFunc("/tools_execution", toolsExecutionHandler)
}

