package main

import (
	"fmt"
)

/*
* 基本类型
*  bool
*  int(32 or 64), int8, int16, int32(rune), int64
*  uint(32 or 64), uint8(byte), uint16, uint32, uint64
*  float32, float64
*  string
*  complex64, complex128
* array    -- 固定长度的数组
 */
var isActive bool
var hello string = `hello` //反引号（` `） or 双引号（""）

//函数
func add(x int, y int) int {
	return x + y
}

/*
* 简短声明 :=
* 只能用在函数内部
 */
func test() {
	var i = 1
	j := 2
	fmt.Println(i, j)
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
	fmt.Println(swap("1", "3"))
}
