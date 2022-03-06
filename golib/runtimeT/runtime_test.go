package runtimeT

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"testing"
)

func Test_GetFuncForPCName(t *testing.T) {
	f := CallExample
	Func := runtime.FuncForPC(reflect.ValueOf(f).Pointer())
	log.Println(Func.Name())
	log.Println(Func.Entry())
}
func CallExample(s string) {
	log.Println(fmt.Sprintf("call:%v", s))
}
