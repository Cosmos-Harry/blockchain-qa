.PHONY: build clean test

build:
	@echo "Building poll-cli (without CGO to avoid hang issue)..."
	CGO_ENABLED=0 go build -o poll-cli cmd/poll-cli/main.go
	@echo "Build complete! Binary: ./poll-cli"

clean:
	@echo "Cleaning build artifacts..."
	rm -f poll-cli poll-cli-*
	@echo "Clean complete!"

test:
	@echo "Running tests..."
	go test ./...

install: build
	@echo "Installing poll-cli to ~/go/bin..."
	cp poll-cli ~/go/bin/
	@echo "Installed! You can now run 'poll-cli' from anywhere"
