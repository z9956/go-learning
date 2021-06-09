package main

import (
	"fmt"
)

/**
* 函数
 */
func add(x int, y int) int {
	return x + y
}

//当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略
func shortAdd(x, y int) int {
	return x + y
}

//多值返回
func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	fmt.Println(shortAdd(1, 2))
}
