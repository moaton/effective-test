package logger

import (
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	log *zap.SugaredLogger
}

func NewZapLogger(space string) *ZapLogger {
	var config = zap.Config{}

	config = zap.NewDevelopmentConfig()
	config.DisableStacktrace = true
	config.EncoderConfig.ConsoleSeparator = "  |  "
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(filepath.Base(caller.FullPath()))
	}
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000")

	logger, err := config.Build(zap.AddCaller(), zap.AddCallerSkip(2))
	if err != nil {
		panic(err)
	}

	return &ZapLogger{logger.Sugar()}
}

func (z *ZapLogger) Info(args ...interface{}) {
	z.log.Info(args...)
}

func (z *ZapLogger) Infow(template string, args ...interface{}) {
	z.log.Infow(template, args...)
}

func (z *ZapLogger) Infof(msg string, args ...interface{}) {
	z.log.Infof(msg, args...)
}

func (z *ZapLogger) Warn(args ...interface{}) {
	z.log.Warn(args...)
}

func (z *ZapLogger) Warnw(template string, args ...interface{}) {
	z.log.Warnw(template, args...)
}

func (z *ZapLogger) Warnf(template string, args ...interface{}) {
	z.log.Warnf(template, args...)
}

func (z *ZapLogger) Error(args ...interface{}) {
	z.log.Error(args...)
}

func (z *ZapLogger) Errorw(template string, args ...interface{}) {
	z.log.Errorw(template, args...)
}

func (z *ZapLogger) Errorf(msg string, args ...interface{}) {
	z.log.Errorf(msg, args...)
}

func (z *ZapLogger) Debug(args ...interface{}) {
	z.log.Debug(args...)
}

func (z *ZapLogger) Debugw(template string, args ...interface{}) {
	z.log.Debugw(template, args...)
}

func (z *ZapLogger) Debugf(msg string, args ...interface{}) {
	z.log.Debugf(msg, args...)
}

func (z *ZapLogger) Fatal(args ...interface{}) {
	z.log.Fatal(args...)
}

func (z *ZapLogger) Fatalw(template string, args ...interface{}) {
	z.log.Fatalw(template, args...)
}

func (z *ZapLogger) Fatalf(msg string, args ...interface{}) {
	z.log.Fatalf(msg, args...)
}
