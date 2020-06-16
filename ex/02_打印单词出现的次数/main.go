package main

import (
	"fmt"
	"strings"
)

func main() {
	st := "How do you do"
	st_space := strings.Split(st, " ")
	fmt.Println(st_space)

	m1 := make(map[string]int,  10)

	for index, word := range st_space {
		fmt.Println(index, word)

		var small string
		small = strings.ToLower(word)
		m1[small] = 0
	}

	fmt.Println(m1)
	
	for _, word := range st_space {
		m1[strings.ToLower(word)] += 1
	}

	fmt.Println(m1)
}
