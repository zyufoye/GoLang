## 0x00. type别名
为int声明一个别名。
```Golang
// 声明一种行的数据类型 myint， 是int的一个别名
type myint int
var a myint = 10
fmt.Println("a = ", a)
fmt.Printf("type of a = %T\n", a)
//type of a = main.myint
```
## 0x01. 如何定义一个结构体类型？以及值传递和指针传递的区别？
```Golang
// 定义一个结构体，此时Book是一个新的数据类型
type Book struct {
	title string
	auth  string
}
//值传递是不修改变量底层存储的内容的
func changeBook(book Book) {
	//传递一个book的副本
	book.auth = "666"
}

func changeBook2(book *Book) {
	//指针传递
	book.auth = "777"
}
changeBook(book1) //{Golang zhang3}
changeBook2(&book1) //{Golang 777}
```
## 0x02. Go中类的封装
在Go中，类的本质就是通过结构体绑定方法来定义实现的。
```Golang
//如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	//如果说类的属性首字母大写, 表示该属性是对外能够访问的，否则的话只能够类的内部访问
	Name  string
	Ad    int
	level int
}
func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(newName string) {
	//this 是调用该方法的对象的一个副本（拷贝）
	this.Name = newName
}
//创建一个对象
hero := Hero{Name: "zhang3", Ad: 100}
```
this Hero表示 Show（）这个方法是绑定到Hero上面的，谁调用这个方法，谁就是this（就是当前对象）。此时GetName、SetName、Show和结构体Hero共同组成了一个英雄类，前面三个是Hero类对应的方法。  
上面SetName方法是调用该对象的一个副本，只读方法是可以的，但是写方法也就是修改方法并不起作用，因为this调用的是该方法的对象的一个副本（拷贝）。
```Golang
func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	//this 是调用该方法的对象，修改该对象底层存储的值
	this.Name = newName
}
hero.SetName("li4")//此时调用就能修改hero中的name属性了
```
上面完成了Hero类的封装，即类定义和绑定在该类上的方法集合。
## 0x03. Go中类的继承
在下面示例中，Human是父类，SuperMan为子类，展示了子类如何继承父类，并且子类如何定义一个子类对象。
```Golang
//首先封装一个基础Human类
type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}
//再定义一个继承Human的SuperMan类
type SuperMan struct {
	Human //SuperMan类继承了Human类的方法
	level int
}

// 重定义父类的方法Eat()
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// 子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}
```
子类会继承父类的全部方法，同时子类可以重写父类方法，也可以定义自己的新方法。  
定义子类的两种方式：
```Golang
s := SuperMan{Human{"li4", "female"}, 88}
var s SuperMan
s.name = "li4"
s.sex = "male"
s.level = 88
```
## 0x03. Go中类的多态
在Go中利用interface接口来实现多态，接口定义一些抽象方法，子类再去实现（重写方法）。在具体的类中，不需要标明使用了哪个接口，只需要把接口里对应的方法全部实现即可。如果没有实现全部的方法，则这个接口的指针就无法指向该类。
```Golang
//本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的种类
}
//具体的类
type Cat struct {
	color string //猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}
//调用
var animal AnimalIF //接口的数据类型， 父类指针
animal = &Cat{"Green"}
animal.Sleep() //调用的就是Cat的Sleep()方法 , 多态的现象
//Cat is Sleep
animal = &Dog{"Yellow"}
animal.Sleep() // 调用Dog的Sleep方法，多态的现象
//Dog is Sleep
```
以上代码表示Cat类已经重写了AnimalIF方法，从而实现了AnimalIF的多态。调用过程中，把子类对象地址给到父类。多态调用的另一种写法：
```Golang
func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}
cat := Cat{"Green"}
dog := Dog{"Yellow"}
showAnimal(&cat)
showAnimal(&dog)
```
多态的必要条件是：有一个父类（有接口interface），有一个或多个子类（实现父类接口的全部方法），父类类型的变量（指针），指向或引用子类的具体数据变量。  
总结起来是：父类定义抽象方法，子类进行实现，把子类指针传给父类，父类调用。

## 0x04. Go中interface通用万能类型
interface{}这种写法表示是一个空接口，Go中包含了很多基本类型，int、string、float32、float64、struct等，这些基础数据类型都实现了interface{}的接口，可以用interface类型引用任意的数据类型。
```Golang
//interface{}是万能数据类型
func myFunc(arg interface{}) {}
```
此时arg就是一种万能数据类型，可以给他传入任何数据类型的数据。有一种办法，是专门为空接口定制的，就是断言机制，可以让万能数据类型arg向下判断到底属于那种类型。
```Golang
// interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//interface{} 改如何区分 此时引用的底层数据类型到底是什么？

	//给 interface{} 提供 “类型断言” 的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)

		fmt.Printf("value type is %T\n", value)
	}
}

type Bookmark struct {
	auth string
}
func main() {
	book := Bookmark{"Golang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
```