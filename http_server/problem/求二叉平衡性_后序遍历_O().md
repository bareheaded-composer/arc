## 题目
- [ ] 标题: [平衡二叉树](https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/)
- [ ] 主体: 二叉树
- [ ] 算法: 后序遍历
- [ ] 时间复杂度: O()
- [ ] 代码:
```go
const unBalancedFlag = -1

func isBalanced(root *TreeNode) bool {
	return getBalancedHeight(root) != unBalancedFlag
}

func getBalancedHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftBalancedHeight := getBalancedHeight(root.Left)
	rightBalancedHeight := getBalancedHeight(root.Right)

	if leftBalancedHeight == unBalancedFlag || rightBalancedHeight == unBalancedFlag {
		return unBalancedFlag
	}
	if abs(leftBalancedHeight-rightBalancedHeight) <= 1 {
		return max(rightBalancedHeight, leftBalancedHeight) + 1
	}
	return unBalancedFlag
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
```
