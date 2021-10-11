package sliceRel

import (
	"fmt"
	"sort"
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

func Test02(t *testing.T) {
	result := make([]int, 0)
	result = append(result, 1, 5, 3, 2)
	fmt.Printf("result: %v\n", result)

	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	fmt.Printf("result: %v\n", result)
}
