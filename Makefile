.PHONY: lint test vendor clean

export GO111MODULE=on

default: lint test

lint:
	golangci-lint run

test:
	go test -v -cover ./...

yaegi_test:
	yaegi test .

vendor:
	go mod vendor

clean:
	rm -rf ./vendor

tools: install-golangci install-yaegi

install-golangci:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1

install-yaegi:
	go install github.com/traefik/yaegi/cmd/yaegi@v0.10.0
