<!-- markdownlint-disable MD033 MD041 -->
<div align="center">

# 🪡log

<!-- prettier-ignore-start -->
<!-- markdownlint-disable-next-line MD036 -->
A lightweight logging framework for golang
<!-- prettier-ignore-end -->

<img src="https://img.shields.io/badge/golang-1.11+-blue" alt="golang">
</div>

## introduction
This is a lightweight Golang logging framework that supports five isolation levels: DEBUG, INFO, WARN, ERROR, and FATAL.

## installing
Select the version to install

`go get github.com/go-needle/log@version`

If you have already get , you may need to update to the latest version

`go get -u github.com/go-needle/log`


## quickly start
```golang
package main

import (
	"fmt"
	"github.com/go-needle/log"
	"os"
)

func main() {
	file := log.File("log.txt")
	defer file.Close()
	// global singleton
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
	log.Set(log.ErrorLevel, file)
	log.Debug("debug")
	log.Error("error")
	log.Info("info")
	log.Warn("warn")
	
	// obj
	l := log.New()
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")
	l.Set(log.ErrorLevel, file)
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")
	l.Fatal("fatal")
}
```
