package main

import (
	"fmt"
	"time"
)

// 并发CPU不一定同时执行多个任务，表示的是他们可以同时推进，执行时间可能互相交错

func worker(id int, jobs <-chan int) {

	for job := range jobs {
		fmt.Printf("worker %d 开始处理任务 %d\n", id, job)

		time.Sleep(time.Second)

		fmt.Printf("worker %d 任务完成 %d\n", id, job)
	}
}

func main29() {

	// ch := make(chan int)  // 无缓冲channel

	// go func() {    //避免了main的阻塞
	// 	ch <- 100
	// }()

	// result  := <-ch

	// fmt.Println(result)

	jobs := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs)
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}

	close(jobs)

	time.Sleep(4 * time.Second)

}
