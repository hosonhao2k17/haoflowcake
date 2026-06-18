package logger

import (
	"haoflowcake/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zapLogger *zap.Logger
}

func NewLogger(cf *config.Config) *Logger {
	var logger *zap.Logger

	if cf.App.IsProduction() {
		logger, _ = zap.NewProduction()
	} else {
		developmentConfig := zap.NewDevelopmentConfig()
		developmentConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, _ = developmentConfig.Build()
	}
	logger = logger.With(
		zap.String("app", "haoflowcake"),
		zap.String("env", cf.App.AppEnv),
	)
	return &Logger{
		zapLogger: logger,
	}
}

func (z *Logger) Debug(msg string, fields ...zap.Field) { z.zapLogger.Debug(msg, fields...) }
func (z *Logger) Info(msg string, fields ...zap.Field)  { z.zapLogger.Info(msg, fields...) }
func (z *Logger) Warn(msg string, fields ...zap.Field)  { z.zapLogger.Warn(msg, fields...) }
func (z *Logger) Error(msg string, fields ...zap.Field) { z.zapLogger.Error(msg, fields...) }
func (z *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{
		zapLogger: z.zapLogger.With(fields...),
	}
}
func (z *Logger) Fatal(msg string, fields ...zap.Field) {
	z.zapLogger.Fatal(msg, fields...)
}
