run:
	@go run main.go
build:
	@echo "Building zero..."
	@mkdir ./.build
	@go build -o ./.build/zero
	@sudo mv ./.build/zero /usr/local/bin/zero
	@rm -d ./.build
	@echo "Build complete."