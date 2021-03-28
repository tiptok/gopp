package copy

import (
	"testing"
)

func TestSkipList(t *testing.T) {
	list := NewSkipList()
	list.AddNode(ValueInt(5))
	list.AddNode(ValueInt(1))
	list.AddNode(ValueInt(7))
	list.AddNode(ValueInt(2))
	list.AddNode(ValueInt(-2))
	list.AddNode(ValueInt(-0))
	list.AddNode(ValueInt(-10))

	printSkipList(list)
	//fmt.Println(list.HasNode(ValueInt(0)))
	//fmt.Println(list.HasNode(ValueInt(100)))
	list.DeleteNode(ValueInt(1))

	printSkipList(list)
}
