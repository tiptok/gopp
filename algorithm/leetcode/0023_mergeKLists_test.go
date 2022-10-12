package leetcode

import "testing"

func Test_mergeKLists(t *testing.T) {
	input := [][]*ListNode{
		{
			NewListNodes(1, 4, 5),
			NewListNodes(1, 3, 4),
			NewListNodes(2, 6),
		},
	}
	for i, _ := range input {
		out := mergeKLists1(input[i])
		PrintResult(out.PrintResult(), "", input[i])
	}
}

func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var retList *ListNode
	for i := 0; i < len(lists); i++ {
		if i == 0 {
			retList = lists[0]
			i++
		}
		retList = mergeTwoLists2(retList, lists[i])
	}
	return retList
}
