package funcRel

import "fmt"

func FuncTest1() {
	//multiParam 可以接受可变数量的参数
	multiParam("jerry1", "herry1")
	names := []string{"jerry", "herry"}
	multiParam(names...)
	multiParam2("jerry1", 1, 2)
}
func multiParam(args ...string) {
	//接受的参数放在args数组中
	for _, e := range args {
		fmt.Println(e)
	}
}

func multiParam2(name string, args ...int) {
	fmt.Println(name)
	//接受的参数放在args数组中
	for _, e := range args {
		fmt.Println(e)
	}
}

type MYFUNC func(int, int) (int, error)

func addFunc(a int, b int) (int, error) {
	var c = a + b

	return c, nil
}

func FuncTest2() {
	var f MYFUNC
	f = addFunc
	c, _ := f(2, 1)
	fmt.Printf("c=%d", c)
}
