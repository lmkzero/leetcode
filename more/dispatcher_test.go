package main

import (
	"testing"
)

func Test_Dispatcher(t *testing.T) {
	tests := []struct {
		name       string
		dispatcher *Dispatcher
		messages   []Message
	}{
		{
			name:       "1",
			dispatcher: NewDispatcher(3, 10), // 初始化 3 个 Worker，每个 Channel 缓冲设为 10
			messages: []Message{ // 模拟上游发送消息
				{"user_A", "登录"},
				{"user_B", "购买商品"},
				{"user_A", "退出登录"}, // 必须和上一条 user_A 消息在同一个 worker 执行
				{"user_C", "浏览页面"},
				{"user_B", "评价商品"}, // 必须和上一条 user_B 消息在同一个 worker 执行
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, msg := range tt.messages {
				tt.dispatcher.Dispatch(msg)
			}
			tt.dispatcher.Close()
		})
	}
}
