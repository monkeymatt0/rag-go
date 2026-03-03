package domain

import (
	"context"
	"errors"
)

var (
	ErrEmptyInput = errors.New("Empty input: write something valid")
)

type EmbeddingService interface {
	GenerateEmbeddings(ctx context.Context, text string) ([]float32, error)
}

// This data structure will represent the vector.
type Vector struct{}

// payload of the request
type EmbeddingRequestPayload struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

// payload of the request
type EmbeddingResponse struct {
	Model      string      `json:"model"`
	Embeddings [][]float64 `json:"embeddings"`
}

func NewEmbeddingRequestPayload() *EmbeddingRequestPayload {
	return &EmbeddingRequestPayload{}
}
