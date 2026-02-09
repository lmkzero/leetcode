package main

func mySqrt(x int) int {
	begin, end := 0, x
	ans := -1
	for begin <= end {
		mid := (begin + end) / 2
		t := mid * mid
		if t > x {
			end = mid - 1
			continue
		}
		if t <= x {
			begin = mid + 1
			ans = mid
			continue
		}
	}
	return ans
}
