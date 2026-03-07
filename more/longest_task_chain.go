package main

import (
	"container/list"
)

// 任务编排系统
// 给一个数组，表示任务依赖关系，比如1完成才能做5​
// [[1,5],[2,6],[3,8],[4,10],[8,9],[5,6],[8,10],[6,7]] ​
// 设计一个任务系统完成这些任务，输出依赖链的最长路径节点的个数​​​

// 这个问题本质上是求有向无环图（DAG） 的最长路径长度（节点个数）：
// 1.构建任务依赖的有向图（邻接表）
// 2.计算每个节点的入度（用于拓扑排序）
// 3.通过拓扑排序 + 动态规划计算每个节点的最长路径长度
// 4.遍历所有节点的最长路径长度，找到最大值
func longestTaskChain(tasks [][]int) int {
	// 1. 构建邻接表（仅存储有后继的节点）+ 入度表（仅存储有前驱的节点）
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	// 收集所有节点（无需单独用nodeSet，最后遍历graph+inDegree的key即可）
	allNodes := make(map[int]bool)

	// 初始化图和入度
	for _, task := range tasks {
		from, to := task[0], task[1]
		graph[from] = append(graph[from], to)
		inDegree[to]++
		allNodes[from] = true
		allNodes[to] = true
	}

	// 2. 初始化DP（直接基于allNodes，无需提前遍历）
	dp := make(map[int]int)
	for node := range allNodes {
		dp[node] = 1 // 每个节点初始路径长度为1（自身）
	}

	// 3. 拓扑排序队列初始化（入度为0的节点）
	queue := list.New()
	for node := range allNodes {
		if inDegree[node] == 0 { // 入度表中无记录则入度为0
			queue.PushBack(node)
		}
	}

	// 4. 拓扑排序 + DP更新
	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(int)
		// 遍历当前节点的后继（无后继则跳过，无需处理）
		for _, next := range graph[curr] {
			if dp[next] < dp[curr]+1 {
				dp[next] = dp[curr] + 1
			}
			inDegree[next]--
			if inDegree[next] == 0 {
				queue.PushBack(next)
			}
		}
	}

	// 5. 找最长路径
	maxLen := 0
	for _, length := range dp {
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}
