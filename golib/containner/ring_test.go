package containner

import (
	"container/ring"
	"fmt"
	"testing"
)

func Test_Ring(t *testing.T) {
	r := ring.New(5)

	n := r.Len()

	// 多循环一次 0->5
	for i := 0; i < n+1; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring and print its contents
	r.Do(func(p interface{}) {
		fmt.Println(p.(int)) // 1 2 3 4 5
	})
}
