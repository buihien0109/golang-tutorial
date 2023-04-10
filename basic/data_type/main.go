package main

import "fmt"

func main() {
	fmt.Println("Data type")

	/* CÁC KIỂU DỮ LIÊU
	- bool
	- Numeric Types
		int8, int16, int32, int64, int
		uint8, uint16, uint32, uint64, uint
		float32, float64
		complex64, complex128
		byte
		rune
	- string
	*/

	// BOOL
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	// NUMBER
	/*
			(Signed integers : Kiểu số nguyên bao gồm cả số âm và số dương)

			- int8: represents 8 bit signed integers
			size: 8 bits
			range: -128 to 127

			- int16: represents 16 bit signed integers
			size: 16 bits
			range: -32768 to 32767

			- int32: represents 32 bit signed integers
			size: 32 bits
			range: -2147483648 to 2147483647

			- int64: represents 64 bit signed integers
			size: 64 bits
			range: -9223372036854775808 to 9223372036854775807

			- int: represents 32 or 64 bit integers depending on the underlying platform.
			You should generally be using int to represent integers unless there is a need to use a specific sized integer.
			size: 32 bits in 32 bit systems and 64 bit in 64 bit systems.
			range: -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems

		(32, 64 bit phụ thuộc vào nền tảng đang sử dụng. Nếu không có gì đặc biệt thì chỉ cần khai báo "int" là đủ)

		Tương tự với uint : Kiểu số nguyên dương cũng thế
	*/

	var (
		number       = 10
		number1 int8 = 20
		number2 int16
		number3 int32
		number4 int64

		age  uint
		age1 uint8
		age2 uint16
		age3 uint32
		age4 uint64
	)

	fmt.Println(number, number1, number2, number3, number4)
	fmt.Println(age, age1, age2, age3, age4)

	// STRING
	var name string
	fmt.Println(name)

	fullName := "Bùi Văn Hiên"
	fmt.Println(fullName)
}
