package logger

import (
	"bytes"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Printer func(string, ...zap.Field)

func (p Printer) Write(b []byte) (int, error) {
	if p != nil {
		p(string(bytes.TrimSpace(b)))
	}
	return len(b), nil
}

func NewLogger() *zap.SugaredLogger {

	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	enc := zapcore.NewConsoleEncoder(cfg)

	enabler := zap.NewAtomicLevelAt(zap.DebugLevel)
	logger := zap.New(zapcore.NewCore(enc, zapcore.AddSync(os.Stdout), enabler))
	zap.ReplaceGlobals(logger)

	log.SetFlags(0)
	log.SetOutput(Printer(logger.Debug))

	return logger.Sugar()
}
