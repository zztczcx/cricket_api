.PHONY: cli
cli:
	docker compose run --rm dev-cli

.PHONY: import-csv
import-csv:
	docker compose run --rm dev-cli go run cmd/load_csv_to_db/main.go

MYSQLCMD=mysql
ifndef CI
	MYSQLCMD=docker compose exec cricket.db mysql
endif

.PHONY: test-db
test-db:
	@$(MYSQLCMD) -h cricket.db -u root -ppassword -e 'DROP DATABASE IF EXISTS cricket_db_test'
	@$(MYSQLCMD) -h cricket.db -u root -ppassword -e 'CREATE DATABASE cricket_db_test'
	docker compose run --rm cricket.db.migrations 'mysql://root:password@tcp(cricket.db:3306)/cricket_db_test' up


.PHONY: docs
docs:
	docker run --rm \
    -v .:/local openapitools/openapi-generator-cli generate \
    -i /local/docs/openAPI_3.yml \
    -g html \
    -o /local/docs/html/

	docker run --rm \
    -v .:/local openapitools/openapi-generator-cli generate \
    -i /local/docs/openAPI_3.yml \
    -g markdown \
    -o /local/docs/markdown/
