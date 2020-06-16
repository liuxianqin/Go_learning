package main

import "fmt"

// 上	 海	  自	来	水	来	自	海	上
// s[0]	s[1] s[2]					 s[len(s)-1]

func main() {
	var s string
	s = "上海自来水来自海上"
	r := make([]rune, 0, len(s))
	for _,char := range s{
		r = append(r, char)
	}
	fmt.Println(r)


	// for index, word := range s {
	// 	if s[index]-s[len(s)-index-1] == 0 {
	// 		// fmt.Println("不是回文", word)
	// 		fmt.Printf("%v %v %c 不是回文\n", s[index], s[len(s)-index-1], word)
	// 		return
	// 	}

	// }
	// fmt.Println("是回文")

	for i:=0;i<len(r)/2;i++{
		if r[i] != r[len(r)-i-1]{
			fmt.Println("不是回文")
			return 
		}
	}
	fmt.Println("是回文")
}
