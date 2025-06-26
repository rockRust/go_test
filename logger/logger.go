package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var globalLogger *zap.Logger

// Config 日志配置结构
type Config struct {
	LogDir     string // 日志目录
	LogFile    string // 日志文件名
	MaxSize    int    // 单个日志文件最大尺寸(MB)
	MaxBackups int    // 保留的日志文件数量
	MaxAge     int    // 保留天数
	Compress   bool   // 是否压缩
	LogLevel   string // 日志级别
}

// DefaultConfig 默认日志配置
func DefaultConfig() *Config {
	return &Config{
		LogDir:     "logs",
		LogFile:    "myapp.log",
		MaxSize:    100,
		MaxBackups: 7,
		MaxAge:     7,
		Compress:   true,
		LogLevel:   "info",
	}
}

// 自定义时间编码器，精确到毫秒
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// Init 初始化全局日志
func Init(config *Config) error {
	if config == nil {
		config = DefaultConfig()
	}

	// 创建日志目录
	if err := os.MkdirAll(config.LogDir, 0755); err != nil {
		return err
	}

	// 配置lumberjack进行日志切割
	writer := &lumberjack.Logger{
		Filename:   filepath.Join(config.LogDir, config.LogFile),
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
		LocalTime:  true,
	}

	// 配置zap编码器 - 使用控制台格式而不是JSON
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = customTimeEncoder // 使用自定义时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// 解析日志级别
	level := parseLogLevel(config.LogLevel)

	// 创建核心配置 - 使用控制台编码器
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // 改为控制台格式
		zapcore.AddSync(writer),
		level,
	)

	// 创建logger
	globalLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return nil
}

// GetLogger 获取全局日志实例
func GetLogger() *zap.Logger {
	if globalLogger == nil {
		// 如果还没有初始化，使用默认配置初始化
		if err := Init(nil); err != nil {
			log.Fatal("Failed to initialize logger:", err)
		}
	}
	return globalLogger
}

// Sync 同步日志缓冲区
func Sync() error {
	if globalLogger != nil {
		return globalLogger.Sync()
	}
	return nil
}

// parseLogLevel 解析日志级别
func parseLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// 便捷方法，直接调用全局logger
func Info(msg string, fields ...zap.Field) {
	GetLogger().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	GetLogger().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	GetLogger().Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	GetLogger().Debug(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	GetLogger().Fatal(msg, fields...)
}
