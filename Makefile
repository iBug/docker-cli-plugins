GO ?= go
LDFLAGS = -s -w

.PHONY: all

all:
	$(GO) build -ldflags="$(LDFLAGS)" .
