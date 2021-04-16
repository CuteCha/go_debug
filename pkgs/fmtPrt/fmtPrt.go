package fmtPrt

import (
	"encoding/json"
	"fmt"
)

type student struct {
	id   int32
	name string
}

func Test1() {
	a := &student{id: 1, name: "xiaoming"}

	fmt.Printf("a=%v	\n", a)
	fmt.Printf("a=%+v	\n", a)
	fmt.Printf("a=%#v	\n", a)

	b:=[]string{"a","b","c"}
	fmt.Printf("b=%v	\n", b)
	fmt.Printf("b=%+v	\n", b)
	fmt.Printf("b=%#v	\n", b)

}

func Test2()  {
	data := `{"data":{"a":1},"name":"张三","age":18,"high":true,"class":{"name":"1班","grade":3}}`
	str := []byte(data)
	request := make(map[string]interface{})
	err := json.Unmarshal(str, &request)
	if err == nil {
		if d, ok := request["data"]; ok {
			switch d.(type) {
			case map[string]interface{}:
				d2 := d.(map[string]interface{})
				println(d2)
			default:
				println("err")
			}
		}

	}
}