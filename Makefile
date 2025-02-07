GOBIN:=$(GOPATH)/bin

install-golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	$(GOBIN)/golangci-lint run ./... --config $(CURDIR)/.golangci.pipeline.yaml