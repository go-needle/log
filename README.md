<!-- markdownlint-disable MD033 MD041 -->
<div align="center">

# 🪡log

<!-- prettier-ignore-start -->
<!-- markdownlint-disable-next-line MD036 -->
A lightweight Golang logging framework
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
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
	log.Set(log.ErrorLevel, log.File("log.txt"))
	log.Debug("debug")
	log.Error("error")
	log.Fatal("fatal")
	log.Info("info")
	log.Warn("warn")
}
```