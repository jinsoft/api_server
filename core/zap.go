package core

import (
	"api_server/global"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"strings"
	"time"
)

var level zapcore.Level

func Zap() (logger *zap.Logger) {
	level = getLevel(global.CONFIG.Zap.Level)

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	return logger
}

func getEncoderCore() (core zapcore.Core) {
	write, err := rotatelogs.New(
		path.Join(global.CONFIG.Zap.Director, "%Y-%m-%d.log"),
		//rotatelogs.WithLinkName(global.CONFIG.Zap.LinkName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), zapcore.AddSync(write), level)
}

func getEncoder() zapcore.Encoder {
	if global.CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		//EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: CustomTimeEncoder,
		//EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	switch global.CONFIG.Zap.EncodeLevel {
	case "LowercaseLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseLevelEncoder // 小写编码器(默认)
	case "LowercaseColorLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder // 小写编码器带颜色
	case "CapitalLevelEncoder":
		config.EncodeLevel = zapcore.CapitalLevelEncoder // 大写编码器
	case "CapitalColorLevelEncoder":
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder // 大写编码器带颜色
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

func getLevel(level string) (lv zapcore.Level) {
	level = strings.ToLower(level)

	switch level {
	case "debug":
		lv = zap.DebugLevel
	case "info":
		lv = zap.InfoLevel
	case "warn":
		lv = zap.WarnLevel
	case "error":
		lv = zap.ErrorLevel
	case "dpanic":
		lv = zap.DPanicLevel
	case "panic":
		lv = zap.PanicLevel
	case "fatal":
		lv = zap.FatalLevel
	default:
		lv = zap.InfoLevel
	}
	return
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05"))
}
