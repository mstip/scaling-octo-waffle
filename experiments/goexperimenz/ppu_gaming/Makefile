# Go parameters
GO=go

all: test build

build:
	$(GO) build -o cmd/ppu_gaming cmd/main.go

run:
	$(GO) run cmd/main.go

test:
	CGO_CFLAGS="-g -O2 -Wno-return-local-addr" $(GO) test ./...

clean:
	$(GO) clean
	rm cmd/ppu_gaming

install:
	$(GO) mod download
