package main

import (
	"fmt"
)

// 1. 定义一种函数类型
type Callback func(result string) error

// 2. 这是一个普通函数
// 它的参数和返回值正好符合 Callback
func printResult(result string) error {
	fmt.Println("收到处理结果：", result)
	return nil
}

// 3. 这个函数接收一个 Callback
func doWork(data string, callback Callback) {

	fmt.Println("开始处理数据：", data)

	// 假装这里做了一些业务处理
	result := "处理完成：" + data

	// 4. 工作完成以后，调用别人传进来的函数
	err := callback(result)

	if err != nil {
		fmt.Println("回调执行失败：", err)
	}
}

func main34() {

	// 5. 把 printResult 函数传进去
	doWork("玩家充值100元", printResult)
}
