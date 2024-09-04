## 0x00. Go反射
在Go语言中，反射（Reflection）主要用于在运行时动态地检查变量的类型和值，以及对这些变量进行操作。反射在某些情况下是非常有用的，特别是当我们需要编写更加通用的代码，而这些代码在编译时无法确定具体的类型时。  
## 0x01. Go中变量的内部构造
Golang中变量的内部构造包含了两部分，一部分是变量类型type，另一部分是变量值value。类型还可以划分成静态类型static type，包含int、string等，还包含了具体数据类型concrete type（interface所指向的具体数据类型，系统看得见的类型），在Go中，每个变量是包含了type和value这个pair的。
```Golang
var a string
//pair<statictype:string, value:"aceld">
a = "aceld"
```
以上代码表示定义一个字符串，其中pair的type是string，value是aceld。我们定义了一个a，a的pair属性一直跟随a不变，并且pair不变且一直成对传递。
```Golang
//tty: pair<type:*os.File, value:"/dev/tty"文件描述符>
tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
//r: pair<type:  , value:>
var r io.Reader
//r: pair<type:*os.File, value:"/dev/tty"文件描述符>
r = tty
//w: pair<type:  , value:>
var w io.Writer
//w: pair<type:*os.File, value:"/dev/tty"文件描述符>
w = r.(io.Writer)
```
在下面这段代码中，Book重写了 ReadBook和WriteBook方法，所以b的type为Book。
```Golang
type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

//具体类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}
//b: pair<type:Book, value:book{}地址>
b := &Book{}
//r: pair<type:, value:>
var r Reader
//r: pair<type:Book, value:book{}地址>
r = b
r.ReadBook()
var w Writer
//r: pair<type:Book, value:book{}地址>
w = r.(Writer) //此处的断言为什么会成功？ 因为w r 具体的type是一致
w.WriteBook()
```