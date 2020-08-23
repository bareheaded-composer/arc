## 题目
- [ ] 标题: [合法二叉搜索树](https://leetcode-cn.com/problems/legal-binary-search-tree-lcci/)
- [ ] 主体: 二叉树
- [ ] 算法: 先序遍历
- [ ] 时间复杂度: O()
- [ ] 代码:
```go
const INF = 1000000000000

func isValidBST(root *TreeNode) bool {
	return getIsValidBST(root, -INF, INF)
}

func getIsValidBST(root *TreeNode, minValue, maxValue int) bool {
	if root == nil {
		return true
	}
	return isInInterval(root.Val, minValue, maxValue) &&
		getIsValidBST(root.Left, minValue, root.Val) &&
		getIsValidBST(root.Right, root.Val, maxValue)
}

func isInInterval(value int, minValue, maxValue int) bool {
	return minValue < value && value < maxValue
}
```
