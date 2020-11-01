package Week_07

// func findCircleNum(M [][]int) int {
//     m, count := len(M), 0
//     visited := make([]int, m)
//     for i := 0; i < m; i++ {
//         if visited[i] == 0 {
//             dfs(M, visited, i)
//             count++
//         }
//     }
//     return count
// }

// func dfs(M [][]int, visited []int, i int) {
//     for j := 0; j < len(M); j++ {
//         if M[i][j] == 1 && visited[j] == 0 {
//             visited[j] = 1
//             dfs(M, visited, j)
//         }
//     }
// }
// 使用并查集
// 时间复杂度：O(N^2)
// 空间复杂度：O(N)
func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	length := len(M)
	ufind := NewUnionFind(length)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {	// 对称矩阵 不需要重复计算
			if M[i][j] == 1 {
				ufind.Union(i, j)
			}
		}
	}
	return ufind.count
}

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		count:  n,
		parent: parent,
	}
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
		u.count--
	}
}

// 路径压缩后 查询为O(1)
func (u *UnionFind) Find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] != i { // 路径压缩
		i, u.parent[i] = u.parent[i], root
	}
	return root
}

