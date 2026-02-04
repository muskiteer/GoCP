package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/muskiteer/GoCP/client/internal"
	"github.com/muskiteer/GoCP/client/internal/chat"
)



func main() {
	log.SetFlags(0)

	

	err := godotenv.Load()
	if err != nil {
		log.Fatal("No env file found")
	}


	models, err := internals.FetchModels()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Ollama server is running")


	

	model, err := internals.SelectModel(models)
	if err != nil {
		log.Fatal("Model selection cancelled")
	}
	log.Printf("Using model: %s\n", model)
	isnomic := internals.CheckNomicModel(models)

	if !isnomic {
		log.Println("Nomic embedding model not found. Please pull 'nomic-embed-text' model. for using the RAG feature.")
	}

	
	log.Println("Starting chat session. Type 'exit' to quit.")
	err = chat.ChatSession(model, isnomic)
	if err != nil {
		log.Fatal(err)
	}	
}
