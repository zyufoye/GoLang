package main //程序包名

/*
两种导包方式都可以，多个包建议用下面这种写法
import "fmt"
import "time"
*/
import (
	"fmt"
	"time"
)

// main函数
// golang要求左括号必须在func之后，他们在同一行，和函数名在同一行
func main() {
	//golang的表达式加不加";"都无所谓，推荐是不加
	fmt.Println("hello")

	time.Sleep(1 * time.Second)

}
