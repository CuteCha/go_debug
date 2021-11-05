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

type TestSlicePair struct {
	Id     string
	Weight int32
}

func Test03(t *testing.T) {
	result := make([]*TestSlicePair, 0)
	result = append(result, &TestSlicePair{Id: "a", Weight: 1}, &TestSlicePair{Id: "b", Weight: 5}, &TestSlicePair{Id: "c", Weight: 3}, &TestSlicePair{Id: "d", Weight: 2})
	fmt.Printf("result: \n")
	for _, item := range result {
		fmt.Printf("%v; ", *item)
	}
	fmt.Printf("\nsorted: \n")

	sort.Slice(result, func(i, j int) bool {
		return result[i].Weight > result[j].Weight
	})

	for _, item := range result {
		fmt.Printf("%v; ", *item)
	}
	fmt.Printf("\n")
}
