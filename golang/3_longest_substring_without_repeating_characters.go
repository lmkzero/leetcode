package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	n := len(s)
	ans := 0
	for l, r := 0, 0; r < n; r++ {
		if idx, ok := m[s[r]]; ok && idx >= l {
			l = idx + 1
		}
		ans = max(ans, r-l+1)
		m[s[r]] = r
	}
	return ans
}
