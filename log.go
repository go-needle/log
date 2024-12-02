package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

// log prefix on console
var prefixes = map[string]string{"DEBUG": "\033[31m[DEBUG]\033[0m ", "INFO": "\033[34m[INFO]\033[0m ", "WARN": "\033[33m[WARN]\033[0m ", "ERROR": "\033[31m[ERROR]\033[0m ", "FATAL": "\033[31m[FATAL]\033[0m "}

// log output levels
const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	Disabled
)

var (
	output   = io.Writer(os.Stdout)
	debugLog = log.New(output, prefixes["DEBUG"], log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(output, prefixes["INFO"], log.LstdFlags|log.Lshortfile)
	warnLog  = log.New(output, prefixes["WARN"], log.LstdFlags|log.Lshortfile)
	errorLog = log.New(output, prefixes["ERROR"], log.LstdFlags|log.Lshortfile)
	fatalLog = log.New(output, prefixes["FATAL"], log.LstdFlags|log.Lshortfile)
	loggers  = map[string]*log.Logger{"DEBUG": debugLog, "INFO": infoLog, "WARN": warnLog, "ERROR": errorLog, "FATAL": fatalLog}
	mu       sync.Mutex
)

// log methods
var (
	Debug  = debugLog.Println
	Debugf = debugLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
	Warn   = warnLog.Println
	Warnf  = warnLog.Printf
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Fatal  = fatalLog.Fatalln
	Fatalf = fatalLog.Fatalf
)

// Set controls log level and output
func Set(level int, out io.Writer) {
	mu.Lock()
	defer mu.Unlock()

	if out == nil {
		out = output
	} else {
		output = out
	}

	for k, logger := range loggers {
		logger.SetOutput(out)
		if out == os.Stdout {
			logger.SetPrefix(prefixes[k])
		} else {
			logger.SetPrefix(fmt.Sprintf("[%s]", k))
		}
	}

	if DebugLevel < level {
		debugLog.SetOutput(io.Discard)
	}
	if InfoLevel < level {
		infoLog.SetOutput(io.Discard)
	}
	if WarnLevel < level {
		warnLog.SetOutput(io.Discard)
	}
	if ErrorLevel < level {
		errorLog.SetOutput(io.Discard)
	}
	if FatalLevel < level {
		fatalLog.SetOutput(io.Discard)
	}
}

func File(path string) *os.File {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file error, err = %v\n", err)
		return nil
	}
	return file
}
