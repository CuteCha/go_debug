package pointerRel

import (
	"fmt"
	"strings"
	"testing"
)

type ManualChatroomRecTop struct {
	ChatroomId string
	StreamerId string
	Weight     int32
	StartTM    int64
	EndTM      int64
}

func Test01(t *testing.T) {
	ret := make([]*ManualChatroomRecTop, 0)
	top1 := &ManualChatroomRecTop{
		ChatroomId: "123_01",
		//StreamerId: "",
		Weight:     100,
		StartTM:    1,
		EndTM:      2,
	}
	fmt.Printf("top: %+v\n", top1)
	setStreamerId(top1)

	ret = append(ret, top1)

	fmt.Printf("top: %+v\n", top1)
	fmt.Printf("ret: %+v\n", ret[0])
}

func setStreamerId(topChatroom *ManualChatroomRecTop) {
	chatroomId := topChatroom.ChatroomId
	arr := strings.Split(chatroomId, "_")
	if len(arr) != 2 {
		fmt.Printf("TopChatroomId:%v not match!!!\n", chatroomId)
		topChatroom.StreamerId = ""
		return
	}

	topChatroom.StreamerId = arr[0]
}
