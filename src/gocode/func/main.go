package main

import "fmt"

func intSum(a int, b int) int {
	return a + b
}
func intSum1(a int, b int) (ret int) {
	ret = a + b
	return
}
func intSum2(args ...int) (ret int) {
	for _, arg := range args {
		ret += arg
	}
	return
}
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}
func calc1(a, b int, op func(int, int) int) int {
	return op(a, b)
}
func a() func() {
	name := "沙河帅哥"
	return func() {
		fmt.Println("hello 沙河小王子" + name)
	}
}
func main() {
	fmt.Println(intSum(2, 5))
	fmt.Println(intSum1(3, 5))
	fmt.Println(intSum2(3, 5, 2, 1, 5))
	sum, sub := calc(100, 200)
	fmt.Println(sum, sub)
	fmt.Println(intSum1(intSum(1, 2), 3))
	fmt.Println(calc1(2, 5, intSum))
	r := a()
	//闭包=函数+外层变量的引用
	r()
}
