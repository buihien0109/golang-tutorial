package main

import "fmt"

func main() {
	// "*T" là kiểu biến con trỏ trỏ đến giá trị kiểu T.
	// Toán tử "&" được sử dụng để lấy địa chỉ của một biến
	// Toán tử "*" được sử dụng để lấy giá trị của một biến
	var b int = 255
	var a *int = &b                    // a point to b (a trỏ đến b)
	fmt.Printf("Type of a is %T\n", a) // Type of a is *int
	fmt.Println("Address of b is", a)  // Address of b is 0xc0000240b8
	fmt.Println("Value of b is", *a)   // Value of b is 255

	// Zero value của pointer
	c := 25
	var d *int
	if d == nil {
		fmt.Println("d is", d) // d is <nil>
		d = &c
		fmt.Println("d after initialization is", d) // d after initialization is 0xc0000240d0
	}

	// Creating pointers using the new function
	size := new(int) // Tạo pointer kiểu int
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)

	// Passing pointer to a function
	number := 58
	fmt.Println("value of number before function call is", number) // 58
	pointer := &number
	changeValue(pointer)
	fmt.Println("value of number after function call is", number) // 55

	// Returning pointer from a function
	newPointer := returnPointer()
	fmt.Println("Value of newPointer", *newPointer) // 5
}

func changeValue(val *int) {
	*val = 55
}

func returnPointer() *int {
	i := 5
	return &i
}
