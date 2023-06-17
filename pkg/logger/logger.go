package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder

	config := zap.Config{
		Encoding:          "json",
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		DisableCaller:     true,
		DisableStacktrace: true,
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync()
	return logger
}
