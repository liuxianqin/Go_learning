package main

import "fmt"

// 数组长度是数组类型的一部分

func main() {
	fmt.Println("hah")

	var a1 [3]bool
	var a2 [4]int

	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a2)

	fmt.Println(a1)
	fmt.Println(a2)

	a1 = [3]bool{true, true, false}

	// 根据初始化自动推断长度
	a3 := [...]int{1, 2, 3, 4, 5, 10: 10}
	fmt.Println(a3)
	fmt.Printf("%T\n", a3)

	//遍历 for i
	citys := [...]string{"北京", "shanghai", "深圳"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	//遍历 for range
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 多维数组
	var a32 [3][2]int
	a32 = [3][2]int{
		[2]int{1, 2},
		[2]int{4, 5},
		[2]int{9, 10},
	}

	for _, v := range a32 {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	fmt.Println()

	// 求[1 3 5 7 8]中sum==8的两个下标
	list := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i]+list[j] == 8 {
				fmt.Println(i, j, list[i], list[j])
			}
		}
	}
	fmt.Println()

	// slice
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil, s2 == nil)

	// 切片的初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"BJ", "SZ", "WX"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil, s2 == nil)

	// 长度 len()  容量 cap()

	// 从数组得到切片
	fa := [...]int{1, 2, 3, 4, 5, 6, 7}
	su := fa[0:4] // 左闭区间 右开区间
	// 切片的容量是底层数组从切片的第一个元素到数组最后一个元素的容量
	fmt.Printf("len(su):%d, cap(su):%d\n", len(su), cap(su))
	// su = fa[:] 	 	// 全部
	su = fa[3:]
	fmt.Printf("len(su):%d, cap(su):%d\n", len(su), cap(su))
	fmt.Println(su)

	// 切片的切片
	su_su := su[1:]
	fmt.Println(su_su)

	fmt.Println()
	// make函数创造切片
	sli1 := make([]int, 5, 10)
	// fmt.Println(sli1)
	fmt.Printf("sli1 = %v, size:%d, capacity:%d\n", sli1, len(sli1), cap(sli1))

	var n1 []int
	n2 := []int{}
	n3 := make([]int, 0)
	fmt.Println(n1 == nil)
	fmt.Println(n2 == nil)
	fmt.Println(n3 == nil)

	fmt.Println()
	// append
	locations := []string{"北京", "上海", "广州"}
	// locations[3] = "广州"   // out of range 编译错误
	fmt.Println(locations)

	// 用原来的变量 接受 返回值
	locations = append(locations, "迪拜", "成都")
	fmt.Println(locations)

	// 把一个切片追加到另一个切片后面 ...
	new_locations := []string{"无锡", "苏州"}
	locations = append(locations, new_locations...)

	fmt.Println(locations)

	// copy
	co1 := []int{1, 2, 3}
	co2 := co1 // 赋值
	var co3 = make([]int, 3)
	copy(co3, co1)
	fmt.Println(co1, co2, co3)

	// 删除 index是2的元素     切片不存数
	drop := []int{0, 1, 2, 3, 4}
	drop_slice := drop[:]
	drop = append(drop[:2], drop[3:]...)   // 修改了底层数组
	fmt.Println(drop)
	fmt.Println(drop_slice)  // 数字被替换

}
