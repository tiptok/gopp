package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
)

type (
	ValueType string
	OpType    string
	Column    struct {
		// 字段
		ColumnId int
		// 列名称
		ColumnName string
		// 值类型
		ValueType ValueType
	}
	Expr struct {
		// 操作符
		OpChar OpType
		// 值
		Value []interface{}
	}
	ExprResult struct {
		// 操作符
		OpChar OpType
		// 值类型
		ValueType ValueType
		// 值
		Value []interface{}
	}
)

type (
	exprCompute interface {
		Append(result []ExprResult) []ExprResult
	}
	InExprCompute struct {
		expr      []Expr
		valueType ValueType
		OpType    OpType
	}
	RangeNumberExprCompute struct {
		expr      []Expr
		valueType ValueType
	}
	RecentDateExprCompute struct {
		expr      []Expr
		valueType ValueType
	}
)

const (
	ValueNumber ValueType = "number"
	ValueDate   ValueType = "date"
	ValueChars  ValueType = "chars"
)

// 表示无穷
const Infinity = -1
const day = 60 * 60 * 24

const (
	Eq       OpType = "="
	LessThan OpType = "<="
	MoreThan OpType = ">="
	Range    OpType = "range"
	In       OpType = "in"
	Like     OpType = "like"
	Recent   OpType = "recent" // 近几天
)

func (er ExprResult) NumberCompare(expr Expr) bool {
	_, x2 := getRange(er.Value[0], er.Value[1])
	y1, y2 := getRange(expr.Value[0], expr.Value[1])
	if y1 <= x2 {
		er.Value[1] = max(x2, y2)
		return true
	}
	if er.Value[1] == Infinity {
		return true
	}
	if er.Value[0] == Infinity && er.Value[1] == Infinity {
		return true
	}
	return false
}

func getRange(vi, vj interface{}) (float64, float64) {
	a, _ := strconv.ParseFloat(fmt.Sprintf("%v", vi), 64)
	b, _ := strconv.ParseFloat(fmt.Sprintf("%v", vj), 64)
	return a, b
}
func max(x, y float64) float64 {
	if x == Infinity {
		return x
	}
	if y == Infinity {
		return y
	}
	return math.Max(x, y)
}
func joinColumnExpr(ec []exprCompute) []ExprResult {
	var result []ExprResult
	for _, ecItem := range ec {
		if ecItem != nil {
			result = ecItem.Append(result)
		}
	}
	return result
}

func NewInExprCompute(expr []Expr, valueType ValueType) exprCompute {
	inExpr := InExprCompute{valueType: valueType}
	for i := 0; i < len(expr); i++ {
		if expr[i].OpChar == Eq || expr[i].OpChar == In {
			inExpr.expr = append(inExpr.expr, expr[i])
		}
	}
	if len(inExpr.expr) > 0 {
		return inExpr
	}
	return nil
}
func NewRangeExprCompute(expr []Expr, valueType ValueType) exprCompute {
	rec := RangeNumberExprCompute{valueType: valueType}
	for i := range expr {
		if expr[i].OpChar == Range || expr[i].OpChar == LessThan || expr[i].OpChar == MoreThan {
			rec.expr = append(rec.expr, expr[i])
		}
	}
	if len(rec.expr) == 0 {
		return nil
	}
	var exprSort = exprSortable(rec.expr)
	sort.Sort(exprSort)
	rec.expr = exprSort
	return rec
}
func NewRecentDateExprCompute(expr []Expr, valueType ValueType) exprCompute {
	inExpr := RecentDateExprCompute{valueType: valueType}
	for i := 0; i < len(expr); i++ {
		if expr[i].OpChar == Recent {
			inExpr.expr = append(inExpr.expr, expr[i])
		}
	}
	if len(inExpr.expr) > 0 {
		return inExpr
	}
	return nil
}
func NewLikeExprCompute(expr []Expr, valueType ValueType) exprCompute {
	inExpr := InExprCompute{valueType: valueType, OpType: Like}
	for i := 0; i < len(expr); i++ {
		if expr[i].OpChar == Like {
			inExpr.expr = append(inExpr.expr, expr[i])
		}
	}
	if len(inExpr.expr) > 0 {
		return inExpr
	}
	return nil
}

// in合并
func (ex InExprCompute) Append(result []ExprResult) []ExprResult {
	var res = ExprResult{
		OpChar:    In,
		ValueType: ex.valueType,
	}
	if ex.OpType != "" {
		res.OpChar = ex.OpType
	}
	for i := range ex.expr {
		res.Value = combine(append(res.Value, ex.expr[i].Value...))
	}
	result = append(result, res)
	return result
}
func combine(arr []interface{}) []interface{} {
	var mapArr = make(map[string]interface{})
	for i := range arr {
		key := fmt.Sprintf("%v", arr[i])
		if _, ok := mapArr[key]; !ok {
			mapArr[key] = arr[i]
		}
	}
	var res []interface{}
	for _, v := range mapArr {
		res = append(res, v)
	}
	return res
}

//范围合并
func (ex RangeNumberExprCompute) Append(result []ExprResult) []ExprResult {
	arr := &ExprResult{
		OpChar:    Range,
		ValueType: ex.valueType,
		Value:     ex.expr[0].Value,
	}
	for i := 1; i < len(ex.expr); i++ {
		if !arr.NumberCompare(ex.expr[i]) {
			result = append(result, *arr)
			arr.Value = ex.expr[i].Value
			continue
		}
	}
	result = append(result, *arr)
	return result
}

// recent范围
func (ex RecentDateExprCompute) Append(result []ExprResult) []ExprResult {
	var res = ExprResult{
		OpChar:    Recent,
		ValueType: ex.valueType,
	}
	var recent int64 = 0
	for i := range ex.expr {
		v, _ := strconv.ParseInt(fmt.Sprintf("%v", ex.expr[i].Value[0]), 10, 64)
		if v > recent {
			recent = v
		}
	}
	res.Value = append(res.Value, []interface{}{time.Now().Unix() - day*recent, Infinity})
	result = append(result, res)
	return result
}

type exprSortable []Expr

func (e exprSortable) Len() int {
	return len(e)
}
func (e exprSortable) Less(i, j int) bool {
	var a, b float64
	initValue := func(vi, vj interface{}) {
		a, _ = strconv.ParseFloat(fmt.Sprintf("%v", vi), 64)
		b, _ = strconv.ParseFloat(fmt.Sprintf("%v", vj), 64)
	}
	if e[i].Value[0] == Infinity && e[j].Value[0] != Infinity {
		return true
	}
	if e[i].Value[0] == Infinity && e[j].Value[0] == Infinity {
		initValue(e[i].Value[1], e[j].Value[1])
		return a < b
	}
	if e[i].Value[0] == e[j].Value[0] {
		initValue(e[i].Value[1], e[j].Value[1])
		return a < b
	}
	initValue(e[i].Value[0], e[j].Value[0])
	return a < b
}
func (e exprSortable) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
