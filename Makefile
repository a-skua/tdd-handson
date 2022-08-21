.PHONY: doc fibonacci calculator

doc:
	docker compose run --rm doc

calculator:
	docker compose run --rm -v ./calculator:/go/src gopher go test ./...

fibonacci:
	docker compose run --rm -v ./fibonacci:/go/src gopher go test ./...
