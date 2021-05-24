package unitTestRel

import (
	"fmt"
	"testing"
)
/*
1 函数测试定义: func TestXxx(*testing.T)
2 这个 TestXxx 函数式放在一个文件尾部名 _test.go 中
*/
func Test01(t *testing.T) {
	fmt.Printf("testing ok")
	t.Log("err.......")
}
