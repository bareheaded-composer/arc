## 题目
- [ ] 标题: [特定深度节点链表](https://leetcode-cn.com/problems/list-of-depth-lcci/)
- [ ] 主体: 二叉树
- [ ] 算法: 先序遍历
- [ ] 时间复杂度: O()
- [ ] 代码:
```go
var listHeads []*ListNode
var listTails []*ListNode

func listOfDepth(tree *TreeNode) []*ListNode {
	listHeads = make([]*ListNode, 0)
	listTails = make([]*ListNode, 0)
	formLists(tree, 0)
	return listHeads
}

func formLists(root *TreeNode, nowDepth int) {
	if root == nil {
		return
	}
	node := &ListNode{Val: root.Val}
	if nowDepth == len(listTails) {
		listTails = append(listTails, node)
		listHeads = append(listHeads, node)
	} else {
		tail := listTails[nowDepth]
		tail.Next = node
		listTails[nowDepth] = node
	}
	formLists(root.Left, nowDepth+1)
	formLists(root.Right, nowDepth+1)
}
```
