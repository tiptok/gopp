package main

import (
	"github.com/tiptok/gocomm/pkg/log"
	"sort"
	"testing"
)

var exprNumberTable = [][]Expr{
	[]Expr{
		{OpChar: Eq, Value: []interface{}{100}},
		{OpChar: Eq, Value: []interface{}{500}},
		{OpChar: Range, Value: []interface{}{50, 200}},
		{OpChar: LessThanEqual, Value: []interface{}{Infinity, 50}},
		{OpChar: Range, Value: []interface{}{60, 100}},
		{OpChar: Range, Value: []interface{}{60, 70}},
	},
	[]Expr{
		{OpChar: Range, Value: []interface{}{100, 200}},
		{OpChar: Range, Value: []interface{}{Infinity, 50}},
		{OpChar: Range, Value: []interface{}{50, 90}},
		{OpChar: Range, Value: []interface{}{150, 300}},
	},
}

var exprDateTable = [][]Expr{
	[]Expr{
		{OpChar: Range, Value: []interface{}{1611731000, 1611735000}},
		{OpChar: LessThanEqual, Value: []interface{}{Infinity, 1611721000}},
		{OpChar: Range, Value: []interface{}{1611734000, 1611737000}},
	},
	[]Expr{
		{OpChar: Range, Value: []interface{}{1611731000, 1611735000}},
		{OpChar: LessThanEqual, Value: []interface{}{Infinity, 1611721000}},
		{OpChar: Range, Value: []interface{}{1611734000, 1611737000}},
		{OpChar: Recent, Value: []interface{}{5}},
	},
}

var exprCharsTable = [][]Expr{
	[]Expr{
		{OpChar: In, Value: []interface{}{"abc", "abd"}},
		{OpChar: Eq, Value: []interface{}{"abc"}},
	},
	[]Expr{
		{OpChar: In, Value: []interface{}{"华南", "华北"}},
		{OpChar: Eq, Value: []interface{}{"华北"}},
		{OpChar: Like, Value: []interface{}{"华"}},
		{OpChar: Like, Value: []interface{}{"中"}},
	},
}

// 数字表达式合并
func TestJoinNumberColumnExpr(t *testing.T) {
	for i := range exprNumberTable {
		out := JoinColumnExprNumber(exprNumberTable[i])
		t.Log(out)
	}
}

// 时间表达式合并
func TestJoinDateColumnExpr(t *testing.T) {
	for i := range exprDateTable {
		out := JoinColumnExprDate(exprDateTable[i])
		t.Log(out)
	}
}

// 字符串表达式合并
func TestJoinCharsColumnExpr(t *testing.T) {
	for i := range exprCharsTable {
		out := JoinCharsExprDate(exprCharsTable[i])
		t.Log(out)
	}
}

// 排序测试
func TestSortExprList(t *testing.T) {
	rec := RangeNumberExprCompute{valueType: ValueNumber}
	expr := exprNumberTable[0]
	for i := range expr {
		if expr[i].OpChar == Range || expr[i].OpChar == LessThanEqual || expr[i].OpChar == GreaterThanEqual {
			rec.expr = append(rec.expr, expr[i])
		}
	}
	var exprSort = exprSortable(rec.expr)
	log.Info("before:", exprSort)
	sort.Sort(exprSort)
	log.Info("after:", exprSort)
	rec.expr = exprSort
}
