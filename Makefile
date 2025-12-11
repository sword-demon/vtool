# Makefile for vtool project

.PHONY: all build test lint fmt vet check clean help

# 默认目标
all: test lint

# 构建项目
build:
	@echo "Building vtool..."
	go build ./...

# 运行所有测试
test:
	@echo "Running tests..."
	go test -v ./...

# 运行测试并生成覆盖率报告
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# 运行基准测试
bench:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

# 代码格式化
fmt:
	@echo "Formatting code..."
	go fmt ./...

# 代码检查
vet:
	@echo "Running go vet..."
	go vet ./...

# 运行 golangci-lint
lint:
	@echo "Running golangci-lint..."
	@if ! command -v golangci-lint > /dev/null 2>&1; then \
		echo "golangci-lint not installed. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	golangci-lint run

# 运行 golangci-lint 并修复可修复的问题
lint-fix:
	@echo "Running golangci-lint with --fix..."
	@if ! command -v golangci-lint > /dev/null 2>&1; then \
		echo "golangci-lint not installed. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	golangci-lint run --fix

# 检查代码质量（包含所有检查）
check: fmt vet lint
	@echo "All checks passed!"

# 清理构建产物
clean:
	@echo "Cleaning..."
	go clean ./...
	rm -f coverage.out coverage.html

# 安装依赖
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

# 安装开发工具
devtools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/swaggo/swag/cmd/swag@latest

# 生成文档
docs:
	@echo "Generating documentation..."
	@if command -v swag > /dev/null 2>&1; then \
		swag init; \
	else \
		echo "swag not installed. Install with: go install github.com/swaggo/swag/cmd/swag@latest"; \
	fi

# 运行特定包的测试
test-package:
	@echo "Running tests for package: $(PKG)"
	go test -v ./$(PKG)

# 运行特定包的 lint
lint-package:
	@echo "Running lint for package: $(PKG)"
	@if command -v golangci-lint > /dev/null 2>&1; then \
		golangci-lint run ./$(PKG); \
	else \
		echo "golangci-lint not installed. Run 'make devtools' first."; \
	fi

# 检查代码复杂度
cyclo:
	@echo "Checking cyclomatic complexity..."
	@if command -v gocyclo > /dev/null 2>&1; then \
		gocyclo -over 15 ./...; \
	else \
		echo "gocyclo not installed. Install with: go install github.com/fzipp/gocyclo/cmd/gocyclo@latest"; \
	fi

# 检查重复代码
dupl:
	@echo "Checking for duplicated code..."
	@if command -v dupl > /dev/null 2>&1; then \
		dupl -threshold 100 ./...; \
	else \
		echo "dupl not installed. Install with: go install github.com/mibk/dupl@latest"; \
	fi

# 安全检查
security:
	@echo "Running security checks..."
	@if command -v gosec > /dev/null 2>&1; then \
		gosec ./...; \
	else \
		echo "gosec not installed. Install with: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest"; \
	fi

# 性能分析
profile:
	@echo "Running CPU profile..."
	go test -cpuprofile=cpu.prof -bench=. ./...
	go tool pprof cpu.prof

# 内存分析
memprofile:
	@echo "Running memory profile..."
	go test -memprofile=mem.prof -bench=. ./...
	go tool pprof mem.prof

# 交叉编译
build-linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build -o bin/vtool-linux-amd64 ./...

build-windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -o bin/vtool-windows-amd64.exe ./...

build-darwin:
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=amd64 go build -o bin/vtool-darwin-amd64 ./...

build-all: clean build-linux build-windows build-darwin
	@echo "All binaries built in bin/"

# 显示帮助信息
help:
	@echo "Available targets:"
	@echo "  all           - Run all checks (default)"
	@echo "  build         - Build the project"
	@echo "  test          - Run all tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  bench         - Run benchmarks"
	@echo "  fmt           - Format code"
	@echo "  vet           - Run go vet"
	@echo "  lint          - Run golangci-lint"
	@echo "  lint-fix      - Run golangci-lint with --fix"
	@echo "  check         - Run all checks (fmt, vet, lint)"
	@echo "  clean         - Clean build artifacts"
	@echo "  deps          - Install dependencies"
	@echo "  devtools      - Install development tools"
	@echo "  docs          - Generate documentation"
	@echo "  test-package  - Run tests for specific package (PKG=...)"
	@echo "  lint-package  - Run lint for specific package (PKG=...)"
	@echo "  cyclo         - Check cyclomatic complexity"
	@echo "  dupl          - Check for duplicated code"
	@echo "  security      - Run security checks"
	@echo "  profile       - Run CPU profiling"
	@echo "  memprofile    - Run memory profiling"
	@echo "  build-linux   - Build for Linux"
	@echo "  build-windows - Build for Windows"
	@echo "  build-darwin  - Build for macOS"
	@echo "  build-all     - Build for all platforms"
	@echo "  help          - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make test-package PKG=internal/bean"
	@echo "  make lint-package PKG=internal/slice"
	@echo "  make build-all"