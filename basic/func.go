package main

import "fmt"

func sum1(nums ...int) int {
	total := 0

	for _, value := range nums {
		total += value
	}

	return total
}

func sum2(nums ...int) int {
	s := 0
	if len(nums) == 0 {
		return s
	}

	s += nums[0]
	s += sum2(nums[1:]...)

	return s
}

func main25() {

	a := []int{10, 20, 30}

	result1 := sum1(a...)
	fmt.Println(result1)

	result2 := sum2(a...)
	fmt.Println(result2)

}
