## 题目
- [ ] 标题: [二叉树的完全性检验](https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/)
- [ ] 主体: 二叉树
- [ ] 算法: 先序遍历
- [ ] 时间复杂度: O()
- [ ] 代码:
```go
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftTreeHeight, rightTreeHeight := getHeight(root.Left), getHeight(root.Right)
	return leftTreeHeight == rightTreeHeight && isFullTree(root.Left) && isCompleteTree(root.Right) ||
		leftTreeHeight == rightTreeHeight+1 && isCompleteTree(root.Left) && isFullTree(root.Right)
}

func isFullTree(root *TreeNode) bool {
	countOfNode := getCountOfNode(root)
	height := getHeight(root)
	return (1<<byte(height))-1 == countOfNode
}

func getCountOfNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getCountOfNode(root.Left) + getCountOfNode(root.Right) + 1
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
