package leetcode

import (
	"bytes"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

/*
LRU是Least Recently Used的简写，就是最近最少使用
参考 https://leetcode.com/problems/lru-cache/discuss/308956/Java-solution-HashMap-with-Doubly-linked-list-solution
*/
type LRUCache struct {
	head     *DoublyLink
	tail     *DoublyLink
	MapCache map[int]*DoublyLink
	cap      int
	Link     *DoublyLink
	len      int
	L        *sync.Mutex
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		MapCache: make(map[int]*DoublyLink),
		cap:      capacity,
		L:        new(sync.Mutex),
		len:      0,
	}
	cache.head = &DoublyLink{Value: 0}
	cache.tail = &DoublyLink{Value: 0}
	cache.head.Next = cache.tail
	cache.tail.Prev = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.MapCache[key]; ok {
		this.remove(v)
		this.moveToHead(v)
		this.MapCache[key] = v
		return v.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	k, ok := this.MapCache[key]
	if this.len < this.cap {
		if ok {
			this.remove(k)
		} else {
			this.len++
		}
	} else {
		if ok {
			this.remove(k)
		} else {
			this.remove(this.tail.Prev)
		}
	}
	this.add(key, value)
}

//add 添加节点
func (this *LRUCache) add(key, value int) {
	defer this.L.Unlock()
	this.L.Lock()
	node := &DoublyLink{
		Value: value,
		Key:   key,
	}
	this.MapCache[key] = node
	this.moveToHead(node)
}

//移除掉最后一个
func (this *LRUCache) remove(n *DoublyLink) {
	defer this.L.Unlock()
	this.L.Lock()
	delete(this.MapCache, n.Key)
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

//将node 移动到最头
func (this *LRUCache) moveToHead(node *DoublyLink) {
	hNext := this.head.Next
	node.Next = hNext
	hNext.Prev = node
	this.head.Next = node
	node.Prev = this.head
}

//Status  打印状态
func (this *LRUCache) Status() {
	log.Println("LRUCache Status->len:", this.len)
	bufMap := bytes.NewBufferString("")
	for k, v := range this.MapCache {
		bufMap.WriteString(fmt.Sprintf("->%v:(%v:%v)", k, v.Key, v.Value))
	}
	//log.Println("Map:",bufMap.String())
	buf := bytes.NewBufferString("")
	tmp := this.head
	for tmp != nil {
		buf.WriteString(fmt.Sprintf("->%v:%v", tmp.Key, tmp.Value))
		tmp = tmp.Next
	}
	log.Println("Link:", buf.String())
}

func Test_LRUCache(t *testing.T) {
	cache := Constructor(10)
	put := func(key, value int) {
		log.Println(fmt.Sprintf("Put %v:%v", key, value))
		cache.Put(key, value)
		cache.Status()
	}
	get := func(key int) {
		log.Println(fmt.Sprintf("Get: %v", key))
		result := cache.Get(key)
		log.Println(fmt.Sprintf("result: 【%v】", result))
		cache.Status()
	}
	inputM := []string{"put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"}
	input := [][]int{{10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26}}

	for i := 0; i < len(inputM); i++ {
		switch inputM[i] {
		case "put":
			put(input[i][0], input[i][1])
		case "get":
			get(input[i][0])
		}
	}

}

func GetAgeFromBirthDate(birth int64) int {
	n := time.Now()
	if birth == 0 {
		return 0
	}
	b := time.Unix(birth, 0)
	if birth > n.Unix() {
		return 0
	}
	age := n.Year() - b.Year()
	if b.AddDate(age, 0, 0).Unix() < n.Unix() {
		age += 1
	}
	return age
}

func Test_GetAgeFromBirthDate(t *testing.T) {
	getAge := func(y, m, d int) time.Time {
		return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	}
	b1 := getAge(1992, 6, 20)
	a1 := GetAgeFromBirthDate(b1.Unix())
	log.Printf("birth:%v age:%v", b1.Unix(), a1)
}
