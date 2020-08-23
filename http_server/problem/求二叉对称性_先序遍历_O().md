## 题目
- [ ] 标题: [对称的二叉树](https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/)
- [ ] 主体: 二叉树
- [ ] 算法: 先序遍历
- [ ] 时间复杂度: O()
- [ ] 代码:
```go
func isSymmetric(root *TreeNode) bool {
	if root==nil{
		return true
	}
	return isTwoTreeSymmetric(root.Left,root.Right)
}

func isTwoTreeSymmetric(treeA *TreeNode, treeB *TreeNode) bool {
	if treeA==nil && treeB==nil{
		return true
	}
	if treeA==nil || treeB==nil{
		return false
	}
	return treeA.Val==treeB.Val && isTwoTreeSymmetric(treeA.Left,treeB.Right) && isTwoTreeSymmetric(treeA.Right,treeB.Left)
}
```
