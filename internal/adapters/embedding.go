package adapters

import "context"

type EmbeddingService struct{}

func NewEmbeddingService() *EmbeddingService {
	return &EmbeddingService{}
}

func (es *EmbeddingService) GenerateEmbeddings(
	ctx context.Context,
	text string,
) ([]float32, error) {
	// TODO: Implement embedding using ollama
	return nil, nil
}
