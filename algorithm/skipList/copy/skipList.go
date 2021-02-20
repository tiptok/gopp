package copy

/*
	跳表实现
*/
import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const MaxLevel = 10
const Probability = 0.5 // 基于时间与空间综合 best practice 值, 越上层概率越小

type (
	SkipList struct {
		maxLevel    int
		Level       int
		HeadNodeArr []*Node
	}
	Node struct {
		Value Value
		Prev  *Node //同层 前节点
		Next  *Node //同层 后节点
		Down  *Node //下层 同节点
	}
	Value interface {
		Less(value Value) bool
		String() string
	}
	ValueInt int
)

var _ Value = (*ValueInt)(nil)

func (v ValueInt) Less(value Value) bool {
	valueInt := value.(ValueInt)
	return v < valueInt
}
func (v ValueInt) String() string {
	return strconv.Itoa(int(v))
}

func (list *SkipList) AddNode(value Value) {
	if list.HasNode(value) != nil {
		// 如果包含相同的数据，就返回，不用添加了
		return
	}
	headNodeInsertPositionArr := make([]*Node, MaxLevel)
	// 只有层级在大于等于 0 的时候在进行循环判断，如果层级小于 0 说明是没有任何数据
	if list.Level >= 0 {
		level := list.Level
		node := list.HeadNodeArr[level].Next
		for node != nil && level >= 0 {
			// 如果节点的值大于传入的值，就应该返回上个节点并进入下一层
			if !node.Value.Less(value) {
				headNodeInsertPositionArr[level] = node.Prev
				if node.Prev.Down == nil {
					if level-1 >= 0 {
						node = list.HeadNodeArr[level-1].Next
					} else {
						node = nil
					}
				} else {
					node = node.Prev.Down
				}
				level -= 1
			} else if node.Value.Less(value) {
				// 如果节点的值小于传入的值就进入下一个节点，如果下一个节点是 nil，说明本层已经查完了，进入下一层，且从下一层的头部开始
				if node.Next == nil {
					headNodeInsertPositionArr[level] = node
					level -= 1
					if level >= 0 {
						// 如果不是最底层继续进入到下一层
						node = list.HeadNodeArr[level].Next
					}
				} else {
					node = node.Next
				}
			}
		}
	}
	list.InsertValue(value, headNodeInsertPositionArr)
}

func (list *SkipList) DeleteNode(value Value) {
	node := list.HasNode(value)
	if node == nil {
		return
	}
	// 如果有节点就删除
	for node != nil {
		prevNode := node.Prev
		nextNode := node.Next

		prevNode.Next = nextNode
		if nextNode != nil {
			nextNode.Prev = prevNode
		}
		node = node.Down
	}
}

func (list *SkipList) HasNode(value Value) *Node {
	if list.Level < 0 {
		return nil
	}
	level := list.Level
	node := list.HeadNodeArr[level].Next
	for node != nil {
		if node.Value == value {
			return node
		}
		if !node.Value.Less(value) {
			// 如果节点的值大于传入的值，就应该返回上个节点并进入下一层
			if node.Prev.Down == nil {
				if level-1 >= 0 {
					node = list.HeadNodeArr[level-1].Next
				} else {
					node = nil
				}
			} else {
				node = node.Prev.Down
			}
			level -= 1
		} else if node.Value.Less(value) {
			// 如果节点的值小于传入的值就进入下一个节点，如果下一个节点是 nil，说明本层已经查完了，进入下一层，且从下一层的头部开始
			node = node.Next
			if node == nil {
				level -= 1
				if level >= 0 {
					// 如果不是最底层继续进入到下一层
					node = list.HeadNodeArr[level].Next
				}
			}
		}
	}
	return nil
}

func (list *SkipList) InsertValue(value Value, headNodeInsertPositionArr []*Node) {
	node := new(Node)
	node.Value = value
	// 插入最底层
	if list.Level < 0 {
		list.Level = 0
		list.HeadNodeArr[0] = new(Node)
		list.HeadNodeArr[0].Next = node
		node.Prev = list.HeadNodeArr[0]
		return
	}

	// 如果不是空的，就插入每一层
	// 插入最底层，
	rootNode := headNodeInsertPositionArr[0]
	nextNode := rootNode.Next

	rootNode.Next = node

	node.Prev = rootNode
	node.Next = nextNode

	if nextNode != nil {
		nextNode.Prev = node
	}

	currentLevel := 1
	for randLevel(0) && currentLevel <= list.Level+1 && currentLevel < MaxLevel {
		if headNodeInsertPositionArr[currentLevel] == nil {
			rootNode = new(Node)
			list.HeadNodeArr[currentLevel] = rootNode
		} else {
			rootNode = headNodeInsertPositionArr[currentLevel]
		}

		nextNode = rootNode.Next

		upNode := new(Node)
		upNode.Value = value
		upNode.Down = node
		upNode.Prev = rootNode
		upNode.Next = nextNode

		rootNode.Next = upNode
		if nextNode != nil {
			nextNode.Prev = node
		}

		node = upNode

		// 增加层数
		currentLevel++
	}
	list.Level = currentLevel - 1
}

// 通过抛硬币决定是否加入下一层
func randLevel(level int) bool {
	if level == 0 {
		randNum := rand.Intn(2)
		if randNum == 0 {
			return true
		}
	}
	//rand.Seed(time.Now().UnixNano())
	//for l := 1; rand.Float32() < Probability && l <= level; level++ {
	//	return true
	//}
	return false
}

// 初始化跳表
func NewSkipList() *SkipList {
	list := new(SkipList)
	list.Level = -1 // 设置层级别
	list.maxLevel = MaxLevel
	list.HeadNodeArr = make([]*Node, list.maxLevel) // 初始化头节点数组
	rand.Seed(time.Now().UnixNano())

	return list
}

func printSkipList(list *SkipList) {
	fmt.Println("====================start===============" + strconv.Itoa(list.Level))
	for i := list.Level; i >= 0; i-- {
		node := list.HeadNodeArr[i].Next
		for node != nil {
			fmt.Print(node.Value.String() + " -> ")
			node = node.Next
		}
		fmt.Println()
	}
	fmt.Println("====================end===============")

}
