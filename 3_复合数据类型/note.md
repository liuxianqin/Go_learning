
### 数组

数组是值类型 复制了一份，新的和旧的没有联系了。

```go

    // 数组长度是数组类型的一部分
	var a1 [3]bool
	var a2 [4]int

	fmt.Printf("%T", a1)
    fmt.Printf("%T", a2)

    // 根据初始化自动推断长度
	a3 := [...]int{1,2,3,4,5,10:10} 
	fmt.Println(a3)
	fmt.Printf("%T\n", a3)

	//遍历 for i
	citys := [...]string{"北京", "shanghai", "深圳"}
	for i:=0;i<len(citys);i++{
		fmt.Println(citys[i])
	}

	//遍历 for range
	for i,v := range citys{
		fmt.Println(i, v)
	}

	// 多维数组
	var a32 [3][2]int
	a32 = [3][2]int{
		[2]int{1,2},
		[2]int{4,5},
		[2]int{9,10},
	}

	for _,v := range a32{
		for _, v2 := range v{
			fmt.Println(v2)
		}
	}

```


### 切片 Slice

支持扩容， 引用类型

```go
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
```

切片限定了一块连续的内存。

切片不能直接比较。
切片能和nil比较。

```go
	var n1 []int 			// 等于nil
	n2 := []int{}			// 已经有地址了
	n3 := make([]int, 0)	// 已经有地址了
	fmt.Println(n1==nil) //true
	fmt.Println(n2==nil) //false
	fmt.Println(n3==nil) //false
```

切片是空的判断： `len(slicename) == 0`

#### 切片的复制

引用

#### append 方法 追加元素

```go
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
```

#### copy

需要提前算好长度
```go
	// copy
	co1 := []int{1, 2, 3}
	co2 := co1 // 赋值
	var co3 = make([]int, 3)
	copy(co3, co1)
	fmt.Println(co1, co2, co3)
```

#### 删除元素

利用切片的特性删除元素

