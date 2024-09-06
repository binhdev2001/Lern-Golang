package main

import (
	"fmt"
)

func main() {

	demo_slice()
}

// Cấu trúc câu lệnh if else
func demo_if_else() {
	//cấu trúc câu lệnh điều kiện if trong go
	x := 10
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than or equal to 10")
	}
}

// Hàm defer trong go
func demo_defer() {
	// Ví dụ áp dụng defer trong go lang
	fmt.Println("Start")

	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")

	fmt.Println("End")
}

// Ví dụ về pointer trong Go

func demo_pointer() {
	i, j := 42, 37

	p := &i                                        // point to i
	fmt.Println("P giờ đang tham chiếu đến i", *p) // read i through the pointer
	*p = 21                                        // set i through the pointer
	fmt.Println("I đang nhận giá trị mới từ P", i) // see the new value of i

	p = &j                                                                   // point to j
	*p = *p / 37                                                             // divide j through the pointer
	fmt.Println("J nhận giá trị mới từ P sau khi P thực hiện phép tính ", j) // see the new value of j
}

// Khai báo mảng trong go lang
func demo_array() {
	// Khai báo mảng trong go
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

// Ví dụ về slice trong Go Lang
func demo_slice() {
	// Khai báo slice trong go
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println("lấy ra bốn phần tử trong mảng nhưng bỏ phần tử đầu tiên:", s)
}

// Demo length và câpcity của slice
// Capacity của slice là số lượng phần tử mà slice có thể chứa
// Length của slice là số lượng phần tử mà slice đang chứa
func demo_length_capacity() {
	// Khai báo slice trong go
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Printf("s=%v, len=%d, cap=%d\n", s, len(s), cap(s))
}
