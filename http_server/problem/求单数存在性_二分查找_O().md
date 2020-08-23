## 题目
- [ ] 标题: [搜索二维矩阵 II](https://leetcode-cn.com/problems/search-a-2d-matrix-ii/)
- [ ] 主体: 矩阵
- [ ] 算法: 二分查找
- [ ] 时间复杂度: O()
- [ ] 代码:
```go
func searchMatrix(matrix [][]int, target int) bool {
	rows, _ := getRowsAndCols(matrix)
	for i := 0; i < rows; i++ {
		if hasExistedInSortedArray(matrix[i], target) {
			return true
		}
	}
	return false
}

func hasExistedInSortedArray(sortedArray []int, target int) bool {
	left, right := 0, len(sortedArray)-1
	for left <= right {
		mid := (left + right) / 2
		if sortedArray[mid] == target {
			return true
		}
		if sortedArray[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}
```
