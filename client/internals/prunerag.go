package internals 

import (
	"strings"
	"github.com/muskiteer/GoCP/client/structs"
)


func PruneRAG(messages *[]structs.OllamaMessages) {
	filtered := (*messages)[:0]

	for _, msg := range *messages {
		if strings.HasPrefix(msg.Content, "[RAG_CONTEXT]") {
			continue
		}
		filtered = append(filtered, msg)
	}

	*messages = filtered
}
