package main

import "fmt"

//住所
type Residence struct {
	Province string //省份
	City     string
}

type Users struct {
	Id            int
	Score         int
	Name, Address string
	Residence     Residence
}

func main5() {
	var u Users
	u.Score = 100
	u = Users{Id: 1, Name: "wukui", Address: "中国",}
	u = Users{Residence: Residence{Province: "山西", City: "大同"}}

	fmt.Printf("%+v\n", u)
	fmt.Println(u.Residence.City)
}
