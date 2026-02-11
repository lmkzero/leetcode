package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	mids := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		a := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			b := int(num2[j] - '0')
			mids[i+j+1] += a * b
		}
	}
	for i := len(mids) - 1; i > 0; i-- {
		mids[i-1] += mids[i] / 10
		mids[i] %= 10
	}
	ans := []byte{}
	for i := 0; i < len(mids); i++ {
		if i == 0 && mids[i] == 0 {
			continue
		}
		ans = append(ans, byte('0'+mids[i]))
	}
	return string(ans)
}
