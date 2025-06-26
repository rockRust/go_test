# Go 项目日志系统

这个项目集成了按天切割的日志系统，支持自动清理旧日志文件。日志功能已封装到独立的包中，便于使用和维护。

## 功能特性

- ✅ 按天自动切割日志文件
- ✅ 自动保留最近7天的日志
- ✅ 自动压缩旧日志文件
- ✅ 单行文本格式日志（易读易解析）
- ✅ 支持不同日志级别
- ✅ 可配置的日志参数
- ✅ 独立的日志包，便于复用

## 项目结构

```
test/
├── main.go                 # 主程序（已简化）
├── go.mod                  # 依赖管理
├── Dockerfile              # Docker 配置
├── .dockerignore           # Docker 忽略文件
├── README.md               # 使用说明
├── pkg/
│   └── logger/             # 日志包
│       ├── logger.go       # 核心日志功能
│       └── example.go      # 使用示例
└── logs/                   # 日志目录（自动创建）
    └── app.log             # 当前日志文件
```

## 日志配置

### 默认配置
- 日志目录：`logs/`
- 日志文件：`app.log`
- 单个文件最大：100MB
- 保留文件数：7个
- 保留天数：7天
- 日志级别：INFO
- 自动压缩：是

### 自定义配置

```go
import "test/pkg/logger"

// 自定义配置
customConfig := &logger.Config{
    LogDir:     "logs",
    LogFile:    "myapp.log",
    MaxSize:    50,    // 50MB
    MaxBackups: 7,     // 保留7个文件
    MaxAge:     7,     // 保留7天
    Compress:   true,  // 压缩旧文件
    LogLevel:   "debug", // 调试级别
}

err := logger.Init(customConfig)
```

## 使用方法

### 基本使用（推荐）
```go
package main

import (
    "test/pkg/logger"
    "go.uber.org/zap"
)

func main() {
    // 初始化日志（使用默认配置）
    if err := logger.Init(nil); err != nil {
        panic("Failed to initialize logger: " + err.Error())
    }
    defer logger.Sync()

    // 使用便捷方法记录日志
    logger.Info("应用启动成功")
    logger.Warn("警告信息")
    logger.Error("错误信息")
    
    // 带字段的结构化日志
    logger.Info("用户登录",
        zap.String("username", "john"),
        zap.String("ip", "192.168.1.1"),
    )
}
```

### 获取Logger实例
```go
// 获取logger实例进行更复杂的操作
log := logger.GetLogger()
log.Info("使用logger实例记录日志")
```

### 日志级别
- `debug`: 调试信息
- `info`: 一般信息
- `warn`: 警告信息
- `error`: 错误信息
- `fatal`: 致命错误

## 日志格式

日志采用单行文本格式，便于阅读和解析：

```
2025-06-26 14:05:23.064 INFO    logger/logger.go:128    Processing item {"item_id": 1}
2025-06-26 14:05:23.064 INFO    logger/logger.go:128    Application started {"version": "1.0.0", "start_time": "2025-06-26 14:05:23.064"}
2025-06-26 14:05:23.064 WARN    logger/logger.go:128    This is a warning message
2025-06-26 14:05:23.064 ERROR   logger/logger.go:128    This is an error message
```

格式说明：
- `时间戳`: 精确到毫秒的时间格式 (YYYY-MM-DD HH:MM:SS.mmm)
- `日志级别`: INFO/WARN/ERROR等
- `调用位置`: 文件名:行号
- `消息内容`: 日志消息
- `字段`: 可选的JSON格式字段（如果有的话）

## 日志文件结构

运行后会在 `logs/` 目录下生成以下文件：
```
logs/
├── app.log          # 当前日志文件
├── app.log.1        # 昨天的日志
├── app.log.2.gz     # 前天的日志（压缩）
├── app.log.3.gz     # 3天前的日志（压缩）
└── ...
```

## 运行项目

```bash
# 下载依赖
go mod tidy

# 运行项目
go run main.go

# 查看日志
tail -f logs/app.log
```

## Docker 部署

```bash
# 构建镜像
docker build -t my-go-app .

# 运行容器
docker run -d -p 8080:8080 -v $(pwd)/logs:/app/logs my-go-app
```

## 高级用法

### 生产环境配置
```go
prodConfig := &logger.Config{
    LogDir:     "/var/log/myapp",
    LogFile:    "app.log",
    MaxSize:    200,   // 200MB
    MaxBackups: 10,    // 保留10个文件
    MaxAge:     30,    // 保留30天
    Compress:   true,
    LogLevel:   "info", // 信息级别
}

err := logger.Init(prodConfig)
```

### 调试环境配置
```go
debugConfig := &logger.Config{
    LogDir:     "logs",
    LogFile:    "debug.log",
    MaxSize:    50,    // 50MB
    MaxBackups: 3,     // 保留3个文件
    MaxAge:     3,     // 保留3天
    Compress:   false, // 不压缩，便于查看
    LogLevel:   "debug", // 调试级别
}

err := logger.Init(debugConfig)
```

## 注意事项

1. 确保 `logs/` 目录有写入权限
2. 日志文件会自动按大小和时间进行切割
3. 超过保留天数的日志文件会自动删除
4. 旧日志文件会自动压缩以节省空间
5. 日志格式为单行文本，便于使用 `grep`、`awk` 等工具处理
6. 时间戳精确到毫秒，便于性能分析和调试
7. 日志包是线程安全的，可以在多个goroutine中使用 