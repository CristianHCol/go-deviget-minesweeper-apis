run-lint:
	./.bin/golangci-lint run

install-linter:
	mkdir -p ./.bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./.bin

unit-test:
	go test -race ./internal/... -count 1 --short

# test folder is excluded since there will be placed the integrtion test and free gitlab ci
# does not support mock dbs or external componenets
unit-test-ci:
	go test -race $$(go list ./internal/... | grep -v /test/) -count 1

# race no running because causes timeouts
integration-test:
	go test ./internal/test/... -count 1

print-coverage:
	go test -coverprofile fmtcoverage.html fmt 

# race no running because causes timeouts
full-test-suite:
	go test ./internal/... -count 1

test-specific:
	go test -race -run $(TEST_NAME) ./internal/... -count 1

test-specific-folder:
	go test -race -run $(TEST_NAME) $(FOLDER) -count 1

test-coverage-report:
	go test -covermode=count -coverprofile test_gen.out $$(go list ./internal/... | grep -v /test/) && go tool cover -html=test_gen.out

test-converage-ci:
	go test -covermode=count -coverprofile coverage.cov $$(go list ./internal/... | grep -v /test/) && go tool cover -func=coverage.cov

test-converage:
	go test -covermode=count -coverprofile test_gen.out ./internal/... && go tool cover -func=test_gen.out | grep total

update-deps:
	go get -u ./...

cleanup-deps:
	go mod tidy

apply-migration:
	go run ./cmd/migration/ --action=upgrade --scriptPath=file://resources/$(SERVICE_NAME)/database/ --dbUrl=$(MIGRATION_DB)

undo-migration:
	go run ./cmd/migration/ --action=downgrade --scriptPath=file://resources/$(SERVICE_NAME)/database/ --dbUrl=$(MIGRATION_DB)
