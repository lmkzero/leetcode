package main

func majorityElement(nums []int) int {
	counter, res := 0, 0
	for _, v := range nums {
		if v == res {
			counter++
		} else if counter == 0 {
			res = v
		} else {
			counter--
		}
	}
	return res
}
