package logger

type Logger interface {
	Info(args ...interface{})
	Infow(template string, args ...interface{})
	Infof(msg string, args ...interface{})

	Warn(args ...interface{})
	Warnw(template string, args ...interface{})
	Warnf(template string, args ...interface{})

	Error(args ...interface{})
	Errorw(template string, args ...interface{})
	Errorf(msg string, args ...interface{})

	Debug(args ...interface{})
	Debugw(template string, args ...interface{})
	Debugf(msg string, args ...interface{})

	Fatal(args ...interface{})
	Fatalw(template string, args ...interface{})
	Fatalf(msg string, args ...interface{})
}

var impl Logger
var loggerSpace = "dev"

func SetLogger(isDebug bool) {
	if isDebug {
		loggerSpace = "prod"
	}
	impl = NewZapLogger(loggerSpace)
}

func Info(args ...interface{}) {
	impl.Info(args...)
}

func Infow(template string, args ...interface{}) {
	impl.Infow(template, args...)
}

func Infof(msg string, args ...interface{}) {
	impl.Infof(msg, args...)
}

func Warn(args ...interface{}) {
	impl.Warn(args...)
}

func Warnw(template string, args ...interface{}) {
	impl.Warnw(template, args...)
}

func Warnf(template string, args ...interface{}) {
	impl.Warnf(template, args...)
}

func Error(args ...interface{}) {
	impl.Error(args...)
}

func Errorw(template string, args ...interface{}) {
	impl.Errorw(template, args...)
}

func Errorf(msg string, args ...interface{}) {
	impl.Errorf(msg, args...)
}

func Debug(args ...interface{}) {
	impl.Debug(args...)
}

func Debugw(template string, args ...interface{}) {
	impl.Debugw(template, args...)
}

func Debugf(msg string, args ...interface{}) {
	impl.Debugf(msg, args...)
}

func Fatal(args ...interface{}) {
	impl.Fatal(args...)
}

func Fatalw(template string, args ...interface{}) {
	impl.Fatalw(template, args...)
}

func Fatalf(msg string, args ...interface{}) {
	impl.Fatalf(msg, args...)
}
