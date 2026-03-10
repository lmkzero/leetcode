package main

import (
	"fmt"
	"sync"
	"time"
)

// 银行存钱
// A、B 两个线程，往同一个账户里各存50块​
// 看账户的余额

// Account 定义银行账户结构体
type Account struct {
	balance int          // 账户余额
	mu      sync.RWMutex // 互斥锁，保证并发安全
}

// Deposit 存钱方法（加锁保证原子性）
func (a *Account) Deposit(amount int) {
	a.mu.Lock()         // 加锁，其他线程需等待解锁
	defer a.mu.Unlock() // 方法结束时自动解锁，避免死锁
	a.balance += amount
	fmt.Printf("存入 %d 元，当前余额: %d 元\n", amount, a.balance)
}

// GetBalance 获取账户余额（加锁保证读取时无并发修改）
func (a *Account) GetBalance() int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.balance
}

func concurrentDeposit() {
	// 初始化账户，初始余额0元
	account := &Account{balance: 0}
	// 开启两个存钱goroutine
	var wg sync.WaitGroup
	wg.Add(2)
	// 线程A：存入50元
	go func() {
		defer wg.Done() // 线程完成后通知等待组
		account.Deposit(50)
	}()
	// 线程B：存入50元
	go func() {
		defer wg.Done()
		account.Deposit(50)
	}()
	// 等待两个线程执行完毕
	wg.Wait()
	// 输出最终余额
	fmt.Printf("\n账户最终余额: %d 元\n", account.GetBalance())
}

func concurrentDepositMore() {
	// 初始化账户，初始余额0元
	account := &Account{balance: 0}
	// 开启两个存钱goroutine
	var wg sync.WaitGroup
	wg.Add(2)
	// 线程A：分10次存入5元（总计50元）
	go func() {
		defer wg.Done()
		for range 10 {
			account.Deposit(5)
			time.Sleep(10 * time.Millisecond) // 模拟耗时操作
		}
	}()
	// 线程B：分10次存入5元（总计50元）
	go func() {
		defer wg.Done()
		for range 10 {
			account.Deposit(5)
			time.Sleep(10 * time.Millisecond)
		}
	}()
	// 等待两个线程执行完毕
	wg.Wait()
	// 输出最终余额
	fmt.Printf("\n账户最终余额: %d 元\n", account.GetBalance())
}
