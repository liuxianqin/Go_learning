### map

```go

    // map
	var m1 map[string]int

	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)

	m1["年龄"] = 3
	m1["叶子"] = 2
	fmt.Println(m1)
	fmt.Println(m1["年龄"])

	value, ok := m1["年龄"]
	fmt.Println(value, ok)
	if !ok {
		fmt.Println("没有这个key")
	} else {
		fmt.Println(value)
	}

```

#### delete value
```go
    // delete value
	delete(m1, "年龄")
	m1["颜色"] = 0
	fmt.Println(m1)

```


#### 顺序输出map

```go


    rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("student%02d", i)
		value := rand.Intn(100) // 0-99
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)

	// 按照人名顺序 把所有的key放到一个切片里面
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
```

#### 元素是 map 的 slice

```go
// 元素类型是map的slice
	var s1 = make([]map[int]string, 10, 10)
	// size 0 : index out of range
	// size 1 : assignment to entry in nil map

	// s1[0][100] = "FF"
	fmt.Println(s1)

	s1[0] = make(map[int]string, 1)
	s1[0][100] = "HH"
	fmt.Println(s1)
```

#### 元素是 slice 的 map

```go
	// 值为切片 的 map
	var m2 = make(map[string][]int, 10)
	m2["BEIJING"] = []int{10,20,30,}
	fmt.Println(m2)
	m2["SHANGHAI"] = []int{20,30,40}
    fmt.Println(m2)
```

