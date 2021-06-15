package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {

		fmt.Printf("2**%d = %d\n", i, v)
	}

	//可以将下标或值赋予 _ 来忽略它
	for i, _ := range pow {

		fmt.Printf("%d", i)
	}
}
