package fmtPrt

import (
	"encoding/json"
	"fmt"
	"reflect"
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

	b := []string{"a", "b", "c"}
	fmt.Printf("b=%v	\n", b)
	fmt.Printf("b=%+v	\n", b)
	fmt.Printf("b=%#v	\n", b)

}

func Test2() {
	data := `{"data":{"a":1},"name":"张三","age":18,"high":true,"class":{"name":"1班","grade":3}}`
	str := []byte(data)
	request := make(map[string]interface{})
	err := json.Unmarshal(str, &request)
	if err == nil {
		if d, ok := request["data"]; ok {
			fmt.Printf("d=%v \n", d)
			switch d.(type) {
			case map[string]interface{}:
				d2 := d.(map[string]interface{})
				fmt.Printf("d2=%v \n", d2)
			default:
				fmt.Printf("err \n")
			}
		}

	}
}

func MyType1(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("default")
	}

}

func MyType2(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println(t)

}

func Test3() {
	var a = "1.5"
	MyType1(a)
	MyType2(a)

	var c interface{} = 10
	var b, ok = c.(int)
	if ok {
		fmt.Printf("b=%d", b)
	}
}
