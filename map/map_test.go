package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	// 光声明map类型，但没有初始化，a就是nil
	var a map[string]int
	fmt.Println(a == nil)
	// map 初始化
	a = make(map[string]int, 8)
	// 初始化后a不是nil
	fmt.Println(a == nil)
	a["seal"] = 888
	a["ball"] = 666
	fmt.Printf("a:%#v\n", a)
	fmt.Printf("type:%T\n", a)

	value, ok := a["ball"]
	if ok {
		fmt.Println("hello ball~ value is ", value)
	} else {
		fmt.Println("no ball ball~")
	}

	delete(a, "ball")

	// map是无序的，与添加的顺序无关
	for k, v := range a {
		fmt.Println(k, v)
	}

	// 按固定顺序遍历map
	scoreMap := make(map[string]int, 100)

	// 添加50个键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.IntN(100) // 0 ~ 99
		scoreMap[key] = value
	}
	//for k, v := range scoreMap {
	//	fmt.Println(k, v)
	//}

	// 1. 按照key大小顺序遍历scoreMap
	keys := make([]string, 0, 100)
	for k := range scoreMap {
		keys = append(keys, k)
	}
	// 2. 对key做排序
	sort.Strings(keys)
	// 3. 按照排序后的key对scoreMap排序
	for _, key := range keys {
		fmt.Println(key)
	}

	s := "how do you do"
	var wordCount = make(map[string]int, 10)
	// 1. 统计有哪些单词
	words := strings.Split(s, " ")
	// 2. 遍历单词统计
	for _, word := range words {
		v, ok := wordCount[word]
		if ok {
			wordCount[word] = v + 1
		} else {
			wordCount[word] = 1
		}
	}
	for k, v := range wordCount {
		fmt.Println(k, v)
	}

}
