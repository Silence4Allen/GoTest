//程序所属包
package main

//导入依赖包
import (
	"fmt"
	"GoTest/src/util"
	"GoTest/src/new"
	"unsafe"
)

//常量定义
const Name string = "allen"

//全局变量的声明与赋值
var contents string = "It's me"

//一般类型声明
type intTest int

//结构体的声明
type structTest struct {
}

//声明接口
type interfaceTest interface {
}

//函数定义
func test() {
	fmt.Println(Name)
}

//主函数
func main() {
	var i int = 1;
	fmt.Println(unsafe.Sizeof(i))
	util.U1()
	new.U2()
	test()
	fmt.Println(contents)
}
