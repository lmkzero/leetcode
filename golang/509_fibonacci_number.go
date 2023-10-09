package main

func fib(n int) int {
	base0, base1 := 0, 1
	if n == 0 {
		return base0
	}
	if n == 1 {
		return base1
	}
	cur := 0
	for i := 2; i <= n; i++ {
		cur = base0 + base1
		base0 = base1
		base1 = cur
	}
	return cur
}
