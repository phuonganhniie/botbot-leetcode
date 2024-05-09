package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

func InitLogger() *zap.SugaredLogger {
	cfg := zap.Config{
		Encoding:    "console",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			TimeKey:      "time",
			LevelKey:     "level",
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeLevel:  customLevelEncoder,
			EncodeTime:   syslogTimeEncoder,
		},
	}

	coreLogger, _ := cfg.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	defer coreLogger.Sync()

	log = coreLogger.Sugar()
	return log
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("02-01-2006 15:04:05"))
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func Info(message string) {
	log.Info(message)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Error(message string) {
	log.Error(message)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Warn(message string) {
	log.Warn(message)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}
