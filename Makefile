build:
	@go build -o bin/api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/api
