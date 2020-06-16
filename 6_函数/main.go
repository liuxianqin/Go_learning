package main

import "fmt"

// 命名返回值，相当于提前声明ret
func sum(x int, y int) (ret int) {
	return x + y
}

func sum2(x int, y int) (ret int) {
	ret = x + y
	return
}

// 多个返回值
func sum3() (int, int) {
	return 5, 5
}

// 参数的类型简写
func sum4(x, y int) int {
	return x + y
}

// 可变长参数  y 可以不传 可以传一个 可以传多个 可变长参数必须放最后
func sum5(x int , y ...int)int{
	ys := x
	for _,value := range y{
		ys += value
	}
	return ys
}

func main() {
	fmt.Println()
	r := sum(1, 2)
	fmt.Println(r)

	r = sum2(3, 4)
	fmt.Println(r)

	m, n := sum3()
	fmt.Println(m, n)

	r = sum4(8,9)
	fmt.Println(r)

	r = sum5(1,2,3,4,5)
	fmt.Println(r)
}
