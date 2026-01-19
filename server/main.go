package main

import (
	"log"
	"net/http"
	"github.com/muskiteer/GoCP/server/routes"
	"github.com/muskiteer/GoCP/server/prompts"
	"github.com/muskiteer/GoCP/server/registery"
	"context"
)

func main() {
	ctx := context.Background()
	path := "/home/muskiteer/Desktop/GoCP/server/schema/tools.json"
	manifest, err := registery.LoadToolManifest(path)
	if err != nil {
		log.Fatal("Failed to load tool manifest: ", err)
	}
	
	registry, err := registery.InitRegistry(manifest)
	if err != nil {
		log.Fatal("Failed to initialize registry: ", err)
	}



	tools_prompt, err := prompts.ToolPromptGenerator()
	if err != nil {
		log.Fatal("Could not read tools_available.txt file: ", err)
	}
	mux:= http.NewServeMux()
	routes.SetupRoutes(mux, tools_prompt, registry, ctx)
	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	
}