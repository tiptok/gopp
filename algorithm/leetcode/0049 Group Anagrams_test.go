package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	var res = make([][]string, 0)
	var mapResult = make(map[string][]string, 0)
	for i := range strs {
		key := []rune(strs[i])
		sort.Slice(key, func(i, j int) bool {
			return key[i] < key[j]
		})
		mapResult[string(key)] = append(mapResult[string(key)], strs[i])
	}
	for _, v := range mapResult {
		res = append(res, v)
	}
	return res
}

type Key [26]byte

func strKey(str string) (key Key) {
	for i := range str {
		key[str[i]-'a']++
	}
	return
}

func groupAnagramsFast(strs []string) [][]string {
	groups := make(map[Key][]string)

	for _, v := range strs {
		key := strKey(v)
		groups[key] = append(groups[key], v)
	}

	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}
