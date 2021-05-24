package trickerRel

import (
	"fmt"
	"time"
)

func tickerOn(name string, f func()) {
	f()
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for range ticker.C {
			f()
		}
	}()
	fmt.Printf("Initial %s ticker: %+v\n", name, ticker)
}

func printCurrentTs() {
	now := time.Now()
	fmt.Printf("printCurrentTs: %+v\n", now)
}

func Test01() {
	tickerOn("T1", printCurrentTs)
	time.Sleep(3 * time.Second)
}

func Test02() {
	t := time.NewTicker(time.Second * 2)
	defer t.Stop()
	for {
		<-t.C
		fmt.Printf("Ticker running...%+v\n", time.Now())
	}
}

func tickerOn02(f func()) {
	t := time.NewTicker(time.Second * 2)
	defer t.Stop()
	for {
		<-t.C
		f()
	}
}

func Test03() {
	tickerOn02(printCurrentTs)
}
