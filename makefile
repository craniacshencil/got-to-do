include .env
run:
	@go run cmd/main.go

up:
	@cd sql/schema && \
	goose postgres ${GOOSE_URL} up

down:
	@cd sql/schema && \
	goose postgres ${GOOSE_URL} down

login:
	@bash scripts/login.sh

create:
	@bash scripts/createList.sh

display:
	@bash scripts/displayList.sh

add:
	@bash scripts/addList.sh

reset-todos:
	@psql -U ${DB_USER} -d ${DB_NAME} -h ${DB_HOST} -p ${DB_PORT} -f scripts/deleteTodos.sql
