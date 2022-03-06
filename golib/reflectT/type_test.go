package reflectT

import (
	"reflect"
	"testing"
)

func Test_type(t *testing.T) {
	type V struct {
		Num     int
		Name    string
		Address string
	}
	//v :=&V{
	//	Num:1,
	//	Name:"tip",
	//	Address:"china",
	//}

	//vs :=[]V{
	//	{Num:2,Name:"cc",Address:"tt"},
	//	{Num:1,Name:"aa",Address:"ff"},
	//}

	parame := []interface{}{1, "cc", 2, "ff"}
	//tp :=reflect.TypeOf(v)
	inFunc := func(a int, b string) {}
	inFuncType := reflect.TypeOf(inFunc)
	var input []reflect.Value
	for i := 0; i < inFuncType.NumIn(); i++ {
		pt := inFuncType.In(i)
		for _, f := range parame {
			if reflect.TypeOf(f).AssignableTo(pt) {
				input = append(input, reflect.ValueOf(f))
				break
			}
		}
	}
}
