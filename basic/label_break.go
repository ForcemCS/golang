package main

import "fmt"

// break只是推出内层循环，外层循环仍然继续
func basic_break() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				break
			}

			fmt.Println(i, j)
		}
	}
}

func label_goto() {

OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				break OuterLoop //直接退出最外层
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("循环结束")
}

func main12() {
	// basic_break()
	label_goto()
}
