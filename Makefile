build:
	@go build -o bin/go-ecommerce-api cmd/main.go

test:
	@go  test -v ./...

run: build
	@./bin/go-ecommerce-api