package main

import "math/rand"

type RandomizedSet struct {
	nums    []int
	indices map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:    []int{},
		indices: map[int]int{},
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.indices[val]; ok {
		return false
	}
	rs.indices[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	idx, ok := rs.indices[val]
	if !ok {
		return false
	}
	last := len(rs.nums) - 1
	lastVal := rs.nums[last]
	rs.nums[idx] = lastVal
	rs.indices[lastVal] = idx
	rs.nums = rs.nums[:last]
	delete(rs.indices, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}
