package main

import "fmt"

func main19() {
	// 切片发生拷贝后，修改一个切片的元素可能也修改了另一个切片

	// arr := make([]int, 3, 5)

	// fmt.Printf("%#v\n",arr)
	// fmt.Println(len(arr))
	// fmt.Printf("%p\n",&arr[0])

	// arr = append(arr,8 )
	// fmt.Printf("%p\n",&arr[0])
	// fmt.Println(len(arr))

	arr := make([]int, 3, 5)
	brr := arr

	brr = append(arr,8 )
	fmt.Println(len(brr))
	fmt.Println(len(arr))    // 需要注意的是arr的len仍然是3 ，8这个元素arr是看不见的

	arr  = append(arr, 7)
	fmt.Println(brr[3])    //此时brr被arr修改为7，7并且对brr是可见的


	brr = append(brr, 8)
	brr = append(brr, 8)

	fmt.Println(cap(brr))   //此时brr进行了扩容，此时arr 和brr毫不相干


	
}