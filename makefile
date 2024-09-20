include .env
run:
	@go run cmd/main.go

up:
	@cd sql/schema && \
	goose postgres ${GOOSE_URL} up

down:
	@cd sql/schema && \
	goose postgres ${GOOSE_URL} down

signup:
	@cd scripts && bash signup.sh

login:
	@cd scripts && bash login.sh

create:
	@cd scripts && bash createList.sh

display:
	@cd scripts && bash displayList.sh

add:
	@cd scripts && bash addList.sh

update:
	@cd scripts && bash updateList.sh

un-update:
	@cd scripts && bash un-updateList.sh

reset-todos:
	@psql -U ${DB_USER} -d ${DB_NAME} -h ${DB_HOST} -p ${DB_PORT} -f scripts/deleteTodos.sql
