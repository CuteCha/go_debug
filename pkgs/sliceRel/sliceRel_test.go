package sliceRel

import (
	"testing"
	"time"
)

func Test01(t *testing.T) {
	result := make([]*string, 0)
	a0 := "aaa"
	result = append(result, &a0)
	t.Log(len(result))
	t.Log(*result[0])
	t.Log(time.Now())
	t.Log(time.Now().Unix())
	t.Log(time.Now().UnixNano())
}
