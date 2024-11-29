package time

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type student struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func TestTime(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	fmt.Println("------")
	now1 := now.Format("2006-01-02 15:04.000 PM")
	fmt.Println(now1)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	timeStr := "2024/11/07 14:00:00"
	timeObj, _ := time.Parse("2006/01/02 15:04:05", timeStr)
	fmt.Println(timeObj)
	timeObjLoc, _ := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	fmt.Println(timeObjLoc)
	reflect.TypeOf(now1)
	//v := reflect.ValueOf()
	//v.Elem()
	//v.IsValid()
	//stu1 := student{
	//	Name: "bob",
	//	Age:  6,
	//}
	//st := reflect.TypeOf(stu1)
	//for i := 0; i < st.NumField(); i++ {
	//	field := st.Field(i)
	//	field.Tag.Get("json")
	//}
	//v.NumField()
	//v.Field()
	//reflect.Value{}
}
