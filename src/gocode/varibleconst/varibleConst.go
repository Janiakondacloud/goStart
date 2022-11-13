package main

import "fmt"

var name1 = "小王子"
var name2, age2 = "小王子", 13
var age1 = 13

const ddr = "this is config file"

func foo() (int, string) {
	return 100, "iloveyou"
}
func main() {
	//fmt.Println("hello")
	//var name string
	//var age int
	//var isOk bool
	//print(name, age, isOk)
	//var (
	//	a string
	//	b bool
	//	c int
	//	d float32
	//)
	//fmt.Println(a, b, c, d)
	//fmt.Println(ddr)
	//fmt.Println(name1, name2, age1, age2)
	////短变量声明一般用于函数中
	//m := 2000
	//fmt.Println(m)
	////哑元变量不会分配内存，所以不存在重复声明的情况
	//_, someString := foo()
	//someNum, _ := foo()
	//fmt.Println(someNum, someString)
	//const (
	//	config1 = "qweqwe"
	//	config2 = "qweqwe1"
	//	config3 = "qweqwe2"
	//)
	//fmt.Println(config1, config2, config3)
	//iota用于常量声音的时候来表示索引,每一行的iota都相同，_来跳过某个不想声明的常量，iota在const出现后就会被重置为0
	const (
		n1 = iota //0
		n2        // 1
		_         //2
		n4        //3
		n5 = 100  //3
	)
	//
	const n6 = iota
	const n7 = iota
	fmt.Println(n4, n5, n6, n7)
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Println(KB, MB, GB, TB, PB)
	const (
		a, b = iota + 1, iota + 2
		c, d
		e, f
	)
	fmt.Println(a, b, c, d, e, f)
}

//var const func等用于函数外的时候必须要以关键字来声明
//：短变量声明不能用于函数外
// _多用于站位表示忽略值
// 常量的定义的时候必须先赋值
