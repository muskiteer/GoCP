package internals

import (
	"regexp"
	"strings"
)

func CleanText(s string) string {
	s = strings.ReplaceAll(s, "\r", "")
	s = regexp.MustCompile(`\n{2,}`).ReplaceAllString(s, "\n")
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}


func ChunkText(text string, size, overlap int) []string {
	var chunks []string
	runes := []rune(text)

	for i := 0; i < len(runes); i += size - overlap {
		end := i + size
		if end > len(runes) {
			end = len(runes)
		}
		chunks = append(chunks, string(runes[i:end]))
	}

	return chunks
}
