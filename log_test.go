package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
	file := File("log.txt")
	defer file.Close()
	Set(ErrorLevel, file)
	Debug("debug")
	Error("error")
	Info("info")
	Warn("warn")
	Fatal("fatal")
}

func TestLogObj(t *testing.T) {
	l := New()
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")
	file := File("log.txt")
	defer file.Close()
	l.Set(ErrorLevel, file)
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")
	l.Fatal("fatal")
}
