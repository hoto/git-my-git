all: clean build test

clean:
	go clean
	rm -rf bin/ dist/

dependencies:
	go mod download
	go mod tidy
	go mod verify

build: dependencies
	go build -o bin/git-my-git

build-with-ldflags: dependencies
	go build -ldflags="-X 'main.Version=0.0.0' -X 'main.ShortCommit=hash' -X 'main.BuildDate=today'" -o bin/git-my-git

test:
	go test -v ./...

run: clean build
	./bin/git-my-git $(arg)

install: clean build
	go install -v ./...

github-release: clean dependencies
	curl -sL https://git.io/goreleaser | VERSION=v0.137.0 bash

github-release-dry-run: clean dependencies
	curl -sL https://git.io/goreleaser | VERSION=v0.137.0 bash -s -- --skip-publish --snapshot --rm-dist

goreleaser-dry-run: dependencies
	goreleaser release --skip-publish --snapshot --rm-dist

