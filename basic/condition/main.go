package main

import "fmt"

func main() {
	hour := 20
	// Câu lệnh if
	if hour < 12 {
		fmt.Println("Good morning")
	}

	// Câu lệnh if/else
	if hour < 12 {
		fmt.Println("Good morning")
	} else {
		fmt.Println("Good afternoon")
	}

	// Cau lệnh if/elseIf/else
	if hour < 12 {
		fmt.Println("Good morning")
	} else if hour >= 12 && hour <= 18 {
		fmt.Println("Good afternoon")
	} else {
		fmt.Println("Good evening")
	}

	// If with assignment
	if number := 10; number > 0 {
		fmt.Println("Số dương")
	}

	/*
		Trong GO không hỗ trợ toán tử 3 ngôi
		num := 10
		message := num % 2 == 0 ? "Số chẵn" : "Số lẻ"
	*/

	// Câu lệnh switch-case
	day := 0
	switch day {
	case 0:
		fmt.Println("Sunday")
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid day")
	}

	switch month := 10; month {
	case 1, 2, 3:
		fmt.Println("Spring")
	case 4, 5, 6:
		fmt.Println("Summer")
	case 7, 8, 9:
		fmt.Println("Autumn")
	case 10, 11, 12:
		fmt.Println("Winter")
	default:
		fmt.Println("Invalid month")
	}

	// Trong switch case có đánh giá điều kiện giống như trong câu lệnh if
	num := 75
	switch {
	case num >= 0 && num <= 50:
		fmt.Println("num lớn hơn 0 và nhỏ hơn 50")
	case num >= 51 && num <= 100:
		fmt.Println("num lớn hơn 51 và nhỏ hơn 100")
	case num >= 101:
		fmt.Println("num lớn hơn 100")
	}

	/*
			Trong Go, việc đi ra khỏi câu lệnh switch là ngay lập tức sau khi một case đúng được thực thi.
			Một lệnh fallthrough được sử dụng để chuyển quyền kiểm soát đến câu lệnh của case ngay sau nó.
			Điều này sẽ giúp tất cả các case đều được kiểm tra.

			Trong ví dụ dưới, mặc dụ case 2 thực thi đúng nhưng không kết thúc luôn bởi vì có "fallthrough"
			Sau khi có "fallthrough" -> Tiếp tục kiểm tra case tiếp

		Kết quả : in ra 2 trường hợp :
			- 75 thì nhỏ hơn 100
			- 75 thì nhỏ hơn 200

	*/
	switch num := 75; { //num không phải là một hằng số
	case num < 50:
		fmt.Printf("%d thì nhỏ hơn 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d thì nhỏ hơn 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d thì nhỏ hơn 200\n", num)
	case num < 300:
		fmt.Printf("%d thì nhỏ hơn 300\n", num)
	}

}
