package leetcode

/*
* 684. 冗余连接
* 在本问题中, 树指的是一个连通且无环的无向图。
* 输入一个图，该图由一个有着N个节点 (节点值不重复1, 2, ..., N) 的树及一条附加的边构成。
* 附加的边的两个顶点包含在1到N中间，这条附加的边不属于树中已存在的边。
* 结果图是一个以边组成的二维数组。每一个边的元素是一对[u, v] ，
* 满足 u < v，表示连接顶点u 和v的无向图的边。
* 返回一条可以删去的边，使得结果图是一个有着N个节点的树。如果有多个答案，则返回二维数组中最后出现的边。
* 答案边 [u, v] 应满足相同的格式 u < v。

* 示例 1：
* 输入: [[1,2], [1,3], [2,3]]
* 输出: [2,3]
* 解释: 给定的无向图为:
  1
 / \
2 - 3

* 示例 2：
* 输入: [[1,2], [2,3], [3,4], [1,4], [1,5]]
* 输出: [1,4]
* 解释: 给定的无向图为:
5 - 1 - 2
    |   |
    4 - 3

* 注意:
* 输入的二维数组大小在 3 到 1000。
* 二维数组中的整数在1到N之间，其中N是输入数组的大小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/redundant-connection

*/

//TODO: 并查集法。
//* 初始化每个节点 i 的 parent[i] 的值为 i 自己，对每条边进行遍历，
//* 且通过union来更新边的父节点从而更新到每个节点 i 的路径。
//* 当新加入判断的边的起始节点 from 的 parent[from] 和 目的地节点 to 的parent[to]一致时，
//* 说明两点已经存在于无环图中，该条边是 冗余的边，由于多余的边只有一条（N个节点的无环图，由N-1条边组成），故可以直接返回。
func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}

	//* 深度遍历路径上的点，找到起始节点
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	//* 查看是否是已经收敛于同一parent的边，不是的话就更行其parent[from]，相当于将节点 to 其加入了无环图中
	union := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		parent[x] = y
		return true
	}

	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return nil
}
