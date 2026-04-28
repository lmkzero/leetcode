package main

import (
	"fmt"
	"hash/fnv"
	"sync"
)

// Message 上游发送的消息体
type Message struct {
	UserID string
	Data   any
}

// Dispatcher 消息分发器与协程池管理器
type Dispatcher struct {
	// TODO 1: 补充需要的字段
	workers  int
	channels []chan Message
	wg       sync.WaitGroup
}

// NewDispatcher 初始化分发器
func NewDispatcher(workerNum int, bufferSize int) *Dispatcher {
	// TODO 2: 实例化 Dispatcher，并初始化内部结构
	d := &Dispatcher{
		workers:  workerNum,
		channels: make([]chan Message, workerNum),
	}
	// TODO 3: 启动 workerNum 个协程 (Worker)
	for i := range workerNum {
		d.channels[i] = make(chan Message, bufferSize)
		d.wg.Add(1)
		go d.worker(i, d.channels[i])
	}
	return d
}

// worker 实际处理消息的协程
func (d *Dispatcher) worker(workerID int, ch <-chan Message) {
	// TODO 4: 不断从 channel 中读取消息并处理，直到 channel 被关闭
	defer d.wg.Done()
	for msg := range ch {
		fmt.Printf("[Worker %d] UserID=%s, Data=%v\n", workerID, msg.UserID, msg.Data)
	}
}

// Dispatch 接收上游消息并分发
func (d *Dispatcher) Dispatch(msg Message) {
	// TODO 6: 核心路由逻辑。根据 UserID 确保相同用户的消息进入同一个 worker
	h := fnv.New32a()
	h.Write([]byte(msg.UserID))
	idx := h.Sum32() % uint32(d.workers)
	d.channels[idx] <- msg
}

// Close 优雅关闭协程池
func (d *Dispatcher) Close() {
	// TODO 7: 触发关闭信号，并阻塞等待所有 worker 将积压的消息处理完毕后再退出
	for _, ch := range d.channels {
		close(ch)
	}
	d.wg.Wait()
}
