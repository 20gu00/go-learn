package main

import (
	"encoding/json"
	"fmt"
)

type s1 struct {
	// tag结构体的元信息
	// `k:"v" k2:"v2"`
	Name string `json:"name"` //序列化后数据是name,前端转递过来的数据也是name
	Age  int
	a    string //小写,不能跨包,不能被json包处理
}

func main() {
	var s = s1{
		Name: "q",
		Age:  100,
		a:    "z",
	}
	// []byte
	// 序列化
	strByte, _ := json.Marshal(s)
	fmt.Println(strByte)
	fmt.Println(string(strByte))

	// 反序列化
	str := `{
		Name:"q",
		Age:100,
		a:"z",
	}`
	var s2 = s1{}
	// 指针
	json.Unmarshal([]byte(str), &s2)
	fmt.Println(s2)

	// 空接口类型 变长参数
}
