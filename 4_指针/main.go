package main

import "fmt"

func main() {
	fmt.Println("指针")

	n := 2
	p := &n
	// fmt.Println(p)
	fmt.Printf("%T，%p\n", p, p)
	m := *p
	fmt.Printf("%T %v\n", m, m)

	// new()分配内存 很少用，给基本数据类型分配，返回类型的指针*int *string
	var a = new(int)
	fmt.Println(a, *a)
	*a = 100
	fmt.Println(a, *a)

	// make()分配内存  返回的是slice,map,chan类型本身
	// ...




}
