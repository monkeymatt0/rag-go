package integration

import (
	"customrag/internal/adapters"
	"testing"
)

const(
	testHost = "localhost"
	testPort = 6334
)

func setupTestVectorRepository(t *testing.T) (*adapters.VectorRepository, error) {
	return adapters.NewVectorRepository(testHost, testPort)
}

// TODO: Write integration tests