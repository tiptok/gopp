package leetcode

/*


Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2]

Example 3:

Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].
*/

func plusOne(digits []int) []int {
	var res []int
	var carry = 1
	for i := len(digits) - 1; i >= 0; i-- {
		var sum = digits[i] + carry
		if sum >= 10 {
			carry = 1
		} else {
			carry = 1 / 10
		}
		res = append([]int{sum % 10}, res...)
	}
	if carry == 1 {
		res = append([]int{carry}, res...)
	}
	return res
}
