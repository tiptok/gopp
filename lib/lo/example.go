package main

import (
	"github.com/samber/lo"
	"log"
	"strconv"
)

func main() {
	sortStrings := lo.Uniq[string]([]string{"a1", "c2", "c1", "b1", "a1"})
	log.Println("1.去重", sortStrings)
	log.Println("1.1. 去重-为数组中的每个元素调用函数以生成计算值,再进行去重", lo.UniqBy[int, int]([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	}))

	transferTtoT1 := lo.Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	log.Printf("2.类型转换 int64->string （并行） %v", transferTtoT1)

	type user struct {
		Id int
	}
	userFlatMap := lo.FlatMap[int, user]([]int{1, 100, 999}, func(x int, _ int) []user {
		return []user{
			{Id: x},
		}
	})
	log.Printf("3.类型转换,扁平到一个列表 int->user %v", userFlatMap)

	log.Printf("4.包含 %v", lo.Contains[int]([]int{0, 1, 2, 3, 4, 5}, 6))
	log.Printf("4.包含-函数 %v", lo.ContainsBy[int]([]int{0, 1, 2, 3, 4, 5}, func(x int) bool {
		return x == 3
	}))

	// Reduces a collection to a single value
	log.Printf("5.合并 %v", lo.Reduce[int, int]([]int{1, 2, 3, 4, 5, 6}, func(agg int, item int, _ int) int {
		return agg + item
	}, 0))

	lo.ForEach[string]([]string{"A"}, func(x string, _ int) {
		log.Printf("6.遍历（并行） %v", x)
	})

	log.Printf("7.N次调用（索引作为调用值）（并行） %v", lo.Times[string](3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	}))

	log.Printf("8.分组(重复项分到同一组) slice -> map（并行） %v", lo.GroupBy[int, int]([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	}))

	log.Printf("9.分块(按数量分到同一组) [] -> [][] %v", lo.Chunk[int]([]int{0, 1, 2, 3, 4, 5}, 2))
	log.Printf("9.分块-函数(按数量分到同一组) [] -> [][] （并行） %v", lo.PartitionBy[int, string]([]int{-2, -1, 0, 1, 2, 3, 4, 5}, func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	}))

	log.Printf("10.倒序 %v", lo.Reverse[int]([]int{0, 1, 2, 3, 4, 5}))

	log.Printf("11.移除 %v", lo.Drop[int]([]int{0, 1, 2, 3, 4, 5}, 2))
	log.Printf("11.移除-DropRight %v", lo.DropRight[int]([]int{0, 1, 2, 3, 4, 5}, 2))
	log.Printf("11.移除-DropWhile %v", lo.DropWhile[int]([]int{0, 1, 2, 3, 4, 5}, func(v int) bool { return v <= 3 }))

	log.Printf("12.过滤 %v", lo.Reject[int]([]int{0, 1, 2, 3, 4, 5}, func(v int, _ int) bool { return v <= 3 }))

	//Range / RangeFrom / RangeWithSteps
	//result := Range(4)
	// [0, 1, 2, 3]
	//result := RangeFrom(1, 5);
	// [1, 2, 3, 4]
	//result := RangeWithSteps(0, 20, 5);
	// [0, 5, 10, 15]
	log.Printf("13.Map KEYS %v", lo.Keys[string, int](map[string]int{"foo": 1, "bar": 2}))
	log.Printf("13.Map VALUES %v", lo.Values[string, int](map[string]int{"foo": 1, "bar": 2}))
	log.Printf("13.Map VALUES-TRANSFORMS  %v", lo.MapValues[string, int, string](map[string]int{"foo": 1, "bar": 2}, func(x int, _ string) string {
		return strconv.Itoa(x)
	}))

	log.Printf("14.交集 %v", lo.Intersect[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2}))
	left, right := lo.Difference[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 6})
	log.Printf("14.差集 %v %v", left, right)
	log.Printf("14.并集 %v", lo.Union[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2}))

	// IndexOf \  LastIndex

	//TryCatch
	caught := false

	lo.TryCatch(func() error {
		panic("error")
		return nil
	}, func() {
		caught = true
	})
	log.Printf("15.TryCatch cought:%v", caught)

	lo.TryCatchWithErrorValue(func() error {
		panic("error1")
		return nil
	}, func(val any) {
		caught = val == "error"
	})
	log.Printf("15.TryCatchWithErrorValue cought:%v", caught)
}
