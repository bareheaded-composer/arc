## 题目
- [ ] 标题: [合法二叉搜索树](https://leetcode-cn.com/problems/legal-binary-search-tree-lcci/)
- [ ] 主体: 二叉树
- [ ] 算法: 中序遍历
- [ ] 时间复杂度: O()
- [ ] 代码:
```go
const INF = 1000000000000

var preNodeValue int

func isValidBST(root *TreeNode) bool {
	preNodeValue = -INF
	return getIsValidBST(root)
}

func getIsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !getIsValidBST(root.Left) {
		return false
	}
	if preNodeValue >= root.Val {
		return false
	}
	preNodeValue = root.Val
	if !getIsValidBST(root.Right) {
		return false
	}
	return true
}
```
