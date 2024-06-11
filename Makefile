TMP_COVER_FILE="/tmp/go-coverage.tmp"

## Quality
code-quality: ## runs code quality checks
	make lint
	make fmt
	make vet

# Append || true below if blocking local developement
lint: ## go linting. Update and use specific lint tool and options
	golangci-lint run --enable-all

vet: ## go vet
	go vet ./...

fmt: ## runs go formatter
	go fmt ./...

## Test
test: ## runs tests and create generates coverage report
	@echo ">>> Running tests with coverage"
	#go test -cover ./...
	go test -covermode=count -timeout 10m ./... -coverprofile=$(TMP_COVER_FILE)

html-coverage: ## displays test coverage report in html mode
	make test
	@echo ">>> Generating html coverage report"
	go tool cover -html=$(TMP_COVER_FILE)