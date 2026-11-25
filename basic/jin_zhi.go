package main

import (
	"fmt"
)

// 整数次方
func intPow(base, exponent int) int {
	result := 1

	for i := 0; i < exponent; i++ {
		result *= base
	}

	return result
}

// 26进制XFD表示多少
func main1() {
	var base int = 'Z' - 'A' + 1 // 进制

	var total int

	// total += 'D' - 'A' + 1
	// total += base * ('F' - 'A' + 1)
	// total += base * base * ('X' - 'A' + 1)

	total += intPow(base, 0) * ('D' - 'A' + 1)
	total += intPow(base, 1) * ('F' - 'A' + 1)
	total += intPow(base, 2) * ('X' - 'A' + 1)

	fmt.Printf("XFD表示的转为十进制后为 %d", total)

}
