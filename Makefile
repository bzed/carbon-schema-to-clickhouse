#!/usr/bin/make -f

NAME:=carbon-schema-to-clickhouse

GO ?= go
export GOPATH := $(CURDIR)/_vendor


all: $(NAME)

$(NAME):
	$(GO) get
	$(GO) build

clean:
	rm -f $(NAME)
test: $(NAME)
	go vet *.go

.PHONY: clean test
