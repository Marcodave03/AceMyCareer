build:
	@go build -o bin/rentaltools

run: build
	@./bin/rentaltools

test:
	@go test -v ./...

postgres-cli:
	# @docker exec -it rentaltool-database psql -U postgres -d postgres --password
	@docker exec -e PGPASSWORD='#rEnt172635' -it rentaltool-database psql -U postgres -d postgres
