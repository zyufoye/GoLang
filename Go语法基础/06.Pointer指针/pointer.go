package main

import "fmt"

func swap(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}

func swap_pointer(pa *int, pb *int) {
	var temp int
	temp = *pa //temp = main::a
	*pa = *pb  // main::a = main::b
	*pb = temp // main::b = temp
}

func main() {
	var a int = 10
	var b int = 20

	swap(a, b)

	fmt.Println("a = ", a, " b = ", b)

	swap_pointer(&a, &b)

	fmt.Println("a = ", a, " b = ", b)

	var p *int

	p = &a

	fmt.Println(&a)
	fmt.Println(p)
	//&a 和 p地址一样
	//0xc000096068
	//0xc000096068

	var pp **int //二级指针

	pp = &p

	fmt.Println(&p)
	fmt.Println(pp)
	//二级指针存放指针的地址
	//0xc000098028
	//0xc000098028
}
