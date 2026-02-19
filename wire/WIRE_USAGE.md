# Wire 工具使用说明

## 什么是 Wire？

Wire 是 Google 开发的一个依赖注入工具，用于在 Go 语言中自动生成依赖注入代码，避免手动编写繁琐的依赖管理代码。

## 安装 Wire

当网络连接恢复后，运行以下命令安装 Wire 工具：

```bash
go install github.com/google/wire/cmd/wire@latest
```

## 生成依赖注入代码

在项目根目录运行以下命令生成依赖注入代码：

```bash
wire
```

这将生成 `wire_gen.go` 文件，包含自动生成的依赖注入代码。

## 项目结构

### 服务层

服务层位于 `internal/services/` 目录，包含各个业务模块的服务接口和实现：

- `user_service.go` - 用户服务

### Wire 配置

- `wire.go` - Wire 配置文件，定义依赖注入关系
- `wire_gen.go` - 自动生成的依赖注入代码（运行 wire 命令后生成）

### 处理器

处理器位于 `internal/api/handlers/` 目录，使用依赖注入接收服务实例：

- `user.go` - 用户处理器，依赖 UserService

### 路由

路由位于 `internal/api/routes/` 目录，接收处理器实例并注册路由：

- `routes.go` - 主路由文件，接收 UserHandler 实例

## 依赖注入流程

1. **定义服务接口和实现** - 在 `services` 目录中定义服务
2. **创建处理器** - 在 `handlers` 目录中创建处理器，依赖服务接口
3. **配置 Wire** - 在 `wire.go` 中定义依赖注入关系
4. **生成代码** - 运行 `wire` 命令生成依赖注入代码
5. **使用注入** - 在 `main.go` 中使用生成的函数初始化处理器

## 示例

### 服务定义

```go
// UserService 用户服务接口
type UserService interface {
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

// NewUserService 创建用户服务实例
func NewUserService() UserService {
	return &userService{}
}
```

### 处理器定义

```go
// UserHandler 用户处理器
type UserHandler struct {
	userService services.UserService
}

// NewUserHandler 创建用户处理器实例
func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
```

### Wire 配置

```go
// WireSet 依赖注入集合
var WireSet = wire.NewSet(
	services.NewUserService,
	handlers.NewUserHandler,
)

// InitializeHandlers 初始化所有处理器
func InitializeHandlers() (*handlers.UserHandler, error) {
	wire.Build(WireSet)
	return nil, nil
}
```

### 使用注入

```go
// 初始化处理器
userHandler, err := InitializeHandlers()
if err != nil {
	log.Fatalf("初始化处理器失败: %v", err)
}

// 设置路由
routes.SetupRoutes(router, userHandler)
```

## 扩展

要添加新的服务和处理器，只需：

1. 在 `services` 目录中创建新的服务
2. 在 `handlers` 目录中创建新的处理器，依赖相应的服务
3. 在 `wire.go` 文件的 `WireSet` 中添加新的依赖
4. 运行 `wire` 命令重新生成代码
5. 在 `main.go` 中使用新的初始化函数
