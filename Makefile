# Define the target
coverage:
	@echo "Running tests and generating coverage report..."
	@go test ./... -coverprofile=coverage.out > /dev/null
	@echo "Calculating total coverage..."
	@go tool cover -func=coverage.out | grep total | awk '{print "Average Coverage:", $$3}'
	@echo "Generating HTML coverage report..."
	@go tool cover -html=coverage.out -o coverage.html
	@echo "HTML coverage report saved as coverage.html"
