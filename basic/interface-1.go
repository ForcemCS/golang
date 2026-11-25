package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "：汪汪")
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println(c.Name, "：喵喵")
}

func makeSpeak(s Speaker) {
	s.Speak()
}

func main31() {

	dog := Dog{Name: "旺财"}
	cat := Cat{Name: "咪咪"}

	makeSpeak(dog)
	makeSpeak(cat)
}
