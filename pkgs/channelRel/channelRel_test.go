package channelRel

import (
	"fmt"
	"testing"
	"time"
)

func Test01(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	x := <-ch
	fmt.Printf("x=%d\n", x)
}

func timer(d time.Duration) <-chan int {
	c := make(chan int)
	go func() {
		time.Sleep(d)
		c <- 1
	}()

	return c
}

func Test02(t *testing.T) {
	for i := 0; i < 24; i++ {
		c := timer(1 * time.Second)
		fmt.Printf("x=%d\n", <-c)
	}

}

func player(table chan int) {
	for {
		b := <-table
		fmt.Printf("b=%d; ", b)
		b++
		fmt.Printf("b=%d\n", b)
		time.Sleep(100 * time.Millisecond)
		table <- b
	}
}

func Test03(t *testing.T) {
	var b int
	table := make(chan int)
	for i := 0; i < 24; i++ {
		go player(table)
	}

	table <- b
	time.Sleep(1 * time.Second)
	<-table
}

