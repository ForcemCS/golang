package main

import "fmt"

func array1() {
	nums := [3]int{10, 20, 30}

	for index, value := range nums {
		fmt.Println(index, value)
	}
}

func array2() {

	//  相当于二维数组，两行三列
	//   +----+----+----+
	//   | 0  | 0  | 0  |
	//   +----+----+----+
	//   | 0  | 0  | 0  |
	//   +----+----+----+
	var nums1 [2][3]int

	nums1 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(nums1)

}

func array3() {

	// 元素是匿名结构体
	var st = [...]struct {
		name string
		age  int
	}{{"a", 10}, {"b", 20}}

	fmt.Printf("%+v\n", st)
	fmt.Println(st[0].name)
	fmt.Printf("%p\n", &st)
}

func array4() {
	nums := [3]int{10, 20, 30}

	for index, value := range nums {
		nums[index] += value
	}

	fmt.Println(nums)
}

func main16() {
	// array1()
	// array2()
	// array3()

	array4()

}
