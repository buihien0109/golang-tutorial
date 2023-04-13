package main

import (
	"fmt"
	"github.com/buihien0109/golang-tutorial/basic/simple_package"
)

func init() {
	fmt.Println("Đây là function init trong package main")
}

func main() {
	// Sử dụng thông tin ở 1 package khác
	fmt.Println(simple_package.Message)
	fmt.Println(simple_package.GetInfo("Bùi Hiên", 1997))
}
