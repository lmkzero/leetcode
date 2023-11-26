package main

func findAnagrams(s string, p string) []int {
	need := make(map[byte]int, len(p))
	window := make(map[byte]int)
	for _, b := range []byte(p) {
		need[b]++
	}
	left, right := 0, 0
	valid := 0
	var res []int
	for right < len(s) {
		b := s[right]
		right++
		if _, ok := need[b]; ok {
			window[b]++
			if window[b] == need[b] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}
