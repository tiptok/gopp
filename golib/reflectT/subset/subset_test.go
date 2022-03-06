package subset

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type test struct {
	Name string
	A    interface{}
	B    interface{}
}

type sample struct {
	Answer  int
	Name    string
	Child   *sample
	private string
}

var (
	subsetTests = []test{
		test{"Integers", 1, 1},
		test{"Strings", "a", "a"},
		test{"Maps",
			map[string]string{"foo": "bar"},
			map[string]string{"foo": "bar"}},
		test{"Maps subset",
			map[string]string{"foo": "bar"},
			map[string]string{"foo": "bar", "answer": "42"}},
		test{"Nil map", nil, map[string]string{"foo": "bar"}},
		test{"Structs", sample{Answer: 1}, sample{Answer: 1}},
		test{"Struct subset", sample{Answer: 1}, sample{Answer: 1, Name: "a"}},
		test{"Nil pointer", sample{}, sample{Child: &sample{}}},
	}
	notSubsetTests = []test{
		test{"Integer", 1, 2},
		test{"Integers of different types", uint(1), int(1)},
		test{"Maps not subset",
			map[string]string{"foo": "bar", "answer": "42"},
			map[string]string{"foo": "bar"}},
		test{"Structs", sample{Answer: 1}, sample{Answer: 2}},
		test{"Struct subset",
			sample{Answer: 1, Name: "b"},
			sample{Answer: 1, Name: "a"}},
	}
)

func TestSubsets(t *testing.T) {
	for _, d := range subsetTests {
		if !Check(d.A, d.B) {
			t.Errorf("Was expecting \"%s\" Check(%v, %v) == true.", d.Name, d.A, d.B)
		}
	}
}

func TestNotSubsets(t *testing.T) {
	for _, d := range notSubsetTests {
		if Check(d.A, d.B) {
			t.Errorf("Was expecting %s Check(%v, %v) == false.", d.Name, d.A, d.B)
		}
	}
}

func TestAddress(t *testing.T) {
	//s:=&sample{Answer:1,}
	ar := []int{5, 6}
	v := reflect.ValueOf(ar[0])
	t.Log(v.CanAddr())

	checkCanAddr()
}

type S struct {
	X int
	Y string
	z int
}

func M() int {
	return 100
}

var x0 = 0

