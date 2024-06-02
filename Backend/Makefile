
# Docker
docker-create:
	@docker run --name AceMyCareerDB -e POSTGRES_PASSWORD=PasswordYangSusah -p 8092:5432 -d postgres

docker-start:
	@docker start AceMyCareerDB

build:
	@go build -o ./bin/api ./src/main.go

run: build
	@./bin/api

