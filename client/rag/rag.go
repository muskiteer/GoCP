package rag

import (
	"bytes"
	"math"
	"sort"

	"github.com/ledongthuc/pdf"
	"github.com/muskiteer/GoCP/client/ollama"
	"github.com/muskiteer/GoCP/client/structs"
)

func ExtractPDFText(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}

	buf.ReadFrom(b)
	return buf.String(), nil
}

func CosineSimilarity(a, b []float64) float64 {
	var dot, magA, magB float64
	for i := range a {
		dot += a[i] * b[i]
		magA += a[i] * a[i]
		magB += b[i] * b[i]
	}
	return dot / (math.Sqrt(magA) * math.Sqrt(magB))
}


const MIN_SCORE = 0.6

func Search(query string, k int) []string {
	qVec, _ := ollama.Embed(query)

	type scored struct {
		Text  string
		Score float64
	}

	var scoredChunks []scored

	for _, c := range structs.MemoryStore {
		score := CosineSimilarity(qVec, c.Vector)
		scoredChunks = append(scoredChunks, scored{
			Text:  c.Text,
			Score: score,
		})
	}

	sort.Slice(scoredChunks, func(i, j int) bool {
		return scoredChunks[i].Score > scoredChunks[j].Score
	})

	var results []string
	for i := 0; i < len(scoredChunks) && len(results) < k; i++ {
		if scoredChunks[i].Score < MIN_SCORE {
			break 
		}
		results = append(results, scoredChunks[i].Text)
	}

	return results
}


