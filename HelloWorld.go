package main

import (
	"Demo1/Demo"
	"fmt"
)

func main() {
	// Kiểu dữ liệu bool
	isTrue := true
	isFalse := false
	fmt.Println("isTrue", isTrue)
	fmt.Println("isFalse", isFalse)
	a := isTrue && isFalse // Toán tử và
	fmt.Println("a:", a)
	b := isTrue || isFalse // Toán tử hoặc
	fmt.Println("b:", b)
	fmt.Println("----------------------")
	// Kiểu dữ liệu số
	number := 10
	fmt.Println(number)
	// Kiểu dữ liệu string
	greet := "Hello"
	fmt.Println(greet)

	// Access Modifier
	Demo.TestAccessModifier() // Khác pakage nhưng vẫn truy cập được

	// Khai báo biến
	var age int // Không giá trị khởi tạo
	fmt.Println("my age is", age)
	var age1 int = 29 // khai báo biến có giá trị khởi tạo
	fmt.Println("my age is", age1)
	var age2 = 29 // Kiểu biến được tự suy
	fmt.Println("my age is", age2)
	var c, d int = 100, 50 // khai báo nhiều biến cùng kiểu
	fmt.Println("c is", c, "d is", d)
	var (
		name   = "vuong"
		age4   = 29 // Khai báo nhiều biến khác kiểu
		height int
	)
	fmt.Println("my name is", name, ", age is", age4, "and height is", height)
	animal, weight := "chicken", 3 // khai báo nhanh
	fmt.Println("my animal is", animal, "weight is", weight)

	// Hằng số
	const e = 93
	// e = 89 //không thể gán lại giá trị của một hằng số
	const f int = 100
	const hello string = "Hello World"

	// for loops
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
	fmt.Println()
	// sử đụng breaks
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //vòng lặp dừng nếu i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n Câu lệnh ngay sau vòng lặp \n")
	// sử dụng continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	// sử dụng nhiều biến trong vòng lặp
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	// for range
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// If - else
	if true {
		fmt.Println("Yes!")
	}

	l := 1993

	if l > 1993 {
		fmt.Println(">")
	} else if l < 1993 {
		fmt.Println("<")
	} else {
		fmt.Println("=")
	}

	// Switch - case
	x := 93
	switch x {
	case 0:
	case 1, 2, 3:
		fmt.Println("Multiple matches")
	case 93:
		fmt.Println("Reached")
		fallthrough
	case 94:
		fmt.Println("Unreached")
		fallthrough
	case 95:
		fmt.Println("1 more turn")
	default:
		fmt.Println("Optional")
	}

}
