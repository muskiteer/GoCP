package internals

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/muskiteer/GoCP/client/structs"

	"github.com/manifoldco/promptui"
)

func FetchModels() ([]string, error) {
	resp, err := http.Get("http://localhost:11434/api/tags")
	if err != nil {
		return nil, fmt.Errorf("ollama server not running (run: ollama serve)")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ollama server unhealthy")
	}

	var tags structs.TagsResponse
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		return nil, err
	}
	
	if len(tags.Models) == 0 {
		return nil, fmt.Errorf("no models found (run: ollama pull <model>)")
	}

	var models []string
	for _, m := range tags.Models {
		// if strings.HasPrefix(m.Name, "nomic-embed-text") {
		// 	continue
		// }
		models = append(models, m.Name)
		log.Println("Found model:", m.Name)
	}

	return models, nil
}

func SelectModel(models []string) (string, error) {
filtered := make([]string, 0, len(models))

	for _, m := range models {
		if !strings.HasPrefix(m, "nomic-embed-text") {
			filtered = append(filtered, m)
		}
	}

	// models = filtered
	prompt := promptui.Select{
		Label: "Select Ollama model",
		Items: filtered,
		Size:  10,
	}

	_, result, err := prompt.Run()
	return result, err
}

func CheckNomicModel(models []string) bool {
	for _, m := range models {
		// log.Println(m)
		if strings.HasPrefix(m, "nomic-embed-text") {
			// log.Println("herer dsjf;ldsjfa")
			return true
		}
	}
	return false
}