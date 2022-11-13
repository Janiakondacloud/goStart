package main

import (
	"fmt"
	"strings"
)

func main() {
	//与slice一样都必须要先申请内存才能操作
	//m := make(map[string]int, 8)
	//fmt.Println(m == nil)
	//m["分数"] = 100
	//m["年龄"] = 10
	//fmt.Println(m)
	//b := map[int]bool{
	//	1: true,
	//	2: false,
	//}
	//fmt.Println(b)
	////判断其值是否存在该map中
	//value, ok := m["分数"]
	//if ok {
	//	fmt.Println(value)
	//} else {
	//	fmt.Println("并没有该数据在m中")
	//}
	////遍历
	//for k, v := range m {
	//	fmt.Println("k:"+k, "value:", v)
	//}
	////只遍历key
	//for k := range m {
	//	fmt.Println(k)
	//}
	////只遍历value
	//for _, value = range m {
	//	fmt.Println(value)
	//}
	////删除key
	//delete(m, "年龄")
	//fmt.Println(m)
	//按照某个顺序遍历map
	//scoreMap := make(map[string]int, 100)
	//for i := 0; i < 50; i++ {
	//	key := fmt.Sprintf("stu%02d", i)
	//	value := rand.Intn(100) //0-99随机整数
	//	scoreMap[key] = value
	//}
	////for k, v := range scoreMap {
	////	fmt.Println(k, v)
	////}
	////fmt.Println("--------------")
	////按照key从小到大去遍历
	//keys := make([]string, 0, 100)
	//for k := range scoreMap {
	//	keys = append(keys, k)
	//}
	//sort.Strings(keys) //keys排序
	//for _, key := range keys {
	//	fmt.Println(key, scoreMap[key])
	//}
	//map类型的切片---
	//mapSlice := make([]map[string]int, 8, 8) //只完成切片的初始化，但是内部的map没有初始化所以为nil
	//fmt.Println(mapSlice[0] == nil)
	////还需要对内部map进行初始化
	//mapSlice[0] = make(map[string]int, 8)
	//mapSlice[0]["电子科技大学"] = 100
	//fmt.Println(mapSlice)
	//value为切片的map---
	//sliceMap := make(map[string][]int, 8)
	////fmt.Println(sliceMap)
	//_, ok := sliceMap["中国"]
	//if ok {
	//	fmt.Println(sliceMap["中国"])
	//} else {
	//	sliceMap["中国"] = make([]int, 5) //创建key为中国的slice值
	//	sliceMap["中国"][0] = 100
	//	sliceMap["中国"][1] = 200
	//	sliceMap["中国"][2] = 300
	//	fmt.Println(sliceMap["中国"])
	//}
	//统计字符串中每个单词出现的次数
	var s = "what are you doing now i am doing this code for dream"
	wordCount := make(map[string]int, 10)
	wordSlice := strings.Split(s, " ")
	for _, word := range wordSlice {
		_, ok := wordCount[word]
		if ok {
			wordCount[word]++
		} else {
			wordCount[word] = 1
		}
	}
	fmt.Println(wordCount)
}
