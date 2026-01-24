package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/muskiteer/GoCP/client/internals"
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
	
	log.Println("Starting chat session. Type 'exit' to quit.")
	err = internals.ChatSession(model)
	if err != nil {
		log.Fatal(err)
	}	
}
