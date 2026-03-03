package inputs

import (
	"context"
	"customrag/internal/core/domain"
)

type EmbeddingInterface interface {
	GenerateEmbeddings(ctx context.Context, text string) (*domain.EmbeddingResponse, error)
}
