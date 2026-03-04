package bootstrap

import (
	"customrag/internal/configuration"
	"customrag/internal/core/ports/inputs"
)

type App struct {
	Configs configuration.Configuration

	EmbeddingService inputs.EmbeddingInterface

	// VectorService
	VectorService inputs.VectorInterface
}

func NewApp() *App {
	return &App{}
}
