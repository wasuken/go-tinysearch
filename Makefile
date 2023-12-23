ifdef update
	u=-u
endif

export GO111MODULE=on

.PHONY: deps

deps:
	go get ${u} -d
	go mod tidy

.PHONY: devel-deps

devel-deps:
	GO111MODULE=off go get ${u} \
							golang.org/x/tools/cmd/goimports

.PHONY: test

test:
	deps devel-deps
		docker-compose up -d
		goimports -l -w .
		go test j-v -cover ./...

.PHONY: install

install: deps
	go install ./cmd/tinysearch
