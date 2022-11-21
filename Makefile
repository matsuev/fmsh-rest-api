# Set path to ./cmd/{app_path} folder
app_path := fmshapi

# Set build application name
app_name := apiserver

.PHONY: run
run:
	@go run ./cmd/$(app_path)

.PHONY: build
build:
	go build -o ./build/$(app_name) -v ./cmd/$(app_path)

.PHONY: tidy
tidy:
	go mod tidy

.DEFAULT_GOAL := run