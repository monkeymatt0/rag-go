package domain

import (
	"context"
)

type EmbeddingService interface {
	GenerateEmbeddings(ctx context.Context, text string) ([]float32, error)
}

// This data structure will represent the vector.
type Vector struct{}
