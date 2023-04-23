GOPATH_DIR=`go env GOPATH`

.PHONY: test
test:
	go test -count 2 -v -race ./...

.PHONY: lint
lint:
	@which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1)
	$(GOPATH_DIR)/bin/golangci-lint run -v --deadline 30s ./...
	@echo "everything is OK"