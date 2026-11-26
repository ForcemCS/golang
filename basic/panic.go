package main

import "fmt"

// panic后，依次会执行：
// 1. 依次执行已经注册的defer,未注册的不会执行
// 2. 打印错误调用堆栈
// 3. 退出os.Exit(2)

// main start
// A
// defer 2
// defer 1
// recover: 程序出错
// main end
func test_p() {

	defer func() {
		defer fmt.Println("defer 1")

		if err := recover(); err != nil {
			fmt.Println("recover: ", err)
		}
	}()

	defer fmt.Println("defer 2")

	fmt.Println("A")

	panic("程序出错")

	//fmt.Println("B")

}

func main28() {

	fmt.Println("main start")
	test_p()
	fmt.Println("end")

}

// main start
// A
// defer 2
// recover:  程序出错
// defer 1
// end
