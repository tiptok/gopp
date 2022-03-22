package main

import (
	"fmt"
	"testing"
)

func SumIntsOrFloats[V int64 | float64](m []V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func TestSum(t *testing.T) {
	input := []int64{1, 2, 3, 4, 5}
	out := SumIntsOrFloats(input)

	t.Log(out)
}

func Sum[T int](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func TestSum1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	out := Sum(input...)
	t.Log(SumAny(input...))
	t.Log(SumAny([]id{1, 2, 3}...))
	t.Log(out)
}

func SumAny[T Numeric](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

type Numeric interface {
	uint | uint8 | ~int | float64
}

type id int

func MapData[F, T any](s []F, f func(F) T) []T {
	var result = make([]T, 0)
	for i := 0; i < len(s); i++ {
		result = append(result, f(s[i]))
	}
	return result
}

func TestMapData(t *testing.T) {
	var s []int
	f := func(i int) int64 { return int64(i) }
	var r []int64
	// 1.明确指定两个类型参数.
	//r = MapData[int, int64](s, f)
	// 2.只指定第一个类型参数，对于 F，并推断出 T
	//r = MapData[int](s, f)
	// 3.不要指定任何类型参数，并让两者都被推断.
	r = MapData(s, f)
	t.Log(r)
}

type SumFn[T Numeric] func(...T) T

func PrintIDAndSum[T ~string, K Numeric](id T, sum SumFn[K], values ...K) {
	// The format string uses "%v" to emit the sum since using "%d" would
	// be invalid if the value type was a float or complex variant.
	fmt.Printf("%s has a sum of %v\n", id, sum(values...))
}

func TestPrintIDAndSum(T *testing.T) {

	// PrintIDAndSum("acct-1", SumAny, 1, 2, 3)
	// 会出错 cannot use generic function Sum without instantiation
	// 无法推断出SumAny函数是数字类型,需要显示指定类型如下

	//PrintIDAndSum("acct-1", SumAny[int], 1, 2, 3)
}

/***声明对象***/

type Ledger[T ~string, K Numeric] struct {
	ID T

	Amounts []K

	SumFn SumFn[K]
}

// 由于 Ledger 具有泛型类型，因此它们的符号必须包含在函数接收器中
func (l Ledger[T, K]) PrintIDAndSum() string {
	return fmt.Sprintf("%s has a sum of %v\n", l.ID, l.SumFn(l.Amounts...))
}

func TestLedger(t *testing.T) {
	result := Ledger[string, int]{
		ID:      "acct-1",
		Amounts: []int{1, 2, 3},
		SumFn:   SumAny[int],
	}.PrintIDAndSum()
	t.Log(result)
}

/*Structural constraints  结构约束*/
type HasID interface {
	~struct {
		ID string
	}
}

type Unique struct {
	ID string
}

type CanSetID interface {
	SetID(string)
}

func (u Unique) SetID(s string) {
	u.ID = s
}

func NewHasID[T HasID]() T {
	var t T
	return t
}

func NewCanSetID[T CanSetID]() T {
	var t T
	return t
}

/*
	以下是不允许的,复合约束不能包含具体类型的联合，
	例如 Go 原语或结构类型和接口类型。
*/
func NewT[T HasID | CanSetID]() T {
	var t T
	return t
}

func TestStructuralConstraint(t *testing.T) {
	t.Logf("%T\n", NewHasID[Unique]())
	t.Logf("%T\n", NewCanSetID[Unique]())
	t.Logf("%T\n", NewCanSetID[*Unique]())

	// output:
	//generic_test.go:149: main.Unique
	//generic_test.go:150: main.Unique
	//generic_test.go:151: *main.Unique
}
