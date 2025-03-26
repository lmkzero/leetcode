package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		chars := []byte(str)
		sort.SliceStable(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		key := string(chars)
		m[key] = append(m[key], str)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
