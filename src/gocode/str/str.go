package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(len("qweqeqwe"))
	s1 := "what are you doing"
	split := strings.Split(s1, " ")
	fmt.Printf("%T\n", split) //切片
	fmt.Println(strings.HasSuffix(s1, "doing"))
	fmt.Println(strings.HasPrefix(s1, "doing"))
	fmt.Println(strings.Contains(s1, "are"))
	fmt.Println(strings.Index(s1, "you"))
	//字符用单引号，字符串用双引号
	//byte uint8的别名 ascii码
	//rune int32别名 unicode编码
	var c1 byte = 'c'
	var c2 rune = 'C'
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Printf("c1:%T,c2:%T", c1, c2)
	s2 := "hello刘雨嘉"
	//输入单一的字符
	for _, r := range s2 {
		fmt.Printf("%c\n", r)
	}
	//这样的方式中文会乱码但是在ascii码范围内就没有问题
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%c,\n", s2[i])
	}
}
