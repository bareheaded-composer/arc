## 题目
- [ ] 标题: [单值二叉树](https://leetcode-cn.com/problems/univalued-binary-tree/)
- [ ] 主体: 二叉树
- [ ] 算法: 先序遍历
- [ ] 时间复杂度: O()
- [ ] 代码:
```go
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return getIsUnivalTree(root, root.Val)
}

func getIsUnivalTree(root *TreeNode, unival int) bool {
	if root == nil {
		return true
	}
	return root.Val == unival && getIsUnivalTree(root.Left, root.Val) && getIsUnivalTree(root.Right, root.Val)
}
```
