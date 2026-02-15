package main

import "strings"

func decodeString(s string) string {
	intQ := []int{}
	stringQ := []string{}
	num := 0
	res := ""
	for _, c := range s {
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
			continue
		}
		if c == '[' {
			intQ = append(intQ, num)
			stringQ = append(stringQ, res)
			num = 0
			res = ""
			continue
		}
		if c == ']' {
			var t strings.Builder
			for i, n := 0, intQ[len(intQ)-1]; i < n; i++ {
				t.WriteString(res)
			}
			res = stringQ[len(stringQ)-1] + t.String()
			intQ = intQ[:len(intQ)-1]
			stringQ = stringQ[:len(stringQ)-1]
			continue
		}
		res += string(c)
	}
	return res
}
