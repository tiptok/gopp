package containner

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_List(t *testing.T) {
	// 双向链表
	list := list.New()
	list.PushBack(9)
	list.PushBack(8)
	list.PushFront(1)
	fmt.Println(list.Front().Value)
	fmt.Println(list.Front().Next().Value)
}
