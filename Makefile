.PHONY: build
build:
	go build -o bin/bugsnag-release-monitor .

.PHONY: lint
lint:
	golangci-lint run --enable-all
