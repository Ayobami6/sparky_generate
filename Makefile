build:
	@go build -o ./bin/sparky

run: build
	@./bin/sparky