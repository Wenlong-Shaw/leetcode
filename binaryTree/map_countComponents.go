package leetcode

/*
* 323. 无向图中连通分量的数目
* 给定编号从 0 到 n-1 的 n 个节点和一个无向边列表（每条边都是一对节点），
* 请编写一个函数来计算无向图中连通分量的数目。

* 示例 1:
* 输入: n = 5 和 edges = [[0, 1], [1, 2], [3, 4]]

     0          3
     |          |
     1 --- 2    4

* 输出: 2

* 示例 2:
* 输入: n = 5 和 edges = [[0, 1], [1, 2], [2, 3], [3, 4]]

     0           4
     |           |
     1 --- 2 --- 3

* 输出:  1

* 注意:
* 你可以假设在 edges 中不会出现重复的边。而且由于所以的边都是无向边，[0, 1] 与 [1, 0]  相同，
* 所以它们不会同时在 edges 中出现。
*/

//* 建立邻接表+深度优先遍历(DFS)每一条未走过的路
func countComponents(n int, edges [][]int) int {
	ans := 0
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = *(new([]int))
	}
	//TODO: 根据edges边表，构造邻接表adj
	for i := range edges {
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}
	//TODO: 如果没有深度遍历过该节点，则以该节点为源开始 dfs 。dfs完毕表示一条连通路走完了，计数ans+1
	visiteds := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visiteds[i] {
			dfs(i, adj, visiteds)
			ans++
		}
	}
	return ans
}

func dfs(src int, adj [][]int, visiteds []bool) {
	//* 源节点置为true，遍历过。
	visiteds[src] = true
	//* 深度遍历该源节点的邻接点。
	successors := adj[src]
	for i := range successors {
		if !visiteds[successors[i]] {
			dfs(successors[i], adj, visiteds)
		}
	}
	return
}

//! 效率低于DFS
//* 建立邻接表+广度优先遍历(BFS)每一条未走过的路
func countComponents_1(n int, edges [][]int) int {
	ans := 0
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = *(new([]int))
	}
	//TODO: 根据edges边表，构造邻接表adj
	for i := range edges {
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}
	//TODO: 如果没有广度遍历过该节点，则以该节点为源开始 bfs，没遍历到一个节点，就将其邻接表全部加入bfs遍历队列，
	//TODO: bfs完毕队列里最后一个未遍历的节点，表示以初始源节点为源的连通图全部点都走完了，进行下一轮未遍历过的初始src点的bfs遍历
	visiteds := make([]bool, n)
	queue := []int{}
	for i := 0; i < n; i++ {
		queue = append(queue, i)
		if !visiteds[i] {
			ans++

			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				if visiteds[cur] {
					continue
				}

				visiteds[i] = true
				for j := range adj[cur] {
					queue = append(queue, adj[cur][j])
				}
			}
		}
	}
	return ans
}
