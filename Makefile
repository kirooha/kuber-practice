DB_DSN:="postgres://127.0.0.1:5432/kuber_practice?sslmode=disable&timezone=utc"
PSQL_DSN:=$(shell echo $(DB_DSN) | sed 's/&timezone=utc//g')
ROOT_PSQL_DSN:=$(shell echo $(PSQL_DSN) | sed 's/kuber_practice/postgres/g')

regenerate-db: regenerate-db-drop-create db-migrate db-gen-structure

regenerate-db-drop-create:
	psql "$(ROOT_PSQL_DSN)" --command="DROP DATABASE IF EXISTS kuber_practice;"
	psql "$(ROOT_PSQL_DSN)" --command="CREATE DATABASE kuber_practice WITH OWNER kuber_practice_user ENCODING = 'UTF8';"

db-migrate:
	goose -dir db/migrations postgres "$(PSQL_DSN)" up
	make db-gen-structure

db-gen-structure:
	pg_dump "$(PSQL_DSN)" --schema-only --no-owner --no-privileges --no-tablespaces --no-security-labels > db/schema.sql

sqlc:
	sqlc generate

docker:
	docker build -t eu.gcr.io/my-new-project-467616/kuber_practice_app:latest .
	docker push eu.gcr.io/my-new-project-467616/kuber_practice_app:latest
	docker build -f Dockerfile.redis -t eu.gcr.io/my-new-project-467616/kuber_practice_app_redis:latest .
	docker push eu.gcr.io/my-new-project-467616/kuber_practice_app_redis:latest

run:
	docker run -it --rm --name kuber_practice_app_running_app eu.gcr.io/my-new-project-467616/kuber_practice_app:latest