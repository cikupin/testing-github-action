.PHONY: test
test:
	go test -race ./...

.PHONY: build
build:
	go build -ldflags="-w -s" -o bin/testing-github-action

.PHONY: build-for-scratch
build-for-scratch:
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o bin/testing-github-action

.PHONY: image-scratch
image-scratch: build-for-scratch
	docker build -t cikupin/testing-github-action:latest .
