package main

import "fmt"

func main() {
	//位运算符
	a := 1               //001
	b := 5               //101
	fmt.Println(a & b)   //与运算符 两位都为1才为1 001
	fmt.Println(a | b)   //或运算符 两位都为0才为0 101
	fmt.Println(a ^ b)   //异或 相同为0 相异为1 100
	fmt.Println(a << 2)  //100
	fmt.Println(b << 2)  //10100
	fmt.Println(a >> 2)  //001=>0000,1 0
	fmt.Println(b >> 2)  //001,01 1
	fmt.Println(1 << 10) //10000000000

}
