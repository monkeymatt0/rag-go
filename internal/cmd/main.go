package main

import (
	"customrag/internal/adapters"
	"customrag/internal/bootstrap"
	"customrag/internal/configuration"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	// DI:
	// 1) Create App
	app := bootstrap.NewApp()
	// 2) Load configs
	configs := configuration.NewConfiguration()
	// 3) read .env vars
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error while reading configs...")
	}
	// 4) Assign env vars to the configs
	configs.EmbeddingModel = os.Getenv("EMBEDDING_MODEL")
	emd, errDim := strconv.ParseInt(os.Getenv("EMBEDDING_MODEL_DIM"), 10, 64)
	if errDim != nil {
		fmt.Println("Failed to retrieve embedding dimension...")
	}
	configs.EmbeddingModelDim = int(emd)
	configs.LLMModel = os.Getenv("LLM_MODEL")
	configs.OllamaLocalEmbed = os.Getenv("OLLAMA_LOCAL_EMBED")

	// 5) Creating needed services
	host := os.Getenv("QDRANT_HOST")
	port, errPort := strconv.ParseInt(os.Getenv("QDRANT_PORT"), 10, 64)
	if errPort != nil {
		fmt.Println("Failed to load the port")
	}
	vectorService, errVec := adapters.NewVectorRepository(host, int(port))
	if errVec != nil {
		fmt.Println("Failed to create vector service")
	}

	app.VectorService = vectorService

	// 6) Assign configs to the app
	app.Configs = *configs

	// 7) TODO: Start the server
}
