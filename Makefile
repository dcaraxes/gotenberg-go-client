GOLANG_VERSION=1.23.2
GOTENBERG_VERSION=8.12.0
GOTENBERG_LOG_LEVEL=ERROR
VERSION=snapshot
GOLANGCI_LINT_VERSION=1.61.0
REPO=dcaraxes/gotenberg-go-client

# gofmt and goimports all go files.
fmt:
	go fmt ./...
	go mod tidy

# run all linters.
lint:
	docker build --build-arg GOLANG_VERSION=$(GOLANG_VERSION) --build-arg GOLANGCI_LINT_VERSION=$(GOLANGCI_LINT_VERSION) -t $(REPO):lint -f build/lint/Dockerfile .
	docker run --rm -it -v "$(PWD):/lint" $(REPO):lint

# run all tests.
tests:
	docker build --build-arg GOLANG_VERSION=$(GOLANG_VERSION) --build-arg GOTENBERG_VERSION=$(GOTENBERG_VERSION) --build-arg GOTENBERG_LOG_LEVEL=$(GOTENBERG_LOG_LEVEL) -t $(REPO):tests -f build/tests/Dockerfile .
	docker run --rm -it -v "$(PWD):/tests" $(REPO):tests