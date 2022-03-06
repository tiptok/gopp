package syncT

import (
	"sync"
	"testing"
)

/*对象池*/
type Message struct {
	Id        int64
	Content   []byte
	SenderId  int64
	ReceiveId int64
	pool      Pool
}

func (m *Message) Reset() {
	m.Id = 0
	m.Content = m.Content[:0]
	m.SenderId = 0
	m.ReceiveId = 0
}

func (m *Message) Free() {
	m.pool.put(m)
}

type Pool struct {
	p *sync.Pool
}

func NewPool() Pool {
	return Pool{
		p: &sync.Pool{
			New: func() interface{} {
				return &Message{}
			},
		},
	}
}

func (p Pool) Get() *Message {
	m := p.p.Get().(*Message)
	m.Reset()
	m.pool = p
	return m
}

func (p Pool) put(m *Message) {
	p.p.Put(m)
}

func BenchmarkMessagePool(b *testing.B) {
	msgPool := NewPool()
	content := []byte("test pool")
	for i := 0; i <= b.N; i++ {
		m := msgPool.Get()
		m.Content = content
		m.SenderId = int64(i)
		m.ReceiveId = int64(i)
		m.Free()
	}
}
