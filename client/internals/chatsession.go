	package internals

	import (
		"bufio"
		"fmt"
		"log"
		"os"
		"strings"
		"github.com/muskiteer/GoCP/client/functions"
		"github.com/muskiteer/GoCP/client/ollama"
		"github.com/muskiteer/GoCP/client/structs"
	)

	func ChatSession(model string) error {
		
		reader := bufio.NewReader(os.Stdin)
		tools_prompt, err := functions.GetToolsPrompt()
		if err != nil || tools_prompt == "" {
			return err	
		}
		var messages structs.OllamaTool
		messages.Model = model
		messages.Stream = false
		messages.Messages = append(messages.Messages, structs.OllamaMessages{
			Role:    "system",
			Content: tools_prompt,
		})
		
		for {
			fmt.Print("GoCP> ")
			cmd, err := reader.ReadString('\n')
			
			if err != nil {
				log.Println("\nTry again")
			}
			cmd = strings.TrimSpace(cmd)

			if cmd == "exit" {
				println("Exiting chat session.")
				break
			}
			if cmd == "\n" || cmd == "" {
				continue
			}

			
			response , err := ollama.GetToolsResult( cmd, &messages)
			if err != nil || response == "" {
				return err
			}

			if(functions.IsToolCall(response)==false){
				fmt.Println(response)
				messages.Messages = append(messages.Messages, structs.OllamaMessages{
					Role:    "assistant",
					Content: response,
				})
				continue
			}
			messages.Messages = append(messages.Messages, structs.OllamaMessages{
				Role:    "assistant",
				Content: response,
			})
			
			final_response, err := ollama.GetFinalResponse(response,cmd, &messages)
			if err != nil || final_response == "" {
				return err
			}
			
			fmt.Println(final_response)
			messages.Messages = append(messages.Messages, structs.OllamaMessages{
				Role:    "assistant",
				Content: final_response,
			})
			
		}
		return nil
	}