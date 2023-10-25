format:
	gofmt -w -l -s .

lint:
	golangci-lint run
