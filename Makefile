BIN = docker-ibug

GO ?= go
LDFLAGS = -s -w

.PHONY: all gzip test $(BIN)

all: $(BIN)

gzip: $(BIN).gz

test:
	$(GO) test ./...

$(BIN):
	$(GO) build -ldflags="$(LDFLAGS)" .

$(BIN).gz: $(BIN)
	gzip -k9 $^
