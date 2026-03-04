package adapters

import (
	"bytes"
	"context"
	"customrag/internal/core/domain"
	"encoding/json"
	"io"
	"net/http"
)

type EmbeddingService struct {
	Service string
	Model   string
	Dim     int
}

func NewEmbeddingService(model string, dim int) *EmbeddingService {
	return &EmbeddingService{}
}

func (es *EmbeddingService) GenerateEmbeddings(
	ctx context.Context,
	text string,
) (*domain.EmbeddingResponse, error) {
	// TODO: Implement embedding using ollama
	// 1) Check of text param to see if it is valid
	if text == "" || text == " " {
		return nil, domain.ErrEmptyInput
	}
	// 2) Create embeddings request
	embeddingRequestPayload := domain.NewEmbeddingRequestPayload()
	embeddingRequestPayload.Model = es.Model
	embeddingRequestPayload.Input = text

	// Marshaling the go structure to obtain JSON
	eb, err := json.Marshal(embeddingRequestPayload)
	if err != nil {
		return nil, err
	}
	// Transforming in bytes
	beb := bytes.NewBuffer(eb)
	resp, err2 := http.Post(es.Service, "application/json", beb)
	if err2 != nil {
		return nil, err2
	}
	defer resp.Body.Close()
	var embeddings domain.EmbeddingResponse
	embeddingBytes, err3 := io.ReadAll(resp.Body)
	if err3 != nil {
		return nil, err3
	}
	err4 := json.Unmarshal(embeddingBytes, &embeddings)
	if err4 != nil {
		return nil, err4
	}

	return &embeddings, nil
}
