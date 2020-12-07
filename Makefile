.PHONY: test
test:
	go test -race ./...

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/testing-github-action
