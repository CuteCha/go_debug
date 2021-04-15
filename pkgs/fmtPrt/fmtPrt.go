package fmtPrt

import "fmt"

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
