package chat

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"errors"

	"github.com/muskiteer/GoCP/client/functions"
	"github.com/muskiteer/GoCP/client/ollama"
	"github.com/muskiteer/GoCP/client/structs"
	"github.com/muskiteer/GoCP/client/rag"
	"github.com/muskiteer/GoCP/client/internals"
)

	func ChatSession(model string) error {
		
		reader := bufio.NewReader(os.Stdin)
		tools_prompt, err := functions.GetToolsPrompt()
		if err != nil || tools_prompt == "" {
			return errors.New("Failed to get tools prompt")
		}
		var messages structs.OllamaTool
		messages.Model = model
		messages.Stream = false
		messages.Messages = append(messages.Messages, structs.OllamaMessages{
			Role:    "system",
			Content: tools_prompt,
		})

		var embeddings []structs.ChunkEmbedding
		
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
			if cmd == "rag it" {
				embeddings, err = rag.RagthePDF()
				if err != nil || embeddings == nil {
				return errors.New("RAG initialization failed")
				}
				if len(embeddings) > 0 {
					fmt.Println("RAG is ready for your questions.")
				}
				continue
			}
			if len(embeddings) > 0 {

				results, err := rag.Search(cmd, embeddings)

				if err != nil{
					return err
				}
				if len(results) > 0 {
						rag_prompt := rag.GenerateRAGPrompt()
						var ragcontext strings.Builder
						ragcontext.WriteString("RAG CONTEXT:\n")
						for _, res := range results {
							ragcontext.WriteString(res + "\n")
						}
						var tempmessages structs.OllamaTool
						tempmessages.Model = model
						tempmessages.Stream = false
						tempmessages.Messages = append(tempmessages.Messages, structs.OllamaMessages{
							Role:    "system",
							Content: rag_prompt + "\n" +ragcontext.String(),
						})
						response , err := ollama.GetToolsResult( cmd, &tempmessages)
						if err != nil {
							return err
						}
						messages.Messages = append(messages.Messages, structs.OllamaMessages{
							Role:    "assistant",
							Content: response,
						})
						fmt.Println(response)
						continue
				}
			}
			
			response , err := ollama.GetToolsResult( cmd, &messages)
			if err != nil {
				return err
			}

			if(functions.IsToolCall(response)==false){
				fmt.Println(response)
				internals.PruneRAG(&messages.Messages)
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
			messages.Messages = messages.Messages[:1] // keep system only

			
		}
		return nil
	}