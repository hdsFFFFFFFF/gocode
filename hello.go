package main	//指明这个文件属于哪个包。package main表示一个可独立执行的程序。每个go应用程序都包含一个名为main的包

import "fmt"	//告诉go编译器这个程序需要使用fmt包

/* 全局变量的声明。注意！全局变量允许声明但不使用，局部变量不行 */
var x, y int	//同一类型的变量可以声明在同一行

var (
	aa int
	bb bool
)

var ff, gg int = 1, 2		//显示声明
var hh, kk = false, "hds"	//隐式声明
/*
	go的多行注释虽然和C的一样，但有一个地方不同，中间行不需要再用星号作连接
	hh, kk := false, "hds" 属于不带声明格式的，这种用法只能在函数体中出现，不能用于全局变量的声明与赋值
*/

/*
	var i1, i2 int
	var s11 string
	i1, i2, s11 = 1, 2, "hds"	//全局变量这么赋值会报错，只有在函数体中可以。
*/
//可以改为这样
var i1, i2, s11 = 1, 2, "hds"

/* func main是程序开始执行的函数，main()函数是每一个可执行程序所必须包含的 */
func main()  {
	fmt.Println(x, y, aa, bb, ff, gg, hh, kk)

	fmt.Println("hello world!" + "hello hds")	//不加\n默认会添加\n，如果加了\n就相当于\n\n

	var s string = "google" + "is" + "good"	//可以简写为s := "googleisgood"。go编译器可以根据值自动判断变量类型
	sss := "googleisgood"
	fmt.Println(s, sss)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d int
	fmt.Println(d)	//没有初始化的变量会默认初始化为0

	var e bool
	fmt.Println(e)	//bool的0值为false

	var i int
	var g float64
	var k string
	fmt.Printf("%v %v %q\n", i, g, k)

	/* 根据值自动判断变量类型 */
	var ss = "googlegood\n"
	var bb = true
	fmt.Println(ss, bb)

	/* 省略var */
	var dd int
	//dd := 1会导致编译中断，因为:=是声明语句
	dd,dd1 := 1, 2
	fmt.Println(dd, dd1)

	/* 多变量声明 */
	var vname1, vname2, vname3 int	//这里定义为int，所以三个变量类型应为一样的
	vname1, vname2, vname3 = 1, 2, 3

	var vname4, vname5, vname6 = 4, "kcm", 6	//不需要显示声明，自动推到。go编译器会根据值自动判断类型

	/* 下面这种格式只能在函数体中出现 */
	/*
		初始化声明
	*/
	vname7, vname8, vname9 := false, 8, "hds"	//:=左面出现的变量不能是已经声明过的
	fmt.Println(vname1, vname2, vname3, vname4, vname5, vname6, vname7, vname8, vname9)

	var iii int	//只有iii变量为int
	//iii := 1
	iii, iii1 := 1, "hds"	//因为iii变量已经被声明(定义就是给定了初值，已经被初始化了)为int，所以它的赋值只能是一个整型值
	fmt.Println(iii, iii1)

	/* go中的编译错误处理 */
	//s_s = 20	/* 报错 undefined: s_s */

	/* 声明了一个局部变量却没有在相同的代码块中使用它 */
	/*
		var s_s = "hello,world!"
		fmt.Println("hello")	//报错：declared but not used

		fmt.Println("hello", s_s)	//移除错误
	 */

	var i11, i22 int
	var s111 string
	i11, i22, s111 = 1, 2, "hds"

	fmt.Println(i11, i22, s111)
}
