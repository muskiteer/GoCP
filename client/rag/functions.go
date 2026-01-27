package rag

import (
	"math"
	"sort"

	"github.com/muskiteer/GoCP/client/ollama"
	"github.com/muskiteer/GoCP/client/structs"
)

func ChunkText(text string) []string {
	const size = 1000
	const overlap = 100
	var chunks []string
	runes := []rune(text)

	for i := 0; i < len(runes); i += size - overlap {
		end := i + size
		if end > len(runes) {
			end = len(runes)
		}else {
			for end > i && runes[end] != ' ' {
			end--
			}
		}
		chunks = append(chunks, string(runes[i:end]))
	}

	return chunks
}

func CosineSimilarity(a, b []float64) float64 {
	var dot, magA, magB float64
	for i := range a {
		dot += a[i] * b[i]
		magA += a[i] * a[i]
		magB += b[i] * b[i]
	}
	denom := math.Sqrt(magA) * math.Sqrt(magB)
	if denom == 0 {
		return 0
	}
	return dot / denom

}

func ChunksToVectors(chunks []string) []structs.ChunkEmbedding {
	var store []structs.ChunkEmbedding
	for _, chunk := range chunks {
		vec, err := ollama.Embed(chunk)
		if err != nil {
			continue
		}
		store = append(store, structs.ChunkEmbedding{
			Text:   chunk,
			Vector: vec,
		})
	}

	return store
}





func Search(query string, store []structs.ChunkEmbedding) ([]string, error) {
	const MIN_SCORE = 0.4
	const k = 5

	qVec, err := ollama.Embed(query)
	if err != nil {
		return nil, err
	}

	scored := make([]structs.Scored, 0, len(store))

	for _, c := range store {
		if len(qVec) != len(c.Vector) {
			continue
		}

		score := CosineSimilarity(qVec, c.Vector)
		scored = append(scored, structs.Scored{
			Text:  c.Text,
			Score: score,
		})
	}

	sort.Slice(scored, func(i, j int) bool {
		return scored[i].Score > scored[j].Score
	})

	var results []string
	for i := 0; i < len(scored) && len(results) < k; i++ {
		if scored[i].Score >= MIN_SCORE {
			results = append(results, scored[i].Text)
		}
	}

	return results, nil
}
