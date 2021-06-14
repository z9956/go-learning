package main

import "fmt"

func printSlice(list []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(list), cap(list), list)
}

func main() {
	//数组
	var a [2]string
	a[0] = "Hello"

	primes := [6]int{2, 3, 5, 7, 11, 13}

	/*
		切片
		它会选择一个半开区间，包括第一个元素，但排除最后一个元素
		更改切片的元素会修改其底层数组中对应的元素
	*/
	var list []int = primes[1:4]
	list[1] = 9

	//切片文法
	results := []struct {
		i int
		b bool
	}{
		{2, true},
	}
	fmt.Println(results)

	//切片的长度与容量
	originList := []int{1, 2, 3, 4, 5, 6}

	printSlice(originList)

	//截取切片使其长度为 0
	originList = originList[:0]
	printSlice(originList)

	//拓展其长度
	originList = originList[:4]
	printSlice(originList)

	//舍弃前两个值
	originList = originList[2:]
	printSlice(originList)

	//切片的零值是 nil
	var s []int
	fmt.Println(s)
	if s == nil {
		fmt.Println("nil!")
	}

	//make 创建切片
	makeList := make([]int, 2)
	printSlice(makeList)

	//向切片追加元素
	var array []int

	//add
	array = append(array, 1)
	printSlice(array)

	//一次性添加多个元素
	array = append(array, 2, 3, 4)
	printSlice(array)
}
