bin/sts:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/graph ./cmd/

.PHONY: dep
dep:
	go mod tidy
