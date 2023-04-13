package main

import "fmt"

func main() {
	fmt.Println("hello world")

	// KHAI BÁO BIẾN
	var name string = "Bui Hien"
	fmt.Printf("name : %s", name)

	// Khai báo biến không cần chỉ định kiểu dữ liệu
	var age = 19
	fmt.Println(age)

	// Khai báo nhiều biến
	var (
		fullName = "naveen"
		year     = 29
		bmi      int
	)
	fmt.Println(fullName, year, bmi)

	var number1, number2 = 10, 20
	fmt.Println(number1, number2)

	var width, height int = 100, 50
	fmt.Println(width, height)

	var firstName, isStatus = "Nguyễn", true
	fmt.Println(firstName, isStatus)

	lastName, isQuit := "Hien", false
	fmt.Println(lastName, isQuit)

	// Khai báo nhanh, không sử dụng từ khóa 'var'
	email := "hien@techmaster.vn"
	fmt.Println(email)
}
