package configuration

type Configuration struct {
	EmbeddingModel    string
	EmbeddingModelDim int
	LLMModel          string

	OllamaLocalEmbed string
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}
