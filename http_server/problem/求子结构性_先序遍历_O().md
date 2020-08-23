## 题目
- [ ] 标题: [树的子结构](https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/)
- [ ] 主体: 双二叉树
- [ ] 算法: 先序遍历
- [ ] 时间复杂度: O()
- [ ] 代码:
```go
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	return A.Val == B.Val && isSubStructureWithSpecificRoot(A.Left, B.Left) && isSubStructureWithSpecificRoot(A.Right, B.Right) ||
		isSubStructure(A.Left, B) ||
		isSubStructure(A.Right, B)
}

func isSubStructureWithSpecificRoot(specificRoot, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if specificRoot == nil {
		return false
	}
	return specificRoot.Val == B.Val && isSubStructureWithSpecificRoot(specificRoot.Left, B.Left) && isSubStructureWithSpecificRoot(specificRoot.Right, B.Right)
}
```
