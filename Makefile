.PHONY: run
run:
	go run ./cmd/fmshapi

.PHONY: build
build:
	go build -o ./build/fmshapi -v ./cmd/fmshapi

.PHONY: tidy
tidy:
	go mod tidy