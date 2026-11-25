package main

import "fmt"

// type
// ├── 自定义基础类型
// ├── struct
// ├── interface
// ├── 函数类型
// └── 类型别名

// ---1.-----
type UserID int64
type ServerID int32

// 自定义类型可以拥有方法
func (u UserID) Valid() bool {
	return u > 0
}

func LoadRole(uid UserID, sid ServerID) {
}

// -------2.-------------
type Role struct {
	UID           int64
	ServerID      int
	Name          string
	Level         int
	TotalRecharge float64
}

// ----3.------------

type Mover interface {
	Move(src string, dest string) error
}

func main33() {

	role := Role{
		UID:           100001,
		ServerID:      40224,
		Name:          "玩家A",
		Level:         86,
		TotalRecharge: 1288.50,
	}

	fmt.Println(role)

}
