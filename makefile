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
