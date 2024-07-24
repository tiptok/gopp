package leetcode

/*

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。
*/

func intToRoman(num int) string {
	var result string
	for num > 0 {
		if num >= 1000 {
			num -= 1000
			result += "M"
		} else if num >= 900 {
			num -= 900
			result += "CM"
		} else if num >= 500 {
			num -= 500
			result += "D"
		} else if num >= 400 {
			num -= 400
			result += "CD"
		} else if num >= 100 {
			num -= 100
			result += "C"
		} else if num >= 90 {
			num -= 90
			result += "XC"
		} else if num >= 50 {
			num -= 50
			result += "L"
		} else if num >= 40 {
			num -= 40
			result += "XL"
		} else if num >= 10 {
			num -= 10
			result += "X"
		} else if num >= 9 {
			num -= 9
			result += "IX"
		} else if num >= 5 {
			num -= 5
			result += "V"
		} else if num == 4 {
			num -= 4
			result += "IV"
		} else {
			num -= 1
			result += "I"
		}
	}
	return result
}

func intToRoman1(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res, i := "", 0
	for num != 0 {
		for values[i] > num {
			i++
		}
		num -= values[i]
		res += symbols[i]
	}
	return res
}
