package main

import "fmt"

func main() {
	type Vertex struct {
		X int
		Y int
	}

	//结构体文法
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
		v3 = Vertex{}      // X:0 Y:0
		v4 = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	)

	v := Vertex{2, 3}
	v.X = 4
	fmt.Println(v.X, v)

	fmt.Println(v1, v2, v3, v4)

	result := &v
	result.Y = 12
	fmt.Println(result, v)
}
