package channelRel

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

func getRoomTm(programId string) int64 {
	if programId == "" || strings.TrimSpace(programId) == "" {
		return 0
	}

	arr := strings.Split(programId, "_")
	if len(arr) != 3 {
		return 0
	}

	return ParseToLong(arr[2], 0)
}

//parse interface to int64 with dfault d
func ParseToLong(e interface{}, d int64) int64 {
	switch e := e.(type) {
	case float64:
		return int64(e)
	case float32:
		return int64(e)
	case int:
		return int64(e)
	case uint:
		return int64(e)
	case int8:
		return int64(e)
	case uint8:
		return int64(e)
	case int16:
		return int64(e)
	case uint16:
		return int64(e)
	case int32:
		return int64(e)
	case uint32:
		return int64(e)
	case int64:
		return e
	case uint64:
		return int64(e)
	case string:
		if i, ok := StrToLong(e); ok == nil {
			return i
		} else {
			return d
		}
	case []byte:
		if i, ok := StrToLong(string(e)); ok == nil {
			return i
		} else {
			return d
		}
	default:
		return d
	}
}

//parse string to int64 with error
func StrToLong(s string) (int64, error) {
	if strings.ContainsAny(s, ".") {
		if f, _e := strconv.ParseFloat(s, 64); _e == nil {
			return int64(f), nil
		} else {
			return math.MaxInt64, fmt.Errorf("date %+v parse to float64 err %+v", s, _e)
		}
	} else {
		if i, ok := strconv.ParseInt(s, 10, 64); ok == nil {
			return i, nil
		} else {
			return math.MaxInt64, fmt.Errorf("date %+v parse to int64 err %+v", s, ok)
		}
	}
}

func IsNeedBoostRoom(liveTmLower, roomTmLower int64) bool {

	//var viewers int32
	firstLiveTm := int64(1623571164)
	programId := "103632022_103632022_1623571164"
	roomStartTm := getRoomTm(programId)

	reqTm := int64(1623576164)
	fmt.Printf("roomStartTm(%v)-firstLiveTm(%v)= %v\n", roomStartTm, firstLiveTm, roomStartTm-firstLiveTm)
	fmt.Printf("reqTm(%v)-roomStartTm(%v)= %v\n", reqTm, roomStartTm, reqTm-roomStartTm)
	//if roomStartTm-firstLiveTm < liveTmLower || reqTm-roomStartTm < roomTmLower {
	//	return false
	//}
	if reqTm-roomStartTm < roomTmLower {
		return false
	}

	return true
}

func Test04(t *testing.T) {
	fmt.Printf("res: %v\n", IsNeedBoostRoom(int64(600), int64(600)))
}

func Test05(t *testing.T) {
	fmt.Printf("res: %v|%T\n", true, true)
}
