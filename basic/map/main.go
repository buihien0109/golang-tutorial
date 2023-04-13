package main

import "fmt"

func main() {
	// Khai báo map
	var employeeSalary = make(map[string]int)
	fmt.Printf("employeeSalary type : %T\n", employeeSalary)

	// Thêm dữ liệu
	employeeSalary["Bui Hien"] = 1200
	employeeSalary["Thu Hang"] = 2300
	employeeSalary["Minh Duy"] = 1500

	fmt.Println("employeeSalary", employeeSalary)

	// Truy cập dữ liệu trong map
	fmt.Println(employeeSalary["Bui Hien"])

	// Vừa khai báo, vừa khởi tạo giá trị
	var employeeEmail = map[string]string{
		"Bui Hien": "hien@gmail.com",
		"Thu Hang": "hang@gmail.com",
		"Minh Duy": "duy@gmail.com",
	}
	fmt.Println("employeeEmail", employeeEmail)

	// Zero value của map là "nil"
	var employeeAge map[string]int
	fmt.Println(employeeAge)

	// Nếu thêm dữ liệu vào trong map -> error
	// employeeAge["Bui Hien"] = 24 // panic: assignment to entry in nil map

	// Checking if a key exists
	newEmp := "Bui Hien"
	value, ok := employeeEmail[newEmp]
	if ok == true {
		fmt.Println("Email of", newEmp, "is", value)
	}

	// Lặp qua các phần tử của map
	for key, val := range employeeEmail {
		fmt.Printf("employeeEmail[%s] = %s\n", key, val)
	}

	// Xóa items trong map
	fmt.Println("employeeEmail trước khi xóa", employeeSalary)
	delete(employeeSalary, "Bui Hien")
	fmt.Println("employeeEmail sau khi xóa", employeeSalary)

	// Độ dài của map
	fmt.Println("Length of employeeSalary is", len(employeeSalary))

	// Map là kiểu tham chiếu
	fmt.Println("Original employee salary", employeeSalary)
	modified := employeeSalary
	modified["Thu Hang"] = 1800
	fmt.Println("Employee salary changed", employeeSalary)

	// NOTE : Không thể so sanh map (==) mà chỉ có thể so sánh map với nil
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	map2 := map1

	//if map1 == map2 {} -> error nếu so sánh 2 map với nhau
	if map2 == nil {
	}
}
