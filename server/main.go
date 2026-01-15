package main

import (
	"log"
	"net/http"
	"github.com/muskiteer/GoCP/routes"
	"github.com/muskiteer/GoCP/prompts"
	// "os"
	// "context"
	
	// "net/http"
	// "github.com/muskiteer/GoCP/registery"
)

func main() {
	
	tools_prompt, err := prompts.ToolPromptGenerator()
	if err != nil {
		log.Fatal("Could not read tools_available.txt file: ", err)
	}
	mux:= http.NewServeMux()
	routes.SetupRoutes(mux, tools_prompt)
	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	
}