# Makefile

# The binary to build (just the basename).
BIN := qscrapper

# Set this to your actual build and output directory
BUILD_DIR := ./build

build:
	@echo "Building $(BIN) version"
	@go build -o $(BUILD_DIR)/$(BIN) cmd/main.go

run: build
	@echo "Running $(BIN)"
	@./$(BUILD_DIR)/$(BIN)

clean:
	@echo "Cleaning"
	@rm -rf $(BUILD_DIR)
