package main

var cache = make([][]string, 100)

func generateParenthesis(n int) []string {
	return generate(n)
}

func generate(n int) []string {
	if cache[n] != nil {
		return cache[n]
	}
	var ans []string
	if n == 0 {
		return []string{""}
	} else {
		for i := 0; i < n; i++ {
			lefts := generate(i)
			rights := generate(n - i - 1)
			for _, left := range lefts {
				for _, right := range rights {
					ans = append(ans, "("+left+")"+right)
				}
			}
		}
		cache[n] = ans
	}
	return cache[n]
}
