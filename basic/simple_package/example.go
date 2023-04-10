package simple_package

import (
	"fmt"
	"time"
)

/*
Mỗi package có chứa 1 init function (không có dữ liệu trả về và không có tham số)
Init func không được gọi 1 cách thủ công (init()) mà được tự động gọi khi package được khởi tạo

Thứ tự khởi tạo package như sau:
	1. Các biến level package được khởi tạo đầu tiên
	2. init func ược gọi tiêp theo. 1 packge có thể có nhiều init func (trong 1 file hoặc nhiều file)
	thứ tự gọi do trình biên dịch

- Khi package được import vào trong package khác -> package được import sẽ được khởi tạo trước
- 1 package sẽ được khởi tạo duy nhất 1 lần nếu được import vào trong nhiều package

*/

func init() {
	fmt.Println("Đây là function init 1 trong package simple_package")
}

func init() {
	fmt.Println("Đây là function init 2 trong package simple_package")
}

func init() {
	fmt.Println("Đây là function init 3 trong package simple_package")
}

/*
Note : Khi cần public 1 biến, 1 function để có thể sử dụng trong package khác thì cần được viết hoa chữ cái đầu
Ví dụ :
	- name -> Name
	- sum() -> Sum()
*/

var (
	Message = "Đây là nội dung message trong simple package"
	count   = 10
	posts   = []string{"Những ngày vui vẻ", "Một ngày nào đó", "Những mảnh ký ức"}
)

// Sử dụng inside packge
func getOld(year int) int {
	now := time.Now()
	currentYear := now.Year()
	return currentYear - year
}

// GetInfo Sử dụng outside, in other package
func GetInfo(name string, year int) string {
	old := getOld(year)
	info := fmt.Sprintf("Xin chào các bạn, tôi tên là %s. Năm nay tôi %d tuổi", name, old)
	return info
}
