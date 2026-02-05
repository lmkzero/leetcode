package main

import "strings"

func restoreIpAddresses(s string) []string {
	n := len(s)
	temp := []string{}
	ans := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i >= n && len(temp) == 4 {
			ans = append(ans, strings.Join(temp, "."))
			return
		}
		if i >= n || len(temp) == 4 {
			return
		}
		x := 0
		for j := i; j < i+3 && j < n; j++ {
			x = x*10 + int(s[j]-'0')
			if x > 255 || (j > i && s[i] == '0') {
				break
			}
			temp = append(temp, s[i:j+1])
			dfs(j + 1)
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0)
	return ans
}
