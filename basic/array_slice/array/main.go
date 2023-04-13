package main

import (
	"fmt"
)

func main() {
	// ARRAY
	// Khai báo mảng trống
	var a [3]int // Khai báo mảng với 3 giá trị, nếu không có khởi ta => nhận default value
	fmt.Println(a)

	a[0] = 10
	a[1] = 20
	fmt.Println(a)

	// Vừa khai báo, vừa khoi tạo giá trị
	var b = [3]int{1, 2, 3}
	fmt.Println(b)

	var c = [4]string{"Hiên"} // Khởi tạo 1 giá trị đầu tiên -> các giá tr còn lại nhận default value
	fmt.Println(c)

	var d = [...]int{1, 2, 3, 4}
	fmt.Println(d)

	var e = [...]int{}
	fmt.Println(e)

	// Array trong go là value type
	names := [...]string{"USA", "China", "India", "Germany", "France"}
	namesOther := names //
	namesOther[0] = "Singapore"
	fmt.Println("names is ", names)
	fmt.Println("namesOther is ", namesOther)

	// Truyền array vào trong func xem có thay đổi gía trị của array không
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num) // [5 6 7 8 8]
	changevalue(num)                                //num is passed by value
	fmt.Println("after passing to function ", num)  // [5 6 7 8 8]

	// Length of array
	fmt.Println("length of num is", len(num))

	// Loop in array
	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
	}

	// Loop use for range
	for _, v := range num {
		fmt.Println(v)
	}
	// SLICE
}

func changevalue(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num) // [55 6 7 8 8]
}
