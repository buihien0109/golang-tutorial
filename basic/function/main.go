package main

import "fmt"

func init() {
	fmt.Println("Đây là function init trong package main")
}

func main() {
	fmt.Println(sum(10, 20))
	fmt.Println(mutil(10, 20))
	sayHello()
	sayHelloWithName("Hien")
	fmt.Println(calcalateVolumeRectangle(4, 5))

	area, perimeter := rectProps(5.4, 6.4)
	fmt.Println(area, perimeter)

	area1, _ := rectProps(5.4, 6.4)
	fmt.Println(area1)

	_, perimeter1 := rectProps(5.4, 6.4)
	fmt.Println(perimeter1)

	area2, _ := rectPropsWithName(5.4, 6.4)
	fmt.Println(area2)

	_, perimeter2 := rectPropsWithName(5.4, 6.4)
	fmt.Println(perimeter2)
}

func sum(a, b int) int {
	return a + b
}

func mutil(a, b int) (result int) {
	result = a * b
	return
}

func calcalateVolumeRectangle(width int, height int) int {
	return width * height
}

func sayHello() {
	fmt.Println("Xin chào các bạn")
}

func sayHelloWithName(name string) {
	message := fmt.Sprintf("Xin chào %s", name)
	fmt.Println(message)
}

// Return mutilple value
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// Return mutilple value with name
func rectPropsWithName(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}
