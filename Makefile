ifdef update
	u=-u
endif

.PHONY: deps

deps:
	go mod tidy

.PHONY: devel-deps

devel-deps:
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: test

test: deps devel-deps
	docker-compose up -d
	goimports -l -w .
	go test -v -cover ./...

.PHONY: install

install: deps
	go install ./cmd/tinysearch
