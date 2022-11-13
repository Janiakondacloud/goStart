package main

import "fmt"

func main() {
	//var a [3]int
	//var b [4]string
	//var c [5]bool
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	////定义并且初始化
	//cityArray := [4]string{"北京", "上海", "成都", "广州"}
	//fmt.Println(cityArray)
	////不确定长度
	//var someArray = [...]any{"北京", 1, false, true, 123.22}
	//fmt.Println(someArray)
	////使用索引值
	//var langArray = [...]string{1: "java", 2: "python", 7: "js"}
	//fmt.Println(langArray[3])     //打印不出来 只能 1 2 7
	//fmt.Printf("%T\n", langArray) //长度为8
	//for index, value := range langArray {
	//	fmt.Print(index)
	//	fmt.Println(value)
	//}
	//二维数组,只能外部用...
	cityArray1 := [...][2]string{
		{"bj", "上海"},
		{"bj1", "上海1"},
		{"bj2", "上海2"},
		{"bj3", "上海3"},
	}
	fmt.Println(cityArray1)
	fmt.Println(cityArray1[2][0])
	fmt.Println("----")
	for _, value := range cityArray1 {
		for _, value1 := range value {
			fmt.Println(value1)
		}
	}
	//practice 找出数组中和为10的两个元素下标所有的结果,(1,2)和(2,1)算一种
	intArray := [...]int{1, 3, 5, 7, 8, 9} //(0,5) (1,3)
	for index, value := range intArray {
		for i := index + 1; i < len(intArray); i++ {
			if value+intArray[i] == 10 {
				fmt.Printf("(%v,%v)\n", index, i)
				break
			}
		}
	}
}
