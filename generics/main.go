package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

// 第一种写法
//func main() {
//	bigger := getBigger[int32](3, 6)
//	println(bigger)
//
//}
//
//func getBigger[T int32 | int64](a, b T) T {
//	if a > b {
//		return a
//	}
//	return b
//}

// 第二种写法
//type Comparable interface {
//	int32 | int64
//}
//
//func main() {
//	getBigger[rune](3, 6)
//}
//
//func getBigger[T Comparable](a, b T) T {
//	if a > b {
//		return a
//	}
//	return b
//}

// 第三种
//type Age int32
//type Comparable interface {
//	~int32 | int64
//}
//
//func main() {
//	//getBigger[rune](3, 6)
//	bigger := getBigger[Age](3, 6)
//	fmt.Println(bigger)
//}
//
//func getBigger[T cmp.Ordered](a, b T) T {
//	if a > b {
//		return a
//	}
//	return b
//}

// ...结构体泛型等 相对于 interface
type GetUserRequest struct{}
type GetBookRequest struct{}

func httpPRC(request any) {
	url := "http://127.0.0.1"
	tp := reflect.TypeOf(request)
	switch tp.Name() {
	case "GetUserRequest":
		url += "user"
	case "GetBookRequest":
		url += "book"
	default:
		panic("unsupported type")
	}
	fmt.Println(url)
	bs, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Json Marshal err")
		return
	}
	http.Post(url, "text/plain", bytes.NewReader(bs))
}

func main() {
	httpPRC(GetBookRequest{})
	httpPRC(GetUserRequest{})
}
