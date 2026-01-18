package internals

import (
	"encoding/json"
	"fmt"
	"net/http"
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
		models = append(models, m.Name)
	}

	return models, nil
}

func SelectModel(models []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select Ollama model",
		Items: models,
		Size:  10,
	}

	_, result, err := prompt.Run()
	return result, err
}