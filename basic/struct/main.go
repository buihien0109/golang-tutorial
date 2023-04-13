package main

import (
	"fmt"
	"github.com/buihien0109/golang-tutorial/basic/struct/sub"
)

// Struct là cấu trúc dữ liệu do người người dùng tự định nghĩa
func main() {
	// Tạo đối tượng mà không có dữ liệu
	newEmployee := Employee{}
	fmt.Println(newEmployee)

	// Tạo đối tượng có dữ liệu, nhưng bỏ qua key mà chỉ điền value
	employee := Employee{"Bùi", "Hiên", 26}
	fmt.Println(employee)

	// Tạo đối tượng có dữ liệu với đầy đủ key và value
	employee1 := Employee{
		firstName: "Thu",
		lastName:  "Hằng",
		age:       24,
	}
	fmt.Println(employee1)

	// Tạo anonymous struct
	student := struct {
		id    int
		name  string
		email string
	}{
		id:    1,
		name:  "Bùi Hiên",
		email: "hien@gmail.com",
	}
	fmt.Println(student)

	// Truy cập các field trong struct
	fmt.Println("employee", employee)
	fmt.Println("firstName", employee.firstName)
	fmt.Println("lastName", employee.lastName)
	fmt.Println("age", employee.age)

	// Zero value của struct
	var employee2 Employee
	fmt.Println("employee2", employee2)
	fmt.Println("firstName", employee2.firstName)
	fmt.Println("lastName", employee2.lastName)
	fmt.Println("age", employee2.age)

	// Pointers to a struct
	employee3 := &Employee{
		firstName: "Hồng",
		lastName:  "Duy",
		age:       40,
	}
	fmt.Println("First Name:", (*employee3).firstName)
	fmt.Println("Last Name:", (*employee3).lastName)
	fmt.Println("Age Before:", (*employee3).age)
	(*employee3).age = 30
	fmt.Println("Age After:", (*employee3).age)
	fmt.Println("employee3", *employee3)

	// Có thể thay thế pointer to struct như sau
	fmt.Println("First Name:", employee3.firstName)
	fmt.Println("Last Name:", employee3.lastName)
	fmt.Println("Age:", employee3.age)

	// Anonymous fields
	s1 := Student{
		string: "naveen",
		int:    50,
	}
	fmt.Println(s1.string)
	fmt.Println(s1.int)

	// Nested structs
	p := Person{
		name: "Bùi Hiên",
		age:  26,
		address: Address{
			city:  "Hà Nội",
			state: "Nam Từ Liêm",
		},
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)

	// Promoted fields
	staff := Staff{
		id:   1,
		name: "Bùi Hiên",
		Address: Address{
			city:  "Thái Bình",
			state: "Thái Thụy",
		},
	}
	fmt.Println("Id:", staff.name)
	fmt.Println("Name:", staff.name)
	fmt.Println("City:", staff.city)
	fmt.Println("State:", staff.state)

	// Exported structs and fields
	post := sub.Post{
		Title:  "Những bài viết hay nhất",
		Author: "Thu Hằng",
	}
	fmt.Println(post)
}

// Nested structs
type Address struct {
	city  string
	state string
}

type Person struct {
	name    string
	age     int
	address Address
}

type Staff struct {
	id   int
	name string
	Address
}

// Anonymous fields không có tên field xác định -> lấy tên kiểu dữ liệu làm tên field
type Student struct {
	string
	int
}

type Employee struct {
	firstName string
	lastName  string
	age       int
}
