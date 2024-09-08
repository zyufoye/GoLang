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
## 0x02. reflect包
reflect包中有两个重要方法，一个是$ValueOf$，用来获取输入参数接口中的数据的值。还有一个是$TypeOf$用来动态获取输入参数接口中值的类型，如果接口为空则返回nil。 
最基本的反射用法： 
```Golang
func reflectNum(arg interface{}) {
	fmt.Println("type : ", reflect.TypeOf(arg))
	fmt.Println("value : ", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.2345

	reflectNum(num)
}
```
结构体反射高阶用法：
```Golang
type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called ..")
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{1, "Aceld", 18}

	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())

	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is:", inputValue)

	//通过type 获取里面的字段
	//1. 获取interface的reflect.Type，通过Type得到NumField ,进行遍历
	//2. 得到每个field，数据类型
	//3. 通过filed有一个Interface()方法等到 对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	//通过type 获取里面的方法,调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

}
```
## 0x03. 结构体标签
结构体标签，就是为当前结构体中的属性添加标签，可以为属性绑定一个或多个标签，利用Golang解析json文件时，结构体标签会发挥很大用途。结构体标签简单用法如下： 
```Golang
type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo, " doc: ", tagdoc)
	}
}

func main() {
	var re resume
	findTag(&re)
}
```
结构体标签在json中的应用，其中Marshal是将结构体编码为json格式的方法，以下代码展示了json编码解码的过程：
```Golang
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"xingye", "zhangbozhi"}}

	//编码的过程  结构体---> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	//解码的过程 jsonstr ---> 结构体
	//jsonStr = {"title":"喜剧之王","year":2000,"rmb":10,"actors":["xingye","zhangbozhi"]}
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error ", err)
		return
	}

	fmt.Printf("%v\n", myMovie)
}
```