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
	var start, middle, end float32
	fmt.Println(start, middle, end)

	var name, left, right = "one", 1, 2
	fmt.Println(name, left, right)

	var Debug, LogLevel, startUpTime = getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)

}
