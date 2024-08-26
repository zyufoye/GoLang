package main

import "fmt"

// 声明全局变量时，方法1，2，3是可以的
var gA int = 100

//但是用方法四去声明全局变量会报错，mod := 不能声明全局变量

// 声明变量的四种方式
func main() {
	//方法一：声明一个变量,如果不赋初值，则默认值为0
	var a int
	fmt.Println("a =", a)
	fmt.Printf("type of a is %T\n", a)

	//方法二：声明一个变量，给一个初始化值
	var b int = 100
	fmt.Println("b =", b)
	fmt.Printf("type of b is %T\n", b)

	//方法三：在初始化时省去数据类型，通过值自动匹配当前变量的数据类型
	var c = 1000
	fmt.Println("c =", c)
	fmt.Printf("type of c is %T\n", c)

	var d = "abcd"
	fmt.Printf("d is %s,type of d is %T\n", d, d)

	//方法四：最常见的方法，省去var，直接自动匹配
	e := 100
	fmt.Println("e =", e)
	fmt.Printf("type of e is %T\n", e)

	//声明多个变量
	var x, y int = 100, 200
	fmt.Println("x = ", x, ", y = ", y)
	var k, l = 100, "abcd"
	fmt.Println("k =", k, "l =", l)

	//多行多变量声明
	var (
		aa int     = 100
		bb float32 = 0.01
	)
	fmt.Println("aa =", aa, "bb = ", bb)

	//打印全局变量
	fmt.Println("Global gA =", gA)

}
