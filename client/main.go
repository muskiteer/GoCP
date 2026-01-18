package main

import (
	"log"
	"github.com/muskiteer/GoCP/client/internals"
)

type TagsResponse struct {
	Models []struct {
		Name string `json:"name"`
	} `json:"models"`
}



func main() {
	log.SetFlags(0)

	models, err := internals.FetchModels()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Ollama server is running")

	model, err := internals.SelectModel(models)
	if err != nil {
		log.Fatal("Model selection cancelled")
	}

	log.Printf("Using model: %s\n", model)
	
}
