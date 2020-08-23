## 题目
- [ ] 标题: [字符串压缩](https://leetcode-cn.com/problems/compress-string-lcci/)
- [ ] 主体: 数组
- [ ] 算法: 双指针
- [ ] 时间复杂度: O()
- [ ] 代码:
```go
func compressString(S string) string {
	first, last := 0, 0
	result := strings.Builder{}
	for last < len(S) {
		for last+1 < len(S) && S[last+1] == S[last] {
			last++
		}
		result.WriteByte(S[first])
		//result.WriteString(fmt.Sprintf("%d", last-first+1))
		result.WriteString(strconv.Itoa(last-first+1))
		first = last + 1
		last = first
	}
	if result.Len() < len(S) {
		return result.String()
	}
	return S
}
```
