package adapters

import (
	"context"
	"customrag/internal/core/domain"
)

type EmbeddingService struct{}

func NewEmbeddingService() *EmbeddingService {
	return &EmbeddingService{}
}

func (es *EmbeddingService) GenerateEmbeddings(
	ctx context.Context,
	text string,
) ([]float32, error) {
	// TODO: Implement embedding using ollama
	// 1) Check of text param to see if it is valid
	if text == "" || text == " " {
		return nil, domain.ErrEmptyInput
	}
	// 2) Create embeddings request
	embeddingRequest := domain.NewEmbeddingRequest()

	// 3) return the embedding
	return nil, nil
}
