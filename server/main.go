package main

import (
	"log"
	"net/http"
	"context"
	"os"
	
	"github.com/muskiteer/GoCP/server/routes"
	"github.com/muskiteer/GoCP/server/prompts"
	"github.com/muskiteer/GoCP/server/registery"
	
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, proceeding with environment variables")
	}
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
	port := os.Getenv("PORT")
	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}