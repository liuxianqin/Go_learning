package main

import (
	"fmt"
	"unicode"
)

func main() {
	s1 := "hello世界12단계"
	var count int
	for index, char := range s1 {
		// fmt.Println(index, char)
		fmt.Printf("%d %c\n", index, char)
		if unicode.Is(unicode.Han, char) {
			count++
		}
		// if unicode.Is(unicode.K)
	}
	fmt.Println("汉字的数量", count)
}
