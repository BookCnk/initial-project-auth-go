migrate:
	make lint
	go run cmd/main.go --with-migrate

run:
	make lint
	go run cmd/main.go

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1
	PATH=$(PATH):$(shell go env GOPATH)/bin golangci-lint run

clean:
	go clean -modcache
	go mod tidy

test:
	make lint
	go test -cover ./...
