package main

import (
	"fmt"
	"slices"
)

func main21() {

	// --------------------
	s := []int{5, 2, 8, 1, 3}
	slices.Sort(s) //原地从小到大排序

	fmt.Println(s)

	// 从大到小排序
	slices.SortFunc(s, func(a, b int) int {
		return b - a
	})
	fmt.Println(s)

	// ---------------
	type User struct {
		Age    int
		Height float64
	}

	brr := []*User{
		{10, 22.3},
		{9, 19.3},
	}
	slices.SortFunc(brr, func(a, b *User) int {
		if a.Height > b.Height {
			return 1
		} else if a.Height < b.Height {
			return -1
		} else {
			return 0
		}
	})

	fmt.Println(*brr[0])

	// ---------------------------------

	fmt.Println("最大者: ", slices.Max(s))
	fmt.Println("最小者: ", slices.Min(s))
	fmt.Println(slices.Contains(s, 5))

	// --------------

	crr := make([]int, len(s))

	copy(crr, s) // 最多只能拷贝len(crr)个元素,性能很高，但是底层数组是不同的内存空间
	fmt.Printf("%v\n", crr)

	fmt.Println(slices.Equal(s, crr)) // true 比较的是数组中的元素内容
}
