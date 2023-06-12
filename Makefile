BIN = docker-ibug

GO ?= go
LDFLAGS = -s -w

.PHONY: all gzip $(BIN)

all: $(BIN)

gzip: $(BIN).gz

$(BIN):
	$(GO) build -ldflags="$(LDFLAGS)" .

$(BIN).gz: $(BIN)
	gzip -k9 $^
