package internals

import (
	"fmt"
	"log"

	"github.com/muskiteer/GoCP/client/functions"
	"github.com/muskiteer/GoCP/client/ollama"
)

func ChatSession(model string) error {
	var cmd string
	for {
		fmt.Print("GoCP> ")
		_, err := fmt.Scanln(&cmd)
		if err != nil {
			return err
		}
		if cmd == "exit" {
			println("Exiting chat session.")
			break
		}
		tools_prompt, err := functions.GetToolsPrompt()
		if err != nil {
			return err
		}
		response , err := ollama.GetToolsResult(tools_prompt, model)
		if err != nil || response == "" {
			return err
		}
		log.Println("Response:\n", response)
		
	}
	return nil
}