package integration

import (
	"context"
	"customrag/internal/adapters"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testText   = "Hi my name is Marcolino and I want to test the eeeeembedding"
	model      = "nomic-embed-text"
	dim        = 768
	serviceURL = "http://localhost:11434/api/embed"
)

func testMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestGenerateEmbedding_Success(t *testing.T) {
	e := adapters.NewEmbeddingService(model, dim)
	e.Model = model
	e.Dim = dim
	e.Service = serviceURL

	ctx := context.Background()
	_, err := e.GenerateEmbeddings(ctx, testText)
	require.NoError(t, err, "Error during the embedding...")
}
