package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

//映射的文法
//映射的文法与结构体相似，不过必须有键名
var result = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
}

//若顶级类型只是一个类型名，你可以在文法的元素中省略它
var short = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
}

func main() {
	//映射将键映射到值。
	//映射的零值为 nil 。nil 映射既没有键，也不能添加键
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(result, short)

	n := make(map[string]int)

	n["key"] = 42
	fmt.Println(n)

	n["age"] = 40
	fmt.Println(n)

	//删除元素
	delete(n, "age")
	fmt.Println(n)

	//通过双赋值检测某个键是否存在
	//若 key 在 m 中，ok 为 true ；否则，ok 为 false。
	//若 key 不在映射中，那么 elem 是该映射元素类型的零值
	v, ok := n["key"]
	fmt.Println(v, ok)
}
