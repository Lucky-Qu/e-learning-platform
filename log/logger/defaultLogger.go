package logger

import "go.uber.org/zap"

var DefaultLogger struct {
	Logger *zap.Logger
}

func initDefaultLogger(zapLogger *zap.Logger) {
	DefaultLogger.Logger = zapLogger
}
