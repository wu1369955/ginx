# ginx - Go 后端模板项目

一个基于 Go 语言和 Gin 框架的后端模板项目，集成了常用的中间件和工具，提供了完整的项目结构和开发环境配置。

## 功能特性

- ✅ Go 1.23.0 + Gin 框架
- ✅ MySQL 数据库集成（GORM）
- ✅ JWT 认证中间件
- ✅ 请求日志中间件
- ✅ 错误处理中间件
- ✅ CORS 中间件
- ✅ HTTPS 重定向中间件
- ✅ 压缩中间件
- ✅ 限流中间件
- ✅ 优雅关闭
- ✅ 配置管理（Viper）
- ✅ Docker 容器化
- ✅ 完整的项目结构

## 项目结构

```
├── bin/              # 构建产物
├── build/            # 构建脚本
├── configs/          # 配置文件
│   └── config.yaml   # 主配置文件
├── internal/         # 内部包
│   ├── api/          # API 相关代码
│   │   ├── middlewares/  # 中间件
│   │   └── routes/       # 路由
│   ├── config/       # 配置管理
│   ├── database/     # 数据库连接
│   ├── errors/       # 错误处理
│   ├── models/       # 数据模型
│   └── utils/        # 工具函数
├── .gitignore        # Git 忽略文件
├── docker-compose.yml # Docker Compose 配置
├── Dockerfile        # Docker 构建文件
├── go.mod            # Go 模块文件
├── go.sum            # Go 依赖校验文件
├── main.go           # 主入口文件
└── README.md         # 项目文档
```

## 快速开始

### 1. 环境要求

- Go 1.23.0 或更高版本
- MySQL 8.0 或更高版本
- Docker（可选，用于容器化部署）

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置数据库

修改 `configs/config.yaml` 文件中的数据库配置：

```yaml
database:
  host: "localhost"
  port: "3306"
  user: "root"
  password: "root"
  dbname: "app"
  charset: "utf8mb4"
  parseTime: true
  loc: "Local"
```

### 4. 运行项目

#### 直接运行

```bash
go run main.go
```

#### 构建并运行

```bash
go build -o ginx main.go
./ginx
```

#### 使用 Docker 运行

```bash
docker-compose up -d
```

### 5. 访问 API

项目启动后，可通过以下地址访问 API：

- 基础地址：`http://localhost:8080`
- 健康检查：`http://localhost:8080/health`

## 开发指南

### 1. 添加新路由

在 `internal/api/routes/routes.go` 文件中添加新的路由：

```go
func SetupRoutes(router *gin.Engine) {
    // API 分组
    api := router.Group("/api")
    
    // 健康检查
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })
    
    // 添加新路由
    api.GET("/users", GetUsers)
    api.POST("/users", CreateUser)
    api.GET("/users/:id", GetUser)
    api.PUT("/users/:id", UpdateUser)
    api.DELETE("/users/:id", DeleteUser)
}
```

### 2. 添加新中间件

在 `internal/api/middlewares/` 目录下创建新的中间件文件，然后在 `internal/api/middlewares/middlewares.go` 文件中注册：

```go
func SetupMiddlewares(router *gin.Engine) {
    router.Use(Logger())
    router.Use(RequestID())
    router.Use(Cors())
    router.Use(Session())
    router.Use(HTTPSRedirect())
    router.Use(Compress())
    router.Use(RateLimit())
    router.Use(ErrorHandler())
}
```

### 3. 添加新模型

在 `internal/models/` 目录下创建新的模型文件，然后在 `main.go` 文件中添加自动迁移：

```go
// 自动迁移数据库表结构
err = database.AutoMigrate(&models.User{}, &models.NewModel{})
if err != nil {
    log.Fatalf("迁移数据库表结构失败: %v", err)
}
```

### 4. 配置管理

项目使用 Viper 管理配置，配置文件位于 `configs/config.yaml`，可通过环境变量覆盖配置值。

## 部署指南

### 1. 本地部署

```bash
# 构建项目
go build -o ginx main.go

# 运行项目
./ginx
```

### 2. Docker 部署

```bash
# 构建镜像
docker build -t ginx .

# 运行容器
docker run -d --name ginx -p 8080:8080 ginx
```

### 3. Docker Compose 部署

```bash
docker-compose up -d
```

## 测试

### 运行单元测试

```bash
go test ./...
```

### 运行基准测试

```bash
go test -bench=. ./...
```

### 性能分析

```bash
# CPU 分析
go test -cpuprofile=cpu.prof ./...
go tool pprof cpu.prof

# 内存分析
go test -memprofile=mem.prof ./...
go tool pprof mem.prof
```

## 监控

项目集成了以下监控功能：

- 请求日志
- 错误日志
- 性能分析

## 常见问题

### 1. 数据库连接失败

- 检查数据库服务是否运行
- 检查数据库配置是否正确
- 检查网络连接是否正常

### 2. 端口被占用

- 使用 `lsof -i :8080` 查看端口占用情况
- 终止占用端口的进程或修改项目端口配置

### 3. 依赖包下载失败

- 检查网络连接
- 设置 Go 模块代理：
  ```bash
  go env -w GOPROXY=https://goproxy.io,direct
  ```

## 技术栈

| 技术/工具 | 版本 | 用途 |
|-----------|------|------|
| Go        | 1.23.0 | 开发语言 |
| Gin       | v1.11.0 | Web 框架 |
| GORM      | v1.25.7 | ORM 框架 |
| Viper     | v1.18.2 | 配置管理 |
| JWT       | v5.2.0 | 认证 |
| MySQL     | 8.0+ | 数据库 |
| Docker    | - | 容器化 |

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

[MIT](LICENSE)
