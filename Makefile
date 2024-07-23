build:
	@go build -o bin/goBackend cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/goBackend