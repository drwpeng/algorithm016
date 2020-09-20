package Week_02

type Node struct {
	 Val int
	 Children []*Node
}
//589. N叉树的前序遍历
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	temp := []*Node{}
	res := []int{}
	temp = append(temp, root)
	for len(temp) > 0 {
		n := len(temp)
		node := temp[n-1]
		res = append(res, node.Val)
		temp = temp[:n-1]
		for i := len(node.Children)-1; i >= 0; i-- {
			temp = append(temp, node.Children[i])
		}
	}
	return res
}
// func preorder(root *Node) []int {
//     res := []int{}
//     pre(root, &res)
//     return res
// }
// func pre(root *Node, res *[]int) {
//     if root == nil {
//         return
//     }
//     *res = append(*res, root.Val)
//     for _, ch := range root.Children {
//         pre(ch, res)
//     }
// }
