package main

// import "fmt"
import (
	"fmt"
	"math"
	"strings"
)

var dadada int

const pi = 3.1415926

const (
	n1 = 100 // 0
	n2
	n3
)

const (
	a1 = iota //0
	a2        //1
	a3
)

const (
	s1 = iota
	s2
	_
	s3
)

func main() {
	fmt.Println("Hello World!")

	var (
		name string
		age  int
		isOK bool
	)

	name = "Zhang"
	age = 16
	isOK = true
	fmt.Printf("name:%s", name)
	fmt.Print(age)
	fmt.Println(isOK)
	fmt.Println(pi)

	fmt.Println(n1, n2, n3)
	fmt.Println(a1, a2, a3)
	fmt.Println(s1, s2, s3)

	const (
		x1 = iota // 0
		x2 = 100  // 100
		x3        // 100
		x4 = iota // 3
		x5        // 4
	)
	fmt.Println(x1, x2, x3, x4, x5)

	var i1 = 101
	fmt.Println(i1)

	fmt.Println(math.MaxFloat32)

	f1 := 1.2234
	fmt.Printf("%T\n", f1)

	b1 := true
	var b2 bool

	fmt.Println(b1, b2)

	s1 := `
一个人
要像一支队伍
	`
	fmt.Println(s1)
	fmt.Println(len(s1))
	ss1 := fmt.Sprint("%s%sff", "haha", "HEHE")
	fmt.Println(ss1)

	pwd := "ABC/home/lxq/Go/src/hello哈哈哈3/5000 방법"
	retPwd := strings.Split(pwd, "/")
	fmt.Println(retPwd)

	// 前缀
	fmt.Println(strings.HasPrefix(pwd, "/home"))
	// 后缀
	fmt.Println(strings.HasSuffix(pwd, "llo"))
	// 出现的位置 不出现：-1
	fmt.Println(strings.Index(pwd, "z"))

	// 拿出具体的字符
	for _, c := range pwd {
		// fmt.Println(c) // 字符的rune类型 一个utf-8类型字符
		fmt.Printf("上一个%c\n", c)
	}

	// 字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红'   // 切片里面保存的是字符
	fmt.Println(string(s3))

	// 双引号 和 单引号
	c1 := "红"   // string
	c2 := '红'	// int32
	fmt.Printf("c1:%T, c2:%T", c1, c2)

}
