package main

import "fmt"

func main18() {
	s1 := []int{10, 20, 30}

	// 此时引用的底层数组发生了变化，之前是len ==cap么有多余的空间了
	// 申请一个更大的数组
	//         ↓
	// 复制旧数据
	//         ↓
	// 添加新元素
	s2 := append(s1, 40)

	for i, v := range s2 {
		fmt.Println(i, v)
	}

}
