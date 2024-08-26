package main

import "fmt"

func func1(a string, b int) int {
	fmt.Println(a)
	fmt.Println(b)
	c := 100
	return c
}

// 返回多个匿名的返回值
func func2(a int, b int) (int, int) {
	fmt.Println(a)
	fmt.Println(b)
	return b, a + b
}

// 返回多个有形参名称的返回值
func func3(a int, b int) (result1 int, result2 int) {
	result1 = a + b
	result2 = a - b
	return result1, result2
}

func func4(a int, b int) (result1 int, result2 int) {
	fmt.Println(a)
	fmt.Println(b)
	//给有名称的返回值变量赋值
	result2 = 2000
	result1 = 1000
	return
}

func func5(a int, b int) (result1, result2 int) {
	fmt.Println(a)
	fmt.Println(b)
	//给有名称的返回值变量赋值
	result2 = 3000
	result1 = 5000
	return
}

func main() {
	c := func1("abc", 100)
	fmt.Println("func1 c = ", c)

	d, e := func2(50, 100)
	fmt.Println("func2 d,e = ", d, e)

	f, g := func3(200, 100)
	fmt.Println("func3 f,g = ", f, g)

	h, i := func4(100, 200)
	fmt.Println("func4 h,i = ", h, i)

	j, k := func5(100, 200)
	fmt.Println("func5 j,k = ", j, k)
}
