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

/*
零值 (没有明确初始值的变量声明会被赋予它们的 零值)
int     0
int8    0
int32   0
int64   0
uint    0x0
rune    0 //rune的实际类型是 int32
byte    0x0 // byte的实际类型是 uint8
float32 0 //长度为 4 byte
float64 0 //长度为 8 byte
bool    false
string  ""
*/
var i int //0

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
