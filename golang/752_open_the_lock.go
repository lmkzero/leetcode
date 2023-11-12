package main

func openLock(deadends []string, target string) int {
	deads := make(map[string]struct{}, len(deadends))
	for _, s := range deadends {
		deads[s] = struct{}{}
	}
	// 记录已经穷举过的密码，防止走回头路
	visited := make(map[string]struct{})
	visited["0000"] = struct{}{}
	q := []string{"0000"}
	// 从起点开始启动广度优先搜索
	step := 0

	for len(q) > 0 {
		sz := len(q)
		/* 将当前队列中的所有节点向周围扩散 */
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]

			/* 判断是否到达终点 */
			if _, ok := deads[cur]; ok {
				continue
			}
			if cur == target {
				return step
			}

			/* 将一个节点的未遍历相邻节点加入队列 */
			for j := 0; j < 4; j++ {
				up := plusOne(cur, j)
				if _, ok := visited[up]; !ok {
					q = append(q, up)
					visited[up] = struct{}{}
				}
				down := minusOne(cur, j)
				if _, ok := visited[down]; !ok {
					q = append(q, down)
					visited[down] = struct{}{}
				}
			}
		}
		/* 在这里增加步数 */
		step++
	}
	// 如果穷举完都没找到目标密码，那就是找不到了
	return -1

}

func plusOne(s string, j int) string {
	arr := []byte(s)
	if arr[j] == '9' {
		arr[j] = '0'
	} else {
		arr[j]++
	}
	return string(arr)
}

func minusOne(s string, j int) string {
	arr := []byte(s)
	if arr[j] == '0' {
		arr[j] = '9'
	} else {
		arr[j]--
	}
	return string(arr)
}
