package main

import "fmt"

func deferTest() int {
	fmt.Println("defer called...")
	return 0
}
func returnTest() int {
	fmt.Println("return called...")
	return 0
}
func returnandDeferTest() int {
	defer deferTest()
	return returnTest()
}
func main() {
	//写入defer关键字
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")

	returnandDeferTest()
}

//main::hello go 1
//main::hello go 2
//return called...
//defer called...
//main end2
//main end1
