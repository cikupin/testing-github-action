.PHONY: test
test:
	go test -race ./...

.PHONY: build
build:
	GOOS=linux go build -ldflags="-w -s" -o bin/testing-github-action

.PHONY: image-scratch
image-scratch:
	GOOS=linux go build -ldflags="-w -s" -o bin/testing-github-action
	docker build -t cikupin/testing-github-action:latest .
