package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type logFunc struct {
	Debug  func(v ...interface{})
	Debugf func(format string, v ...interface{})
	Info   func(v ...interface{})
	Infof  func(format string, v ...interface{})
	Warn   func(v ...interface{})
	Warnf  func(format string, v ...interface{})
	Error  func(v ...interface{})
	Errorf func(format string, v ...interface{})
	Fatal  func(v ...interface{})
	Fatalf func(format string, v ...interface{})
}

type Logger struct {
	debugLog *log.Logger
	infoLog  *log.Logger
	warnLog  *log.Logger
	errorLog *log.Logger
	fatalLog *log.Logger
	loggers  map[string]*log.Logger
	output   io.Writer
	mu       sync.Mutex
	*logFunc
}

func New() *Logger {
	output := io.Writer(os.Stdout)
	debugLog := log.New(output, prefixes["DEBUG"], log.LstdFlags|log.Lshortfile)
	infoLog := log.New(output, prefixes["INFO"], log.LstdFlags|log.Lshortfile)
	warnLog := log.New(output, prefixes["WARN"], log.LstdFlags|log.Lshortfile)
	errorLog := log.New(output, prefixes["ERROR"], log.LstdFlags|log.Lshortfile)
	fatalLog := log.New(output, prefixes["FATAL"], log.LstdFlags|log.Lshortfile)
	loggers := map[string]*log.Logger{"DEBUG": debugLog, "INFO": infoLog, "WARN": warnLog, "ERROR": errorLog, "FATAL": fatalLog}
	return &Logger{debugLog: debugLog, infoLog: infoLog, warnLog: warnLog, errorLog: errorLog, fatalLog: fatalLog, output: output, loggers: loggers,
		logFunc: &logFunc{Debug: debugLog.Println, Debugf: debugLog.Printf, Info: infoLog.Println, Infof: infoLog.Printf, Warn: warnLog.Println, Warnf: warnLog.Printf, Error: errorLog.Println, Errorf: errorLog.Printf, Fatal: fatalLog.Fatalln, Fatalf: fatalLog.Fatalf}}
}

// Set controls log level and output
func (l *Logger) Set(level int, out io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if out == nil {
		out = l.output
	} else {
		l.output = out
	}

	for k, logger := range l.loggers {
		logger.SetOutput(out)
		if out == os.Stdout {
			logger.SetPrefix(prefixes[k])
		} else {
			logger.SetPrefix(fmt.Sprintf("[%s]", k))
		}
	}

	if DebugLevel < level {
		l.debugLog.SetOutput(io.Discard)
	}
	if InfoLevel < level {
		l.infoLog.SetOutput(io.Discard)
	}
	if WarnLevel < level {
		l.warnLog.SetOutput(io.Discard)
	}
	if ErrorLevel < level {
		l.errorLog.SetOutput(io.Discard)
	}
	if FatalLevel < level {
		l.fatalLog.SetOutput(io.Discard)
	}
}
