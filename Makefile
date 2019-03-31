# Go parameters
GOCMD=go
GOLINT=golint
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BUILD_TARGET_MAIN=main
#BUILD TARGET

all: build

.PHONY: fmt
fmt:
	$(GOCMD) fmt ./...

.PHONY: test
test:
	$(GOTEST) -short ./...

.PHONY: public
public:
	$(GOBUILD) -o ./bin/google -v ./integrated/googlenews/example/main.go

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f ./bin/*
