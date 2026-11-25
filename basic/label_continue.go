package main

import "fmt"
// 普通的 continue 只会跳过当前内层循环的本次执行。
// 标签 continue 可以直接开始外层循环的下一次迭代。
func label_conti() {

OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				continue OuterLoop   //只能放在一个for循环的外部
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("循环结束")
}

func main13() {
	label_conti()
}
