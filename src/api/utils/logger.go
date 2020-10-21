package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type ILogger interface {
	Init()
	SetLogLevel(loglevel string)
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, err error, args ...interface{})
	Panic(format string, err error, args ...interface{})
}

type Logger struct{
	logLevel string
	Log *logrus.Logger
}

func (l *Logger) Init(){
	l.Log = &logrus.Logger{
		Out:       os.Stdout,
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
		Formatter: &logrus.TextFormatter {
			DisableColors: true,
			FullTimestamp: true,
		},
	}
}
func (l *Logger) SetLogLevel(logLevel string) {
	if level, error := logrus.ParseLevel(logLevel); error != nil {
		panic(error)
	} else {
		l.Log.Level = level
	}
}
func (l *Logger) Debug(format string, args ...interface{}){
	if l.Log.Level >= logrus.DebugLevel {
		l.Log.Debug(fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Info(format string, args ...interface{}){
	if l.Log.Level >= logrus.InfoLevel {
		l.Log.Info(fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Warn(format string, args ...interface{}){
	if l.Log.Level >= logrus.WarnLevel {
		l.Log.Warn(fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Error(format string, err error, args ...interface{}){
	if l.Log.Level >= logrus.ErrorLevel {
		l.Log.Error(fmt.Sprintf(format, args...), err)
	}
}

func (l *Logger) Panic(format string, err error, args ...interface{}){
	if l.Log.Level >= logrus.PanicLevel {
		l.Log.Panic(fmt.Sprintf(format, args...), err)
	}
}