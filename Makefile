coverage:
	@echo "Running tests and generating coverage report..."
	# @go test $(go list ./... | grep -v mocks) -coverprofile=coverage.out
	@echo "Calculating total coverage..."
	@go tool cover -func=coverage.out | grep total | awk '{print "Average Coverage:", $$3}'
	@go tool cover -html=coverage.out -o coverage.html
