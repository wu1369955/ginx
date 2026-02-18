# 定义变量
APP_NAME=main

# 构建应用
build:
	@echo "Building application..."
	@go build -o $(APP_NAME) .

# 运行应用
run:
	@echo "Running application..."
	@go run .

# 运行测试
test:
	@echo "Running tests..."
	@go test ./tests/...

# 构建Docker镜像
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(APP_NAME) .

# 运行Docker容器
docker-run:
	@echo "Running Docker container..."
	@docker run -p 8080:8080 $(APP_NAME)

# 使用Docker Compose运行
docker-compose-up:
	@echo "Running with Docker Compose..."
	@docker-compose up -d

# 停止Docker Compose服务
docker-compose-down:
	@echo "Stopping Docker Compose services..."
	@docker-compose down

# 清理
tidy:
	@echo "Tidying up..."
	@go mod tidy

# 查看依赖
vendor:
	@echo "Vendorizing dependencies..."
	@go mod vendor

# 格式化代码
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# 检查代码
vet:
	@echo "Vetting code..."
	@go vet ./...

# 帮助
help:
	@echo "Available commands:"
	@echo "  make build           - Build the application"
	@echo "  make run             - Run the application"
	@echo "  make test            - Run tests"
	@echo "  make docker-build    - Build Docker image"
	@echo "  make docker-run      - Run Docker container"
	@echo "  make docker-compose-up - Run with Docker Compose"
	@echo "  make docker-compose-down - Stop Docker Compose services"
	@echo "  make tidy            - Tidy up dependencies"
	@echo "  make vendor          - Vendorize dependencies"
	@echo "  make fmt             - Format code"
	@echo "  make vet             - Vet code"
	@echo "  make help            - Show this help"

.PHONY: build run test docker-build docker-run docker-compose-up docker-compose-down tidy vendor fmt vet help
