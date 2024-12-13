package logger

import (
	AppConfig "e-learning-platform/config"
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	basicLogger *zap.Logger
	err         error
)

func InitLogger() {
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:          "消息：",
			LevelKey:            "等级：",
			TimeKey:             "时间：",
			NameKey:             "名字：",
			CallerKey:           "调用者：",
			FunctionKey:         "调用函数：",
			StacktraceKey:       "堆栈键名：",
			SkipLineEnding:      false,
			LineEnding:          zapcore.DefaultLineEnding,
			EncodeLevel:         zapcore.LowercaseLevelEncoder,
			EncodeTime:          zapcore.ISO8601TimeEncoder,
			EncodeDuration:      zapcore.SecondsDurationEncoder,
			EncodeCaller:        zapcore.FullCallerEncoder,
			EncodeName:          zapcore.FullNameEncoder,
			NewReflectedEncoder: nil,
			ConsoleSeparator:    "\t",
		},
		OutputPaths:      []string{AppConfig.Config.Log.Path},
		ErrorOutputPaths: []string{AppConfig.Config.Log.Path},
		InitialFields:    nil,
	}
	basicLogger, err = config.Build()
	if err != nil {
		panic(errors.New("日志初始化失败"))
	}
	initDefaultLogger(basicLogger)
	initGormLogger(basicLogger)
}
