package main

import (
	"fmt"
	"time"
)

func ticker(loc *time.Location) {
	fmt.Printf("现在的时间是: %s\n", time.Now().In(loc).Format("2006-01-02 15:04:05 -07:00"))

	tk := time.NewTicker(time.Second)

	for i := 0; i < 5; i++ {

		<-tk.C
		fmt.Printf("现在的时间是: %s\n", time.Now().In(loc).Format("2006-01-02 15:04:05 -07:00"))
	}

	tk.Stop()
}

// 2006-01-02 15:04:05

func main26() {

	//---time.Time表示的是一个时间点

	filename := "backup_" + time.Now().Format("20060102_150405") + ".sql"
	fmt.Println(filename)

	fmt.Println(time.Now().Format("03:04:05 PM"))

	s26 := "2026-08-16 18:30:00"
	// 字符串转为time.Time对象 默认按照UTC解析
	t1, err := time.Parse("2006-01-02 15:04:05", s26)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t1)
	fmt.Printf("%T", t1)

	//------时区
	fmt.Println(time.Now().Location())

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	t2, err := time.ParseInLocation("2006-01-02 15:04:05", s26, loc)
	if err != nil {
		panic(err)
	}

	fmt.Println(t2)

	//---------time.Duration 时间间隔
	var d1 time.Duration = 5 * time.Second
	fmt.Println(d1)

	var d2 time.Duration = 2*time.Hour + 30*time.Second
	fmt.Println(d2)

	// 涉及夏令时 DST 时，某一天可能 23 小时或者 25 小时。
	tomorrow := time.Now().AddDate(0, 0, 1)
	fmt.Println(tomorrow)

	ticker(loc)
}
