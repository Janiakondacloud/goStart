package main

import "fmt"

func b() {
	//处理错误panic，defer在最后执行
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b error")
		}
	}()
	panic("panic in b")
}
func a() {
	fmt.Println("func a")
}
func c() {
	fmt.Println("func c")
}
func main() {
	a()
	b()
	c()
}
