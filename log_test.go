package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
	Set(ErrorLevel, File("log.txt"))
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
	l.Set(ErrorLevel, File("log.txt"))
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")
	l.Fatal("fatal")
}
