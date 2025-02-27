build:
	@go build -o ./bin/sparky

run: build
	@./bin/sparky

build-mac:
	@go build -o ./bin/mac/sparky

run-mac: build-mac
	@./bin/mac/sparky