package main

import (
	"fmt"
	"sort"
)

func main() {
	//切片更灵活，不用指定长度
	//var a []string
	//var b []int
	//c := []bool{true, false}
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//var aArray = [5]string{"a23", "123123", "aasd", "123czx", "zxsea"}
	//a = aArray[1:4]
	//a1 := a[1:len(a)] //再切片
	//fmt.Println(a)
	//fmt.Println(a1)
	////make来创造切片
	//d := make([]string, 5, 10)
	//fmt.Println(d)
	//fmt.Printf("%T\n", d)
	//fmt.Println(len(d))
	//fmt.Println(cap(d))
	//扩容,切片动态调整长度，指针指向数组头，每次扩容指向新的地址
	//var intArray []int
	//for i := 0; i < 10; i++ {
	//	intArray = append(intArray, i)
	//	fmt.Printf("%v len:%d cap:%d ptr:%p\n", intArray, len(intArray), cap(intArray), intArray)
	//}
	//newArray := make([]int, 10)
	//newArray1 := intArray
	////copy复制切片的值，如果直接赋值的话引用一样，如果用copy复制了值但指针则是新的。
	//copy(newArray, intArray)
	//intArray[0] = 100
	//fmt.Println(newArray)
	//fmt.Println(intArray)
	//fmt.Println(newArray1)
	//fmt.Printf("%p\n", newArray)
	//fmt.Printf("%p\n", intArray)
	//fmt.Printf("%p\n", newArray1)
	////删除index=3的元素 公式:append(a[:index],a[index+1:]...)
	//newArray = append(newArray[:3], newArray[4:]...)
	//fmt.Println(newArray)
	strings := make([]string, 5, 10) //创建5个空格字符串容量为10
	for i := 0; i < 10; i++ {
		strings = append(strings, fmt.Sprintf("%v", i))
	}
	fmt.Println(strings)
	//排序数组
	ints := [13]int{1, 2, 5, 3, 7, 4, 8, 123, 22, 3332, 234, 643, 15}
	//对ints数组切片再指向ints
	sort.Ints(ints[:])
	fmt.Println(ints)
}
