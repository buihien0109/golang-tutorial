package main

import (
	"fmt"
	"strings"
)

func main() {
	// Tạo nil slice (slice rỗng)
	var sliceEmpty []int
	sliceEmpty1 := make([]int, 0)
	sliceEmpty2 := []int{}

	fmt.Println(sliceEmpty, sliceEmpty1, sliceEmpty2)

	// Slice không chứa bất kỳ data nào
	// Nó tham chiếu đến array đang tồn tại
	a := [5]int{76, 77, 78, 79, 80}
	var b = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b)

	c := []int{6, 7, 8} // Tạo array và trả về 1 slice tham chiếu
	c = append(c, 9)    // Thêm phần tử vào slice sử dụng append
	fmt.Printf("type of : %T", c)
	fmt.Println(c)

	// Tham chiếu 1 slice
	// Khi thay đổi giá trị trong slice -> thay đổi giá trị trong array được tham chiếu
	darr := [...]int{1, 2, 3, 4, 5, 6, 7}
	dslice := darr[2:5]
	fmt.Println("array before", darr) // [1 2 3 4 5 6 7]
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr) // [1 2 4 5 6 6 7]

	// Ví dụ 2
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa) // [78 79 80]
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa) // [100 79 80]
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa) // [100 101 80]
	fmt.Println("nums1", nums1)
	fmt.Println("nums2", nums2)

	// Ví dụ 3
	// Tạo một slices string
	// Bao gồm độ dài và sức chứa của 5 phần tử
	slice := make([]string, 5)
	fmt.Println(slice)

	// Bao gồm độ dài là 3 và sức chứa sẵn sàng là 5
	slice1 := make([]int, 3, 5)
	fmt.Println(slice1)

	// Tạo slice với value của các index
	sliceNames := []string{0: "hoho", 9: "hihi", 10: "haha"}
	fmt.Println("sliceNames", sliceNames)

	// length and capacity of a slice
	/*
		slice[i:j] và sức chứa sẵn sàng là k.
		Độ dài: j - i
		Sức chứa sẵn sàng: k - i

		Vậy trong trường hợp ví dụ trên của mình sẽ là slice[1:3]
		Độ dài: 3 - 1 = 2
		Sức chứa sẵn sàng: 7 - 1 = 6
	*/

	// Ví dụ 1
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6

	// Ví dụ 2
	names := []string{"Hien", "Hoa", "An"}
	fmt.Printf("length of slice %d capacity %d\n", len(names), cap(names))

	newNames := append(names, "Hằng")
	fmt.Printf("length of slice %d capacity %d\n", len(newNames), cap(newNames))

	newNames = append(newNames, "A", "B", "C")
	fmt.Printf("length of slice %d capacity %d\n", len(newNames), cap(newNames))

	newNames = append(newNames, strings.Split(strings.Repeat("D", 6), "")...)
	fmt.Printf("length of slice %d capacity %d\n", len(newNames), cap(newNames))

	newNames = append(newNames, strings.Split(strings.Repeat("E", 12), "")...)
	fmt.Printf("length of slice %d capacity %d\n", len(newNames), cap(newNames))

	newNames = append(newNames, strings.Split(strings.Repeat("F", 25), "")...)
	fmt.Printf("length of slice %d capacity %d\n", len(newNames), cap(newNames))
	/*
		Kết quả :
		length of slice 3 capacity 3
		length of slice 4 capacity 6
		length of slice 7 capacity 12
		length of slice 13 capacity 24
		length of slice 25 capacity 48
		length of slice 50 capacity 96
	*/
}
