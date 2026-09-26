package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func newLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

func (l *Logger) Debug [T any] (values ...T) {
	l.debug.Println(values)
}

func (l *Logger) Info [T any] (values ...T) {
	l.info.Println(values)
}

func (l *Logger) Warning [T any] (values ...T) {
	l.warning.Println(values)
}

func (l *Logger) Error [T any] (values ...T) {
	l.err.Println(values)
}


func (l *Logger) DebugF [T any] (format string, values ...T) {
	l.debug.Printf(format, values)
}

func (l *Logger) InfoF [T any] (format string, values ...T) {
	l.info.Printf(format, values)
}

func (l *Logger) WarningF [T any] (format string, values ...T) {
	l.warning.Printf(format, values)
}

func (l *Logger) ErrorF [T any] (format string, values ...T) {
	l.err.Printf(format, values)
}
