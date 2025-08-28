.PHONY: generate verify test benchmark
generate:
	go run ./internal/generator

verify:
	go run ./internal/verifier

test:
	go test -v ./...

benchmark:
	go test -bench=. -benchtime=1s ./...
