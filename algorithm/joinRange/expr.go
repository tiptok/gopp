package main

func JoinColumnExprNumber(expr []Expr) []ExprResult {
	var ec = []exprCompute{NewInExprCompute(expr, ValueNumber), NewRangeExprCompute(expr, ValueNumber)}
	return joinColumnExpr(ec)
}

func JoinColumnExprDate(expr []Expr) []ExprResult {
	var ec = []exprCompute{NewRangeExprCompute(expr, ValueDate), NewRecentDateExprCompute(expr, ValueDate)}
	return joinColumnExpr(ec)
}

func JoinCharsExprDate(expr []Expr) []ExprResult {
	var ec = []exprCompute{NewInExprCompute(expr, ValueChars), NewLikeExprCompute(expr, ValueChars)}
	return joinColumnExpr(ec)
}
