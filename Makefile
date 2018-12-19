# Go parameters

GOCMD=go
GOBIN=${GOPATH}/bin
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=gochain
BINARY_UNIX=$(BINARY_NAME)_unix


build:
	@echo "  >  Building binary..."
	@echo "  >  GOPATH: ${GOPATH}"
	$(GOBUILD) -o $(GOBIN)/$(BINARY_NAME) -v *.go
