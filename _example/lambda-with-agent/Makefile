.PHONY: build
build:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap main.go
	zip archive.zip bootstrap
