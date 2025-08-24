.PHONY: help build run clean docker-up docker-down migrate test

help: ## 显示帮助信息
	@echo "可用命令:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## 编译程序
	@echo "编译程序..."
	go build -o bin/wallet-manager cmd/main.go

run: ## 运行程序
	@echo "启动钱包管理服务..."
	go run cmd/main.go

clean: ## 清理编译文件
	@echo "清理编译文件..."
	rm -rf bin/

docker-up: ## 启动Docker服务（数据库）
	@echo "启动Docker服务..."
	docker-compose up -d

docker-down: ## 停止Docker服务
	@echo "停止Docker服务..."
	docker-compose down

docker-logs: ## 查看Docker日志
	docker-compose logs -f

migrate: ## 运行数据库迁移
	@echo "运行数据库迁移..."
	@echo "数据库迁移将在服务启动时自动运行"

test: ## 运行测试
	@echo "运行测试..."
	go test -v ./...

dev: docker-up ## 开发环境启动
	@echo "等待数据库启动..."
	sleep 5
	@echo "启动开发服务器..."
	GIN_MODE=debug go run cmd/main.go

install: ## 安装依赖
	@echo "安装Go依赖..."
	go mod tidy
	go mod download

fmt: ## 格式化代码
	@echo "格式化代码..."
	go fmt ./...

lint: ## 代码检查
	@echo "代码检查..."
	golangci-lint run

.DEFAULT_GOAL := help
