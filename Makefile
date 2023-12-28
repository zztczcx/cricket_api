.PHONY: cli
cli:
	docker compose run --rm dev-cli

.PHONY: import-csv
import-csv:
	docker compose run --rm dev-cli go run cmd/load_csv_to_db/main.go
