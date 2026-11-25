package main

import "fmt"

func main9() {

	var sum int

	// for a := 6; a > 0; a-- {
	// 	sum  +=  a
	// }

	// a :=6
	// for ; a > 0; a-- {
	// 	sum  +=  a
	// }

	// a :=6
	// for ; a > 0; a-- {
	// 	if  a == 2  {
	// 		break
	// 	}
	// 	sum  +=  a
	// }

	a := 6
	for ; a > 0; a-- {
		if a == 2 {
			continue
		}
		sum += a
	}

	fmt.Println(sum)

	for {
		var input string

		fmt.Print("请输入内容，输入 exit 退出：")
		fmt.Scanln(&input)

		if input == "exit" {
			fmt.Println("程序结束")
			break
		}

		fmt.Println("你输入的是：", input)
	}
}
