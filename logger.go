package paulog

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	logLevels = map[string]LogLevel{}
)

func SetLevel(pkgPrefix string, level LogLevel) {
	logLevels[pkgPrefix] = level
}

func ClearLogLevels() {
	logLevels = map[string]LogLevel{}
}

type Logger struct {
	pkgName       string
	errorLogger   *log.Logger
	warningLogger *log.Logger
	infoLogger    *log.Logger
	debugLogger   *log.Logger
}

const (
	green      = "\x1b[32;20m"
	grey       = "\x1b[36;20m"
	yellow     = "\x1b[33;20m"
	red        = "\x1b[31;20m"
	resetColor = "\x1b[0m"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func newLogger(pkgName string, level LogLevel) *log.Logger {
	levelColor := resetColor
	switch level {
	case DEBUG:
		levelColor = grey
	case INFO:
		levelColor = green
	case WARNING:
		levelColor = yellow
	case ERROR:
		levelColor = red
	}
	return log.New(os.Stderr, fmt.Sprintf("\b\b --- [%s %s %s] %s: ", levelColor, level, resetColor, pkgName), log.Ldate|log.Ltime|log.Lmsgprefix|log.Lshortfile)
}

func GetLogger(pkgName string) *Logger {
	if len(pkgName) == 0 {
		panic("pkgName is empty")
	}

	return &Logger{
		pkgName:       pkgName,
		errorLogger:   newLogger(pkgName, ERROR),
		warningLogger: newLogger(pkgName, WARNING),
		infoLogger:    newLogger(pkgName, INFO),
		debugLogger:   newLogger(pkgName, DEBUG),
	}
}

func (l *Logger) canLog(level LogLevel) bool {
	subParts := strings.Split(l.pkgName, ".")
	lastLevel := INFO
	for i, _ := range subParts {
		pkgPrefix := strings.Join(subParts[:i+1], ".")
		pkgLevel, ok := logLevels[pkgPrefix]
		if !ok {
			continue
		}
		lastLevel = pkgLevel
	}

	return lastLevel <= level
}

func (l *Logger) Debugf(format string, args ...any) {
	if !l.canLog(DEBUG) {
		return
	}
	l.debugLogger.Output(2, fmt.Sprintf(format, args...))
}

func (l *Logger) Infof(format string, args ...any) {
	if !l.canLog(INFO) {
		return
	}
	l.infoLogger.Output(2, fmt.Sprintf(format, args...))
}

func (l *Logger) Warnf(format string, args ...any) {
	if !l.canLog(WARNING) {
		return
	}
	l.warningLogger.Output(2, fmt.Sprintf(format, args...))
}

func (l *Logger) Warningf(format string, args ...any) {
	l.Warnf(format, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	if !l.canLog(ERROR) {
		return
	}
	l.errorLogger.Output(2, fmt.Sprintf(format, args...))
}
