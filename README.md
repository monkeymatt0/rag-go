# rag-go
Rag system built in golang.

## Compose
In the compose file we have qdrant service and ollama service.
Ollama in this case is thought as a service to use both LLM and produce embeddings.

To act directly on the container of ollama is possible to use the makefile that provides:

- pull_models: Download a model (used for both embeddings model or LLM)
     ```bash
     make pull_models model=<model_name>
     ```
- list_models: Show lists of downloaded models.
     make list_models
     ```bash
     make list_models
     ```