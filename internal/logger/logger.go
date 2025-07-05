package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log      *zap.SugaredLogger
	initOnce sync.Once
)

func initLogger(f *os.File) *zap.SugaredLogger {
	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(pe)
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(f), zap.DebugLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.DebugLevel),
	)

	logger := zap.New(core)
	return logger.Sugar()
}

func Init(f *os.File) {
	initOnce.Do(func() {
		log = initLogger(f)
	})
}

func Get() *zap.SugaredLogger {
	if log == nil {
		panic("logger is not initialized. Call logger.Init() before using logger.Get()")
	}
	return log
}
