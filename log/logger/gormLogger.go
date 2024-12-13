package logger

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"time"
)

var GormLogger gormLogger

type gormLogger struct {
	Logger   *zap.Logger
	LogLevel logger.LogLevel
}

func (g gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := g
	newLogger.LogLevel = level
	return &newLogger
}

func (g gormLogger) Info(ctx context.Context, msg string, i ...interface{}) {
	if g.LogLevel >= logger.Info {
		g.Logger.Info(msg, zap.Any("Data", i))
	}
}

func (g gormLogger) Warn(ctx context.Context, msg string, i ...interface{}) {
	if g.LogLevel >= logger.Warn {
		g.Logger.Warn(msg, zap.Any("Data", i))
	}
}

func (g gormLogger) Error(ctx context.Context, msg string, i ...interface{}) {
	if g.LogLevel >= logger.Error {
		g.Logger.Error(msg, zap.Any("Data", i))
	}
}

func (g gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.LogLevel < logger.Silent {
		return
	}
	// 计算执行时间
	elapsed := time.Since(begin)
	sql, rows := fc()
	// 如果有错误，记录错误日志
	if err != nil {
		g.Logger.Error("查询出错",
			zap.Error(err),
			zap.String("sql指令：", sql),
			zap.Duration("查询时间：", elapsed),
			zap.Int64("查询影响行：", rows),
		)
	} else if elapsed > 200*time.Millisecond { //  200ms 作为慢查询的阈值
		g.Logger.Warn("慢查询",
			zap.String("sql指令：", sql),
			zap.Duration("查询时间：", elapsed),
			zap.Int64("查询影响行：", rows),
		)
	} else {
		g.Logger.Info("查询成功",
			zap.String("sql指令：", sql),
			zap.Duration("查询时间：", elapsed),
			zap.Int64("查询影响行：", rows),
		)
	}
}

func initGormLogger(basicLogger *zap.Logger) {
	GormLogger.Logger = basicLogger
}
