package libs

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

////////////////// Logger Functions //////////////////

// CreateLogger creates a new Zap logger to be used for RF API Message Logging

func CreateLogger() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
			logfile,
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	return zap.Must(config.Build())
}

func logmaker() *zap.Logger {
	logger := CreateLogger()
	defer logger.Sync()
	return logger
}
