package main

import (
	"fmt"
	"time"
)

// 这定义了类型和返回值
func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	//Debug, LogLevel, startUpTime := getConfig()     // 用函数的返回值作为初始值
	var Debug, LogLevel, startUpTime = getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
}
