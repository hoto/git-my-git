.PHONY: clean dependencies build test run install github-release github-release-dry-run goreleaser-dry-run

REPO_NAME = github.com/hoto/git-my-git

clean:
	go clean
	rm -rf bin/ dist/

dependencies:
	go mod download
	go mod tidy
	go mod verify

build: dependencies
	go build -ldflags="-X '${REPO_NAME}/pkg/config.Version=0.0.0' -X '${REPO_NAME}/pkg/config.ShortCommit=HASH' -X '${REPO_NAME}/pkg/config.BuildDate=DATE'" -o bin/git-my-git ./cmd/git-my-git/main.go

test:
	go test -v ./...

run: clean build
	./bin/git-my-git $(arg)

install: clean build
	go install -v ./...

goreleaser-release: clean dependencies
	curl -sL https://git.io/goreleaser | VERSION=v0.137.0 bash

goreleaser-dry-run: clean dependencies
	curl -sL https://git.io/goreleaser | VERSION=v0.137.0 bash -s -- --skip-publish --snapshot --rm-dist

goreleaser-dry-run-local: dependencies
	goreleaser release --skip-publish --snapshot --rm-dist

