# Go 命令和项目配置
GO := go
PROJECT_NAME := GINCHAT
BIN_NAME := $(GINCHAT)
BUILD_DIR := ./bin

# Swag 文档生成配置
SWAG_DOCS_DIR := ./docs
SWAG_MAIN := ./main.go  # 主文件路径（根据项目调整）

# Wire 依赖注入配置
WIRE_DIR := ./cmd/app      # Wire 配置文件所在目录（根据项目调整）

.PHONY: all build run test clean install-tools swag wire

# 默认目标：构建项目
all: build

# 安装必要的工具（Swag 和 Wire）
install-tools:
	@echo "Installing tools..."
	$(GO) install github.com/swaggo/swag/cmd/swag@latest
	$(GO) install github.com/google/wire/cmd/wire@latest

# 生成 Swagger 文档
swag:
	@echo "Generating Swagger docs..."
	swag init -g $(SWAG_MAIN) -o $(SWAG_DOCS_DIR)

# 生成 Wire 依赖注入代码
#wire:
#	@echo "Generating Wire dependencies..."
#	cd $(WIRE_DIR) && wire

# 构建项目（依赖 Swagger 和 Wire）
build:
	@echo "Building binary..."
	$(GO) build -v -o $(BUILD_DIR)/$(BIN_NAME) ./...

# 运行项目
run: build
	@echo "Starting application..."
	$(GO) run main.go

# 执行测试
test:
	@echo "Running tests..."
	$(GO) test -v -race ./...

# 清理生成文件和二进制
clean:
	@echo "Cleaning up..."
	rm -rf $(SWAG_DOCS_DIR)
	rm -f $(BUILD_DIR)/$(BIN_NAME)
	find . -name "wire_gen.go" -delete