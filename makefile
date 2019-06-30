EXEC=./monkey-go
SRC=$(shell find . -type f -name "*.go")

all: build

run: $(EXEC)
	$(EXEC)

$(EXEC): $(SRC)
	go build

build: $(EXEC)

test: build
	@go test ./...

clean: 
	@go clean --testcache

.PHONY: all run clean test