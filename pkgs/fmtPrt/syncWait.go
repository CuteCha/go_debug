package fmtPrt

import (
	"fmt"
	"sync"
	"time"
)

var n = 10

func SWTest1() {
	for i := 0; i < n; i++ {
		go fmt.Println(i)
	}
	time.Sleep(time.Second)
}

func SWTest2() {
	c := make(chan bool, n)
	for i := 0; i < n; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}

	for i := 0; i < n; i++ {
		<-c
	}
}

func SWTest3() {
	fmt.Println("start....")
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("end....")
}
