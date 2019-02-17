PKG := github.com/syllabix/versioner
MAIN := cmd/cli/*.go

BUILD_VERSION=`git describe --tags`

# Expose compile-time information as linker symbols
LDFLAGS += -X 'github.com/syllabix/versioner/internal.diagnostic.AppVersion=${BUILD_VERSION}'
LDFLAGS += -X 'github.com/syllabix/versioner/internal.diagnostic.BuildTimestamp=`date +%Y-%m-%dT%H:%M:%S%:z`'
LDFLAGS += -X 'github.com/syllabix/versioner/internal.diagnostic.CommitHash=`git rev-parse HEAD`'
LDFLAGS += -X 'github.com/syllabix/versioner/internal.diagnostic.GoVersion=`go version`'
# Omit symbol table and debug info, leads to a smaller binary
LDFLAGS += -s

GO_BUILD = CGO_ENABLED=0 go build -a -ldflags "$(LDFLAGS)"

.PHONY: help build test clean

## Print the help message.
# Parses this Makefile and prints targets that are preceded by "##" comments.
help:
	@echo "" >&2
	@echo "Available targets: " >&2
	@echo "" >&2
	@awk -F : '\
			BEGIN { in_doc = 0; } \
			/^##/ && in_doc == 0 { \
				in_doc = 1; \
				doc_first_line = $$0; \
				sub(/^## */, "", doc_first_line); \
			} \
			$$0 !~ /^#/ && in_doc == 1 { \
				in_doc = 0; \
				if (NF <= 1) { \
					next; \
				} \
				printf "  %-15s %s\n", $$1, doc_first_line; \
			} \
			' <"$(abspath $(lastword $(MAKEFILE_LIST)))" \
		| sort >&2
	@echo "" >&2

## Build binary using the go link tool to set dianostic variables
build:
	$(mkdir -p ./build)
	$(GO_BUILD) -o ./.build/versioner $(PKG)/cmd/cli

## Build and run the binary
exec:
	$(mkdir -p ./build)
	$(GO_BUILD) -o ./.build/versioner $(PKG)/cmd/cli
	.build/versioner

## Run all tests.
test:
	go test -race -cover -coverprofile=coverage.out ./...

## Removes generated files
clean:
	rm -f ./coverage.out
	rm -rf ./.build