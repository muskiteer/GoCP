// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
	
// 		log.SetFlags(0)

// 	_,err :=http.Get("http://localhost:11434/")
// 	if err !=nil{
// 		log.Fatal("Ollama server not running.\nPlease use ollama serve to start the server.")
// 	}
// 	log.Println("Ollama server is running.")
// 	log.Println("Enter a model you want to use like 'llama3.1:8b'")
// 	log.Println("you can see the name of the models by running 'ollama list'")
// 	var model string
// 	fmt.Print("Model: ")
// 	_, err = fmt.Scanln(&model)
// 	if err != nil {
// 		log.Fatal("Failed to read model input:", err)
// 	}

// 	log.Printf("Using model: %s\n", model)
	
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/manifoldco/promptui"
)

type TagsResponse struct {
	Models []struct {
		Name string `json:"name"`
	} `json:"models"`
}

func fetchModels() ([]string, error) {
	resp, err := http.Get("http://localhost:11434/api/tags")
	if err != nil {
		return nil, fmt.Errorf("ollama server not running (run: ollama serve)")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ollama server unhealthy")
	}

	var tags TagsResponse
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		return nil, err
	}

	if len(tags.Models) == 0 {
		return nil, fmt.Errorf("no models found (run: ollama pull <model>)")
	}

	var models []string
	for _, m := range tags.Models {
		models = append(models, m.Name)
	}

	return models, nil
}

func selectModel(models []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select Ollama model",
		Items: models,
		Size:  10,
	}

	_, result, err := prompt.Run()
	return result, err
}

func main() {
	log.SetFlags(0)

	models, err := fetchModels()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Ollama server is running")

	model, err := selectModel(models)
	if err != nil {
		log.Fatal("Model selection cancelled")
	}

	log.Printf("Using model: %s\n", model)
	

	// next:
	// startInteractiveShell(model)
}
