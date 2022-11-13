package main

import (
	"fmt"
	"math"
)

func main() {
	n := 10
	//%b 2进制 %d 10进制 %o 八进制
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	m := 31
	//%x 16进制
	fmt.Printf("%x\n", m)
	var age uint8
	fmt.Println(age)
	//age = 256报错
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	//打印windows平台下的路径
	fmt.Println("c:\\:code\\go.exe")
	s1 := `这是多行
			字符串
			来表示的
\t根据这里的格式原样输出
		`
	fmt.Println(s1)
}
