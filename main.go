package main

import (
	"fmt"
	"time"

	"test/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	// 初始化日志
	if err := logger.Init(nil); err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logger.Sync()

	logger.Info("Application started",
		zap.String("version", "1.0.0"),
		zap.Time("start_time", time.Now()),
	)

	fmt.Println("hello, vscode")
	logger.Info("Hello message printed")

	// 测试日志功能
	testLogging()

	// 测试递归函数
	result := fact(5)
	logger.Info("Factorial calculation completed",
		zap.Int("input", 5),
		zap.Int("result", result),
	)
	fmt.Printf("fact(5) = %d\n", result)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// 测试日志功能
func testLogging() {
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")

	// 模拟一些业务逻辑
	for i := 1; i <= 3; i++ {
		logger.Info("Processing item", zap.Int("item_id", i))
		time.Sleep(100 * time.Millisecond) // 模拟处理时间
	}
}
