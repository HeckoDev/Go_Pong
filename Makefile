.PHONY: build run clean install-deps help

BINARY_NAME=pong
MAIN_FILE=main.go

all: build

build:
	@echo "🔨 Building Pong Master..."
	@go build -o $(BINARY_NAME) $(MAIN_FILE)
	@echo "✅ Build complete: $(BINARY_NAME)"

run: build
	@echo "🎮 Starting Pong Master..."
	@echo "📋 Controls:"
	@echo "   Player 1: W (up) / S (down)"
	@echo "   Player 2: ↑ (up) / ↓ (down)"
	@echo "   Enter: Launch ball"
	@echo ""
	@./$(BINARY_NAME)

dev:
	@echo "🎮 Running in development mode..."
	@go run $(MAIN_FILE)

clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -f $(BINARY_NAME) $(BINARY_NAME)-* $(BINARY_NAME).exe
	@echo "✅ Clean complete"

deps:
	@echo "📦 Installing Go dependencies..."
	@go mod tidy
	@go mod download
	@echo "✅ Dependencies installed"

install-deps:
	@echo "📦 Installing system dependencies..."
	@sudo apt-get update
	@sudo apt-get install -y libx11-dev libxcursor-dev libxrandr-dev \
		libxinerama-dev libxi-dev libgl1-mesa-dev libxxf86vm-dev \
		libasound2-dev libasound2 libasound2-plugins alsa-utils
	@echo "✅ System dependencies installed"

build-linux:
	@echo "🐧 Building for Linux..."
	@GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux $(MAIN_FILE)
	@echo "✅ Created: $(BINARY_NAME)-linux"

build-windows:
	@echo "🪟 Building for Windows..."
	@GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe $(MAIN_FILE)
	@echo "✅ Created: $(BINARY_NAME).exe"

build-mac:
	@echo "🍎 Building for macOS..."
	@GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-mac $(MAIN_FILE)
	@echo "✅ Created: $(BINARY_NAME)-mac"

build-all: build-linux build-windows build-mac
	@echo "✅ Multi-platform build complete"

help:
	@echo "Makefile for Pong Master"
	@echo ""
	@echo "Available commands:"
	@echo "  make build         - Build the game"
	@echo "  make run           - Build and run the game"
	@echo "  make dev           - Run in development mode (go run)"
	@echo "  make clean         - Remove build artifacts"
	@echo "  make deps          - Install Go dependencies"
	@echo "  make install-deps  - Install system dependencies (Linux)"
	@echo "  make build-linux   - Cross-compile for Linux"
	@echo "  make build-windows - Cross-compile for Windows"
	@echo "  make build-mac     - Cross-compile for macOS"
	@echo "  make build-all     - Build for all platforms"
	@echo "  make help          - Show this help message"
