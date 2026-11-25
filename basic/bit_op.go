package main

import (
	"fmt"
)

func binaryFormat1(a uint64) {
	var c uint64 = 1 << 63
	for i := 0; i < 64; i++ {
		if c&a == c {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
		c = c >> 1
	}
}

func binaryFormat2(a uint64) {
	var c uint64 = 1 << 63
	started := false

	for i := 0; i < 64; i++ {
		if a&c != 0 {
			started = true
			fmt.Print("1")
		} else if started {
			fmt.Printf("0")
		}

		c = c >> 1
	}

	if !started {
		fmt.Print("0")
	}
}

const (
	StatusOnline   = 1 << 0 // 0001
	StatusMuted    = 1 << 1 // 0010
	StatusVIP      = 1 << 2 // 0100
	StatusFighting = 1 << 3 // 1000
)

func main15() {

	var a int = 5
	fmt.Printf("%b\n", a) // 转为二进制

	var status int

	// 玩家上线
	//status |= StatusOnline
	status = status | StatusOnline

	// 玩家成为 VIP
	status |= StatusVIP

	fmt.Printf("%04b\n", status) // 0101

	// 判断是否在线
	if status&StatusOnline != 0 {
		fmt.Println("玩家在线")
	}

	// 判断是否被禁言
	if status&StatusMuted != 0 {
		fmt.Println("玩家被禁言")
	} else {
		fmt.Println("玩家没有被禁言")
	}

	binaryFormat2(10)

}
