package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

//流程控制语句：for、if、else、switch 和 defer

//processControl
func forFun() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//初始化语句和后置语句是可选的
	total := 1
	for total < 10 {
		total += total
	}
	fmt.Println(total)
}

//for 是 Go 中的 while
func whileFun() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	//if 的简短语句,该语句声明的变量作用域仅在 if 之内。
}

//if
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//switch
func switchFunc() {
	fmt.Print(`Go runs on `)

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf(os)
	}
}

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	fmt.Println(time.Saturday - today)

	//defer 语句会将函数推迟到外层函数返回之后执行
	defer fmt.Println("world")
	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
