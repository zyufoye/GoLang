## 0x00. 第一个Go程序
```Golang
package main//程序包名

import "fmt" //单个包导入方式

//多个包导入方式
import (
    "fmt"
    "time"
)

func main(){ //Go要求函数左括号必须在和func在同一行
    //Go表达式推荐不加";"结尾，但加不加都可以
    fmt.println("hello,world!")
}
```
## 0x01. 变量声明
在Go中，声明一个变量有四种方式，首先是声明一个默认值为0的变量，然后在声明的同时，我们可以为变量赋值来初始化这个变量，在初始化时可以省去数据类型，通过值自动匹配当前变量的数据类型。最后也是最常用的方法是省去var关键字，直接用":="赋值自动匹配。  
多变量声明方式分为单行写法和多行写法。  
在声明全局变量时，方法1，2，3都是可以的，但是不能用:=直接赋值。
```Golang
//方法1：声明一个变量
var a int

//方法2：声明变量时初始化
var b int = 100

//方法3：在初始化时省去数据类型，通过值自动匹配当前的数据类型
var c = 100

//方法4：常用方法，直接自动匹配
d := 100

//多变量声明方式-单行写法
var x,y int = 100,200

//多变量声明方式-多行写法
var (
		m  int     = 100
		n  float32 = 0.01
	)

//全局变量声明，即在main外声明，方法1，2，3均可，方法4报错
var gA int = 100
```
## 0x02. 常量声明
在Go中，常量声明需要关键字const，常量一经声明不得修改。  
常量可以用来定义枚举类型，同时有一个iota特性。iota只能配合const一起使用，iota在const中有累加效果。
```Golang
//常量(只读属性)
const length int = 10

//const 来定义枚举类型
const (
	//可以在const() 添加一个关键字 iota， 每行的iota都会累加1, 第一行的iota的默认值是0
	BEIJING = 10*iota	 //iota = 0
	SHANGHAI 		  //iota = 1
	SHENZHEN          //iota = 2
)

const (
	a, b = iota+1, iota+2 // iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
	c, d				  // iota = 1, c = iota + 1, d = iota + 2, c = 2, d = 3
	e, f				  // iota = 2, e = iota + 1, f = iota + 2, e = 3, f = 4

	g, h = iota * 2, iota *3  // iota = 3, g = iota * 2, h = iota * 3, g = 6, h = 9 
	i, k					   // iota = 4, i = iota * 2, k = iota * 3 , i = 8, k = 12
)
```
## 0x03. 
