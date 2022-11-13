package main

import "fmt"

func main() {
	i := 10
	for i > 0 {
		fmt.Println(i)
		i--
	}
	//无线循环
	//for {
	//	fmt.Println("无限循环中....")
	//}
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Print("----")
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}
