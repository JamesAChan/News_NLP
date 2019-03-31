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
	$(GOBUILD) -o ./bin/rpc_server -v ./rpc/server/main.go
	$(GOBUILD) -o ./bin/mds_record -v ./db/mds/record/main.go
	$(GOBUILD) -o ./bin/mds_recordFut -v ./db/mds/recordFut/main.go
	$(GOBUILD) -o ./bin/mds_recordOpt -v ./db/mds/recordOpt/main.go

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f ./bin/*
