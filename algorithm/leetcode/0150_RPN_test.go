package leetcode

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

//逆波兰式  指的是后缀表达法
func Test_evalRPN(t *testing.T) {
	input := [][]string{
		[]string{"2"},
		[]string{"2", "1", "+", "3", "*"},
		[]string{"4", "13", "5", "/", "+"},
		[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
	}
	for i := range input {
		out := evalRPN(input[i])
		log.Println(fmt.Sprintf("input :%v out:%v", input[i], out))
	}
}

func Benchmark_evalRPN(b *testing.B) {
	input := []string{"2", "1", "+", "3", "*"}
	for i := 0; i < b.N; i++ {
		evalRPN(input)
	}
}

func evalRPN(tokens []string) int {
	result := 0
	var stack []string
	pop := func(stack []string) (string, []string, bool) {
		if len(stack) == 0 {
			return "", stack, false
		}
		return stack[len(stack)-1], stack[:len(stack)-1], true
	}
	atoiconv := func(s string) int {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return num
	}
	if len(tokens) == 1 {
		return atoiconv(tokens[0])
	}
	compute := func(v1, v2 int, op string) (ret int, err error) {
		switch op {
		case "*":
			ret = v1 * v2
			return
		case "/":
			ret = v1 / v2
			return
		case "+":
			ret = v1 + v2
			return
		case "-":
			ret = v1 - v2
			return
		}
		err = fmt.Errorf("op not found:" + op)
		return
	}
	var (
		num1, num2, op string
		ok             bool
	)
	for i := 0; i < len(tokens); i++ {
		op = tokens[i]
		if _, err := compute(1, 1, op); err != nil {
			stack = append(stack, op)
		} else {
			num1, stack, ok = pop(stack)
			num2, stack, ok = pop(stack)
			if !ok {
				//log.Println("stack out index")
				return 0
			}

			result, err = compute(atoiconv(num2), atoiconv(num1), op)
			if err != nil {
				log.Println(err)
				return 0
			}
			//log.Println(num2,op,num1,"=",result)
			stack = append(stack, fmt.Sprintf("%v", result))
		}
	}
	return result
}
