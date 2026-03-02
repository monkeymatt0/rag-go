package inputs

import "context"

type EmbeddingInterface interface {
	GenerateEmbeddings(ctx context.Context, text string) ([]float32, error)
}
