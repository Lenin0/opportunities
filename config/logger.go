package config

import (
	"io"
	"log"
	"os"
)

// Log level
type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

// Interface to allow flexible use or future replacement
type LoggerInterface interface {
	Debug(args ...interface{})
	Debugf(format string, v ...interface{})
	Info(args ...interface{})
	Infof(format string, v ...interface{})
	Warn(args ...interface{})
	Warnf(format string, v ...interface{})
	Error(args ...interface{})
	Errorf(format string, v ...interface{})
}

// Logger structure
type Logger struct {
	level   LogLevel
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

// Function to instantiate the logge
func NewLogger(prefix string, level LogLevel, outputPaths ...string) *Logger {
	var writers []io.Writer

	// Always writes to stdout
	writers = append(writers, os.Stdout)

	// If files paths are passed, write to them as well

	for _, path := range outputPaths {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Printf("error opening log file %s: %v", path, err)
			continue
		}
		writers = append(writers, file)
	}

	writer := io.MultiWriter(writers...)
	flags := log.Ldate | log.Ltime | log.Lshortfile

	return &Logger{
		level:   level,
		debug:   log.New(writer, prefix+" DEBUG: ", flags),
		info:    log.New(writer, prefix+" INFO: ", flags),
		warning: log.New(writer, prefix+" WARNING: ", flags),
		err:     log.New(writer, prefix+" ERROR ", flags),
		writer:  writer,
	}
}

// Functions with level checking

func (l *Logger) Debug(args ...interface{}) {
	if l.level <= DebugLevel {
		l.debug.Println(args...)
	}
}
func (l *Logger) Info(args ...interface{}) {
	if l.level <= InfoLevel {
		l.info.Println(args...)
	}
}
func (l *Logger) Warning(args ...interface{}) {
	if l.level <= WarnLevel {
		l.warning.Println(args...)
	}
}
func (l *Logger) Error(args ...interface{}) {
	if l.level <= ErrorLevel {
		l.err.Println(args...)
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level <= DebugLevel {
		l.debug.Printf(format, v...)
	}
}
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level <= InfoLevel {
		l.info.Printf(format, v...)
	}
}
func (l *Logger) Warningf(format string, v ...interface{}) {
	if l.level <= WarnLevel {
		l.warning.Printf(format, v...)
	}
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level <= ErrorLevel {
		l.err.Printf(format, v...)
	}
}
