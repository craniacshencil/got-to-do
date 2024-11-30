# got-to-do

- todo app written in go, requires postgres.
- utils used are bash, make and curl
- also dockerized
- super impractical, made this to learn go, curl, make and docker

## Local Dev Setup

- Make sure you have make and curl installed

- Create database on postgres

```sql
CREATE DATABASE got_to_do;
```

- For setting up database create `.env.local` file for local development (in root project directory)

```sh
PORT=8080
SECRET_JWT_KEY=<some_random_string>
GOOSE_URL=postgres://<username>:<password>@localhost:<postgres_port>/got_to_do
DB_URL=postgres://<username>:<password>@localhost:<postgres_port>/got_to_do?sslmode=disable
DB_NAME=got_to_do
DB_HOST=localhost
DB_USER=<username> #(can be postgres)
DB_PORT=<postgres_port> #(usually 5432)
```

- Run migrations on database

```sh
make up
```

- Run `setup.sh` to seed database

```sh
bash setup.sh
```

- Run the following make commands in order (one by one would be better)

```sh
make signup
make login
make create
```

- Other functionalites are also present in the `makefile`

## Docker setup

- The first 5 steps are common from above, perform them first

- Create separate `.env.docker` file for docker

```sh
PORT=8080
SECRET_JWT_KEY=<some_random_string>
GOOSE_URL=postgres://<username>:<password>@host.docker.internal:<postgres_port>/got_to_do
DB_URL=postgres://<username>:<password>@host.docker.internal:<postgres_port>/got_to_do?sslmode=disable
DB_NAME=got_to_do
DB_HOST=host.docker.internal
DB_USER=<username>
DB_PORT=<postgres_port>
```

Here the `username`, `postgres_port`, `password` would be same as that of the `.env.local` file. We are accessing the database on your machine and not on a postgres container.

- Create the binary for the app

```sh
make build
```

- Create docker image

```sh
make docker-build
```

- Run the docker container

```sh
make docker-run
```

- Now run the scripts mentioned above (last 2 steps of local dev setup)

### NOTE

- Configure postgres to accept connections from docker
