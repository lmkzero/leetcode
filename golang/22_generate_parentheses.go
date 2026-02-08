package main

func generateParenthesis(n int) []string {
	ans := []string{}
	var dfs func(int, int, string)
	dfs = func(left, right int, s string) {
		if left > n || right > n || left < right {
			return
		}
		if left == n && right == n {
			ans = append(ans, s)
			return
		}
		dfs(left+1, right, s+"(")
		dfs(left, right+1, s+")")
	}
	dfs(0, 0, "")
	return ans
}
