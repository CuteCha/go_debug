package loopRel

import "fmt"

func ForTest1() {
	l := 9
	for {
		if l < 0 {
			break
		}
		fmt.Printf("l=%+v\n", l)
		l--
	}

}
