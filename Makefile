#!/usr/bin/make -f

NAME:=carbon-schema-to-clickhouse

GO ?= go

all: $(NAME)

$(NAME):
	$(GO) get
	$(GO) build

clean:
	rm -f $(NAME)
test: $(NAME)
	go vet *.go

.PHONY: clean test
