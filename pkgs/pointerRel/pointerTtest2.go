package pointerRel

import "fmt"

type StringArr []*string
type Context struct {
	Id     int
	Name   string
	Result *StringArr
}

func (context *Context) Build(c *Context) {
	result := make(StringArr, 0, 10)
	for i := 0; i < 10; i++ {
		t := fmt.Sprintf("abc_%d", i)
		result = append(result, &t)
	}

	c.Result = &result
}

func PointerTest22() {
	a := Context{
		Id:     123,
		Name:   "test",
		Result: nil,
	}

	b := Context{
		Id:     123,
		Name:   "test",
		Result: nil,
	}

	b.Build(&a)

	for i, e := range *a.Result {
		fmt.Printf("i=%d, e=%s\n", i, *e)
	}
}
