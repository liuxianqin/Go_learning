package main

import "fmt"

func main() {
	// 位运算

	// 按位 与 或 抑或
	//    101 5
	//    010 2
	// &  000 与=>0
	// |  111 或=>7
	// ^  111 不一样就是1=>7
	fmt.Println(5 & 2)
	fmt.Println(5 | 2)
	fmt.Println(5 ^ 2)

	// 移位
	fmt.Println(5 << 1)  // 101 -> 1010 (10)
	fmt.Println(1 << 10) // 1 -> 10000000000 (1024)

	fmt.Println(1024 >> 10)
	fmt.Println(5 >> 1)   // 101 -> 10 (2)

	fmt.Println()
	// 超出范围的移位
	var m = int8(1)
	fmt.Println(m << 9)


}
