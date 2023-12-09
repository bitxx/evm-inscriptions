package log

import (
	"github.com/bitxx/logger"
	"github.com/bitxx/logger/logbase"
)

var logHelper *logbase.Helper

type LoggerConf struct {
	Type      string
	Path      string
	Level     string
	Stdout    string
	EnabledDB bool
	Cap       uint
}

func Init(loggerConf LoggerConf) {
	/*logHelper = logger.NewLogger(
		logger.WithType(loggerConf.Type),
		logger.WithPath(loggerConf.Path),
		logger.WithLevel(loggerConf.Level),
		logger.WithStdout(loggerConf.Stdout),
		logger.WithCap(loggerConf.Cap),
	)*/
	logHelper = logger.NewLogger(
		logger.WithType("default"),
		logger.WithPath(""),
		logger.WithLevel("debug"),
		logger.WithStdout("default"),
		logger.WithCap(100),
	)
}

func Info(args ...interface{}) {
	logHelper.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logHelper.Infof(template, args...)
}

func Trace(args ...interface{}) {
	logHelper.Trace(args...)
}

func Tracef(template string, args ...interface{}) {
	logHelper.Tracef(template, args...)
}

func Debug(args ...interface{}) {
	logHelper.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logHelper.Debugf(template, args...)
}

func Warn(args ...interface{}) {
	logHelper.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logHelper.Warnf(template, args...)
}

func Error(args ...interface{}) {
	logHelper.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logHelper.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	logHelper.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logHelper.Fatalf(template, args...)
}

func WithError(err error) *logbase.Helper {
	return logHelper.WithError(err)
}

func WithFields(fields map[string]interface{}) *logbase.Helper {
	return logHelper.WithFields(fields)
}
