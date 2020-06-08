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
	ss1 := fmt.Sprint("ff", "haha", "HEHE")
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
	s3[0] = '红' // 切片里面保存的是字符
	fmt.Println(string(s3))

	// 双引号 和 单引号
	c1 := "红" // string
	c2 := '红' // int32
	fmt.Printf("c1:%T, c2:%T", c1, c2)

	fmt.Println()

	if age > 18 {
		fmt.Println("赛马")
	} else if age > 19 {
		fmt.Println("彩票")
	} else {
		fmt.Println("储蓄")
	}

	// for 循环 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 变种1 前面已经定义了i
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// 变种2 前面已经定义了i
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 无限循环
	// for {
	// 	fmt.Println("123456789")
	// }

	// for range 循环
	s := "合理的string"
	for index, v := range s {
		// fmt.Println(index, v)
		fmt.Printf("%d, %c\n", index, v) // 从中可以看出汉字一般占据3个字节
	}

	str := "只获取规划的string"

	for _, v := range str {
		// fmt.Println(v)
		fmt.Printf("%c %T  ", v, v)
	}
	fmt.Println()
	fmt.Println(str[1])

	// 跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println()

	//跳过某次for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()

	var finger = 2
	switch finger {
	case 1:
		fmt.Println("大木自己")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	}

	// 作用域里面的
	switch finger := 3; finger {
	case 3:
		fmt.Println("小木制")
	case 2:
		fmt.Println("无名指")
	default:
		fmt.Println("手指头")
	}

}
