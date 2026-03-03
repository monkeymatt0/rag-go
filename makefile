.PHONY: pull_model
pull_model:
	docker exec internal_ollama ollama pull $(model)

.PHONY: list_models
list_models:
	docker exec internal_ollama ollama list