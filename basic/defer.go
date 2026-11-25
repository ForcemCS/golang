package main

import (
	"fmt"
	"time"
)

// third: 9
// second: 9
// first: 5
func test() int {
	defer fmt.Println("defer") //在return之前(扔给外边的调用者之前)执行

	fmt.Println("return 前")

	return 100
}

func defer_exe_time() (i int) {

	i = 9
	defer func() {
		fmt.Printf("first: %d\n", i) //先执行注册
	}()

	defer func(i int) {
		fmt.Printf("second: %d\n", i)
	}(i)

	defer fmt.Printf("third: %d\n", i)

	return 5
}

func timeOfWork(arg int) int {
	begin := time.Now()

	defer func() { fmt.Printf("耗时： %f\n", time.Since(begin).Seconds()) }()

	if arg > 10 {
		time.Sleep(2 * time.Second)

		return 100

	} else {
		time.Sleep(3 * time.Second)

		return 200
	}

}

func main27() {
	// x := 10
	// // defer登记的是一个整体，需要注意
	// defer fmt.Println(x)
	// x = 20

	// x := 10

	// defer func() {
	// 	fmt.Println(x)
	// }()

	// x = 20

	defer_exe_time()
	timeOfWork(5)
}
