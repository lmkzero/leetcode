package main

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int, len(s1))
	window := make(map[byte]int)
	for _, c := range []byte(s1) {
		need[c]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
