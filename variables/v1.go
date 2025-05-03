package main

import (
	"fmt"
	"time"
)

var (
	Debug       bool //省略初始值
	LogLevel    = "info"
	startUpTime = time.Now() // 省略类别
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}
