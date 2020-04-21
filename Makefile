test:
	go test -v -cover -coverprofile=cover.out ./...

cover:
	go tool cover -html=cover.out

lint:
	golangci-config-generator
	golangci-lint run

install-gcg:
	go install github.com/rbee3u/golangci-config-generator/cmd/golangci-config-generator@latest

install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
