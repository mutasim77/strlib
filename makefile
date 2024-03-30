run-main:
	@echo "Running Main..."
	cd cmd && go run main.go

run-tests:
	@echo "Running Tests..."
	go test ./... -v