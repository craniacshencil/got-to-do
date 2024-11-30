include .env.local
run:
	@go run cmd/main.go

build:
	@go build -o bin/got_to_do cmd/main.go 

docker-build:
	@docker build -t got_to_do:final .

docker-run:
	@docker run -dp 127.0.0.1:8080:8080 --add-host=host.docker.internal:host-gateway got_to_do:final

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

create-delete:
	@cd scripts && bash createDeleteList.sh

display:
	@cd scripts && bash displayList.sh

add:
	@cd scripts && bash addList.sh

update:
	@cd scripts && bash updateList.sh

un-update:
	@cd scripts && bash un-updateList.sh

delete-task:
	@cd scripts && bash deleteTask.sh

delete-list:
	@cd scripts && bash deleteList.sh

reset-todos:
	@psql -U ${DB_USER} -d ${DB_NAME} -h ${DB_HOST} -p ${DB_PORT} -c "DELETE FROM todo_lists;"


