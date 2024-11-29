package main

import "fmt"

func main() {
	//支持的写法
	a := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	//xx不支持xx多维数组的内层使用...
	//b := [3][...]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}

	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}

	c := [3]int{10, 20, 30}
	modifyArray(c) //在modify中修改的是a的副本x
	fmt.Println(c) //[10 20 30]
	d := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(d) //在modify中修改的是b的副本x
	fmt.Println(d)  //[[1 1] [1 1] [1 1]]
}

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
