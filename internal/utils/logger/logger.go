package logger

import (
	"fmt"
	"io"
	"os"
	"time"
	"unicode/utf8"
)

type Logger struct {
	prefix string
	writer io.Writer
}
type level string

const (
	_LOG     level = "LOG"
	_ERROR   level = "ERROR"
	_WARNING level = "WARN"
)

func (l Logger) log(level level, raw string, values ...any) {
	if !Enabled {
		return
	}

	fmt.Fprintf(l.writer, "[%s][%s][%s] ", time.Now().Format(time.DateTime), level, l.prefix)
	fmt.Fprintf(l.writer, raw, values...)
	var buf [4]byte
	n := utf8.EncodeRune(buf[:], '\n')
	l.writer.Write(buf[:n])
}

func (l Logger) Log(raw string, values ...any) {
	l.log(_LOG, raw, values...)
}

func (l Logger) Warn(raw string, values ...any) {
	l.log(_WARNING, raw, values...)
}

func (l Logger) Error(raw string, values ...any) {
	l.log(_ERROR, raw, values...)
}

var loggers = map[string]*Logger{}

func Get(prefix string) *Logger {
	logger, ok := loggers[prefix]
	if !ok {
		loggers[prefix] = NetW(prefix, os.Stdout)
		return loggers[prefix]
	}
	return logger
}

func NetW(prefix string, w io.Writer) *Logger {
	logger, ok := loggers[prefix]
	if !ok {
		loggers[prefix] = &Logger{
			prefix: prefix,
			writer: w,
		}
		return loggers[prefix]
	}
	return logger
}

var Enabled bool = false
