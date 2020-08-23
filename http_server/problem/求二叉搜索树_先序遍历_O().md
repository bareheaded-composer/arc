## 题目
- [ ] 标题: [最小高度树](https://leetcode-cn.com/problems/minimum-height-tree-lcci/)
- [ ] 主体: 数组
- [ ] 算法: 先序遍历
- [ ] 时间复杂度: O()
- [ ] 代码:
```go
class Solution {

    /**
     * @param Integer[] $nums
     * @return TreeNode
     */
    function sortedArrayToBST($nums) {
        return $this->getSortedArrayToBST($nums,0,count($nums)-1);
    }

    function getSortedArrayToBST($nums,$left,$right) {
        if ($left>$right){
            return null;
        }
        $mid=floor(($left+$right)/2);
        $root=new TreeNode($nums[$mid]);
        $root->left = $this->getSortedArrayToBST($nums,$left,$mid-1);
        $root->right = $this->getSortedArrayToBST($nums,$mid+1,$right);
        return $root;
    }
}
```
