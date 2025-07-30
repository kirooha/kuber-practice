DB_DSN:="postgres://127.0.0.1:5432/kuber_practice?sslmode=disable&timezone=utc"
PSQL_DSN:=$(shell echo $(DB_DSN) | sed 's/&timezone=utc//g')
ROOT_PSQL_DSN:=$(shell echo $(PSQL_DSN) | sed 's/kuber_practice/postgres/g')

regenerate-db: regenerate-db-drop-create db-migrate

regenerate-db-drop-create:
	psql "$(ROOT_PSQL_DSN)" --command="DROP DATABASE IF EXISTS kuber_practice;"
	psql "$(ROOT_PSQL_DSN)" --command="CREATE DATABASE kuber_practice WITH OWNER postgres ENCODING = 'UTF8';"

db-migrate:
	goose -dir db/migrations postgres "$(PSQL_DSN)" up
	make db-gen-structure

db-gen-structure:
	pg_dump "$(PSQL_DSN)" --schema-only --no-owner --no-privileges --no-tablespaces --no-security-labels > db/schema.sql

sqlc:
	sqlc generate