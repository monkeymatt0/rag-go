package adapters

import (
	"bytes"
	"context"
	"customrag/internal/core/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

type EmbeddingService struct {
	service string
	model   string
	dim     int
}

func NewEmbeddingService(model string, dim int) *EmbeddingService {
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
	embeddingRequestPayload := domain.NewEmbeddingRequestPayload()
	embeddingRequestPayload.Model = es.model
	embeddingRequestPayload.Input = text

	// Marshaling the go structure to obtain JSON
	eb, err := json.Marshal(embeddingRequestPayload)
	if err != nil {
		return nil, err
	}
	// Transforming in bytes
	beb := bytes.NewBuffer(eb)
	resp, err := http.Post(es.service, "application/json", beb)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	// 3) return the embedding
	return nil, nil
}
