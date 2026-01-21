package handler

import (
	"encoding/json"
	"log"
	"context"
	"net/http"
	"github.com/muskiteer/GoCP/server/registery"
	"github.com/muskiteer/GoCP/server/structs"
)

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func StartSession(w http.ResponseWriter, r *http.Request) {
	// Placeholder for session initialization logic
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Session Started"})
}

func ToolsPromptsHandler(w http.ResponseWriter, r *http.Request, tools_prompt string) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var response = map[string]string{
		"prompt_tools": tools_prompt,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)		
}

func ToolsExecutionHandler(w http.ResponseWriter, r *http.Request, registry registery.Registry, ctx context.Context) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var tools_needed structs.ToolExecute
	json.NewDecoder(r.Body).Decode(&tools_needed)
	log.Println("Received tool execution request for tool:", tools_needed.ToolName)
	
	response_str, err := registery.ToolsExec(ctx, tools_needed, &registry)
	if err != nil {
		http.Error(w, "Error executing tools", http.StatusInternalServerError)
		log.Println("Error in ToolsExecutionHandler:", err)
		return
	}
	
	var response = map[string]string{
		"tools_response": response_str,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	

}