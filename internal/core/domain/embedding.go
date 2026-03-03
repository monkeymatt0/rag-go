package domain

import (
	"context"
	"errors"
	"net/url"
)

var (
	ErrEmptyInput = errors.New("Empty input: write something valid")
)

type EmbeddingService interface {
	GenerateEmbeddings(ctx context.Context, text string) ([]float32, error)
}

// This data structure will represent the vector.
type Vector struct{}

// Embedding request has all the needed attribute to perform an embedding request
// url -> is a private attribute
// model -> is private for now, in future could became selectable
type EmebeddingRequest struct {
	url   url.URL
	model string
	Input string
}

func NewEmbeddingRequest() *EmebeddingRequest {
	return &EmebeddingRequest{}
}
