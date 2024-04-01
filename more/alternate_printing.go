// Package main 其他面试题目
package main

import "sync"

// 使用两个goroutine交替打印1-10
func alternatePrinting() {
	ch1, ch2 := make(chan int), make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		printFunc(ch1, ch2, "ch2")
	}()
	go func() {
		defer wg.Done()
		printFunc(ch2, ch1, "ch1")
	}()
	ch1 <- 1
	wg.Wait()
}

func printFunc(send, rec chan int, receiver string) {
	// chan关闭 or v+1 > 10时退出循环
	for v := range rec {
		println(receiver, ": ", v)
		if v+1 > 10 {
			close(send)
			break
		}
		send <- v + 1
	}
}
