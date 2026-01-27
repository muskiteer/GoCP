package rag

import (
	"errors"

	"github.com/muskiteer/GoCP/client/internals"
	"github.com/muskiteer/GoCP/client/structs"
)

func RagthePDF() ([]structs.ChunkEmbedding, error) {
	filepath, err := internals.SelectPDF()
	if err != nil {
		return nil, err
	}

	text, err := internals.ExtractPDFText(filepath)
	if err != nil {
		return nil, err
	}

	chunk := ChunkText(text)
	vectors := ChunksToVectors(chunk)
	if len(vectors) == 0 {
		return nil, errors.New("No vectors were created from the document")
	}

	return vectors , nil
}