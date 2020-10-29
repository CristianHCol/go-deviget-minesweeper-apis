SHELL:=/bin/bash -O extglob
BINARY=ms-api
VERSION=0.1.0
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

#go tool commands
build:
	go build ${LDFLAGS} -o ${BINARY} cmd/ms/main.go

run:
	@go run cmd/ms/main.go

run-lint:
	./.bin/golangci-lint run

install-linter:
	mkdir -p ./.bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./.bin

unit-test:
	go test -race ./internal/... -count 1 --short

print-coverage:
	go test -coverprofile fmtcoverage.html fmt 

test-coverage-report:
	go test -covermode=count -coverprofile test_gen.out $$(go list ./internal/... | grep -v /test/) && go tool cover -html=test_gen.out

test-converage:
	go test -covermode=count -coverprofile test_gen.out ./internal/... && go tool cover -func=test_gen.out | grep total

update-deps:
	go get -u ./...

cleanup-deps:
	go mod tidy
## docker compose
up:
	docker-compose up --build
down:
	docker-compose down --remove-orphans