func checkCanAddr() {
	// 可寻址的情况
	v := reflect.ValueOf(x0)
	fmt.Printf("x0: %v \tcan be addressable and set: %t, %t\n", x0, v.CanAddr(), v.CanSet()) //false,false
	var x1 = 1
	v = reflect.Indirect(reflect.ValueOf(x1))
	fmt.Printf("x1: %v \tcan be addressable and set: %t, %t\n", x1, v.CanAddr(), v.CanSet()) //false,false
	var x2 = &x1
	v = reflect.Indirect(reflect.ValueOf(x2))
	fmt.Printf("x2: %v \tcan be addressable and set: %t, %t\n", x2, v.CanAddr(), v.CanSet()) //true,true
	var x3 = time.Now()
	v = reflect.Indirect(reflect.ValueOf(x3))
	fmt.Printf("x3: %v \tcan be addressable and set: %t, %t\n", x3, v.CanAddr(), v.CanSet()) //false,false
	var x4 = &x3
	v = reflect.Indirect(reflect.ValueOf(x4))
	fmt.Printf("x4: %v \tcan be addressable and set: %t, %t\n", x4, v.CanAddr(), v.CanSet()) // true,true
	var x5 = []int{1, 2, 3}
	v = reflect.ValueOf(x5)
	fmt.Printf("x5: %v \tcan be addressable and set: %t, %t\n", x5, v.CanAddr(), v.CanSet()) // false,false
	var x6 = []int{1, 2, 3}
	v = reflect.ValueOf(x6[0])
	fmt.Printf("x6: %v \tcan be addressable and set: %t, %t\n", x6[0], v.CanAddr(), v.CanSet()) //false,false
	var x7 = []int{1, 2, 3}
	v = reflect.ValueOf(x7).Index(0)
	fmt.Printf("x7: %v \tcan be addressable and set: %t, %t\n", x7[0], v.CanAddr(), v.CanSet()) //true,true
	v = reflect.ValueOf(&x7[1])
	fmt.Printf("x7.1: %v \tcan be addressable and set: %t, %t\n", x7[1], v.CanAddr(), v.CanSet()) //true,true
	var x8 = [3]int{1, 2, 3}
	v = reflect.ValueOf(x8[0])
	fmt.Printf("x8: %v \tcan be addressable and set: %t, %t\n", x8[0], v.CanAddr(), v.CanSet()) //false,false
	// https://groups.google.com/forum/#!topic/golang-nuts/RF9zsX82MWw
	var x9 = [3]int{1, 2, 3}
	v = reflect.Indirect(reflect.ValueOf(x9).Index(0))
	fmt.Printf("x9: %v \tcan be addressable and set: %t, %t\n", x9[0], v.CanAddr(), v.CanSet()) //false,false
	var x10 = [3]int{1, 2, 3}
	v = reflect.Indirect(reflect.ValueOf(&x10)).Index(0)
	fmt.Printf("x9: %v \tcan be addressable and set: %t, %t\n", x10[0], v.CanAddr(), v.CanSet()) //true,true
	var x11 = S{}
	v = reflect.ValueOf(x11)
	fmt.Printf("x11: %v \tcan be addressable and set: %t, %t\n", x11, v.CanAddr(), v.CanSet()) //false,false
	var x12 = S{}
	v = reflect.Indirect(reflect.ValueOf(&x12))
	fmt.Printf("x12: %v \tcan be addressable and set: %t, %t\n", x12, v.CanAddr(), v.CanSet()) //true,true
	var x13 = S{}
	v = reflect.ValueOf(x13).FieldByName("X")
	fmt.Printf("x13: %v \tcan be addressable and set: %t, %t\n", x13, v.CanAddr(), v.CanSet()) //false,false
	var x14 = S{}
	v = reflect.Indirect(reflect.ValueOf(&x14)).FieldByName("X")
	fmt.Printf("x14: %v \tcan be addressable and set: %t, %t\n", x14, v.CanAddr(), v.CanSet()) //true,true
	var x15 = S{}
	v = reflect.Indirect(reflect.ValueOf(&x15)).FieldByName("z")
	fmt.Printf("x15: %v \tcan be addressable and set: %t, %t\n", x15, v.CanAddr(), v.CanSet()) //true,false
	v = reflect.Indirect(reflect.ValueOf(&S{}))
	fmt.Printf("x15.1: %v \tcan be addressable and set: %t, %t\n", &S{}, v.CanAddr(), v.CanSet()) //true,true
	var x16 = M
	v = reflect.ValueOf(x16)
	fmt.Printf("x16: %p \tcan be addressable and set: %t, %t\n", x16, v.CanAddr(), v.CanSet()) //false,false
	var x17 = M
	v = reflect.Indirect(reflect.ValueOf(&x17))
	fmt.Printf("x17: %p \tcan be addressable and set: %t, %t\n", x17, v.CanAddr(), v.CanSet()) //true,true
	var x18 interface{} = &x11
	v = reflect.ValueOf(x18)
	fmt.Printf("x18: %v \tcan be addressable and set: %t, %t\n", x18, v.CanAddr(), v.CanSet()) //false,false
	var x19 interface{} = &x11
	v = reflect.ValueOf(x19).Elem()
	fmt.Printf("x19: %v \tcan be addressable and set: %t, %t\n", x19, v.CanAddr(), v.CanSet()) //true,true
	var x20 = [...]int{1, 2, 3}
	v = reflect.ValueOf([...]int{1, 2, 3})
	fmt.Printf("x20: %v \tcan be addressable and set: %t, %t\n", x20, v.CanAddr(), v.CanSet()) //false,false
}
