.PHONY: default

default: swagger-update mock run

init:
	brew install golangci-lint
	brew install vektra/tap/mockery
	brew upgrade mockery
	go mod download

clean:
	rm -rf ./build
	rm -rf mocks

linter: mock
	golangci-lint run ./...

test: mock
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean swagger-update mock test linter
	CGO_ENABLED=0 go build -ldflags="-w -s"

mock:
	go generate ./...

run:
	go run main.go

mutation-test: mock
	go get -t -v github.com/zimmski/go-mutesting/...
	go-mutesting ./application/...  ./pkg/...
