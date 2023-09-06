package logger

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

type LogLevel int

const (
	Debug   LogLevel = 0
	Info    LogLevel = 1
	Warning LogLevel = 2
	Error   LogLevel = 3
)

type logger struct {
	log   *logrus.Entry
	level LogLevel
}

func New(service string, level LogLevel) Logger {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]

	log := logrus.WithField("file", filename).WithField("function", fn).WithField("service", service)
	return &logger{log, level}
}

func (l *logger) Debug(msg string) {
	l.log.Debugln(msg)
}

func (l *logger) Info(msg string) {
	l.log.Infoln(msg)
}

func (l *logger) Warning(msg string) {
	l.log.Warningln(msg)
}

func (l *logger) Error(msg string) {
	l.log.Errorln(msg)
}
