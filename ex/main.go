package main

import "fmt"

// import "fmt"

func main() {
	fmt.Println("打印99乘法表")

	// 1*1=1
	// 1*2=2 2*2=4
	// 1*3=3 2*3=6 3*3=9

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()

	}
}
