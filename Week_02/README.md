# 学习笔记

### golang 的堆必须实现以下几个接口
```
// An IntHeap is a min-heap of ints.
type IntHeap []int  //类型根据实际需要实现

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }  // 大顶堆（<）或小顶堆(>)
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
```


### 二叉树的遍历
* 前序：根左右
* 中序：左根右
* 后序：左右根

示例代码：后序 **递归**
```
func postorder(root *TreeNode, res *[]int) {
      if root == nil {
          return
      }
      postorder(root.Left, res)
      postorder(root.Right, res)
      *res = append(*res, root.Val)
  }
```
后序---**迭代**
```
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	slice := make([]int, 0)
	for len(stack) > 0 {
		root := stack[len(stack) - 1]
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Left == nil && root.Right == nil {
			stack = stack[:len(stack) - 1]
			slice = append(slice, root.Val)
		}
		root.Left = nil
		root.Right = nil
	}
	return slice
}
```
### 347. 前 K 个高频元素
* 使用hash记录每个元素的次数
* 元素个数为k的小顶堆，大于堆顶则插入，小于堆顶则舍弃
```
func topKFrequent(nums []int, k int) []int {
	myhash := make(map[int]int)
	for _, v := range nums {
		myhash[v]++
	}

	h := intHeap{}
	heap.Init(&h)
	for key, v := range myhash {
		heap.Push(&h, [2]int{key, v})
		if h.Len() > k {
			heap.Pop(&h)
		}
	}
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&h).([2]int)[0])
	}
	return res
}
```


