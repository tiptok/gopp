package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

func main() {
	r := rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64()))

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()

	show := func(name string, v1, v2, v3 any) {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", name, v1, v2, v3)
	}
	// Float32和Float64的值在[0,1)范围内
	show("Float32", r.Float32(), r.Float32(), r.Float32())
	show("Float64", r.Float64(), r.Float64(), r.Float64())

	//ExpFloat64值的平均值为1，但呈指数衰减。
	show("ExpFloat64", r.ExpFloat64(), r.ExpFloat64(), r.ExpFloat64())

	//NormFloat64值的平均值为0，标准差为1。
	show("NormFloat64", r.NormFloat64(), r.NormFloat64(), r.NormFloat64())

	//Int32、Int64和Uint32生成给定宽度的值。
	show("Int32", r.Int32(), r.Int32(), r.Int32())
	show("Int64", r.Int64(), r.Int64(), r.Int64())
	show("Uint32", r.Uint32(), r.Uint32(), r.Uint32())

	//IntN、Int32N和Int64N将它们的输出限制为<n。
	//它们比使用r.Int()%n更加小心。
	show("IntN(10)", r.IntN(10), r.IntN(10), r.IntN(10))
	show("Int32N(10)", r.Int32N(10), r.Int32N(10), r.Int32N(10))
	show("Int64N(10)", r.Int64N(10), r.Int64N(10), r.Int64N(10))

	//Perm生成[0,n)范围内的随机排列。
	show("Perm", r.Perm(5), r.Perm(5), r.Perm(5))

	//打印一个位于半开区间[0,100)内的int64。
	fmt.Println("rand.N():", rand.N(int64(100)))

	//打印一个位于半开区间[0,100)内的uint32
	fmt.Println("rand.N():", rand.N(uint32(100)))

	//睡眠一个在0到100毫秒之间的随机时间。
	time.Sleep(rand.N(100 * time.Millisecond))

	//Shuffle使用默认的随机源对元素的顺序进行伪随机化
	words := strings.Fields("inkrunsfromthecornersofmymouth")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
}
