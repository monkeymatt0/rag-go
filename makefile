.PHONY: pull_model
pull_model:
	docker exec internal_ollama ollama pull $(model)

list_models:
	docker exec internal_ollama ollama list