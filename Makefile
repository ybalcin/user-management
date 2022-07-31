.PHONY: run-tests
run-tests:
	go test ./internal/application/test ./internal/domain/test

.PHONY: run-api
run-api:
	go run main.go