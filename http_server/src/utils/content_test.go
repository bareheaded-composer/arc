package utils

import (
	"arc/src/model"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"testing"
)
func TestMakeFilerContentToFile(t *testing.T) {
	unitJson := `{
      "topic": "最大黑方阵",
      "link": "https://leetcode-cn.com/problems/max-black-square-lcci/",
      "main_body": "矩阵",
      "solution": "动态规划",
      "purpose": "求最大方阵位置",
      "comment": "方阵形成条件: 可以是空心的。",
      "code": "func findSquare(matrix [][]int) []int {\n\tx, y, size := getMaxBlankSquareCoordinateAndSize(matrix)\n\tif x == -1 {\n\t\treturn []int{}\n\t}\n\treturn []int{x - size + 1, y - size + 1, size}\n}\n\nfunc getMaxBlankSquareCoordinateAndSize(matrix [][]int) (int, int, int) {\n\tmaxWidth, maxHeight := getMaxWidthAndMaxHeight(matrix)\n\trows, cols := getRowsAndCols(matrix)\n\tmaxBlankSquareX, maxBlankSquareY, maxSize := -1, -1, 0\n\tfor i := 0; i < rows; i++ {\n\t\tfor t := 0; t < cols; t++ {\n\t\t\tcurBlankSquareSize := getBlankSquareSize(maxWidth, maxHeight, i, t)\n\t\t\tif curBlankSquareSize > maxSize {\n\t\t\t\tmaxBlankSquareX = i\n\t\t\t\tmaxBlankSquareY = t\n\t\t\t\tmaxSize = curBlankSquareSize\n\t\t\t}\n\t\t}\n\t}\n\treturn maxBlankSquareX, maxBlankSquareY, maxSize\n}\n\nfunc getMaxWidthAndMaxHeight(matrix [][]int) ([][]int, [][]int) {\n\trows, cols := getRowsAndCols(matrix)\n\tmaxWidth := get2DSlice(rows, cols)\n\tmaxHeight := get2DSlice(rows, cols)\n\tmaxWidth[0][0] = matrix[0][0] ^ 1\n\tmaxHeight[0][0] = matrix[0][0] ^ 1\n\tfor i := 1; i < rows; i++ {\n\t\tmaxWidth[i][0] = matrix[i][0] ^ 1\n\t\tif matrix[i][0] == 0 {\n\t\t\tmaxHeight[i][0] = maxHeight[i-1][0] + 1\n\t\t}\n\t}\n\tfor t := 1; t < cols; t++ {\n\t\tmaxHeight[0][t] = matrix[0][t] ^ 1\n\t\tif matrix[0][t] == 0 {\n\t\t\tmaxWidth[0][t] = maxWidth[0][t-1] + 1\n\t\t}\n\t}\n\tfor i := 1; i < rows; i++ {\n\t\tfor t := 1; t < cols; t++ {\n\t\t\tif matrix[i][t] == 0 {\n\t\t\t\tmaxHeight[i][t] = maxHeight[i-1][t] + 1\n\t\t\t\tmaxWidth[i][t] = maxWidth[i][t-1] + 1\n\t\t\t}\n\t\t}\n\t}\n\treturn maxWidth, maxHeight\n}\n\nfunc getBlankSquareSize(maxWidth, maxHeight [][]int, x, y int) int {\n\tmaxSize := 0\n\tfor size := 1; size <= min(maxHeight[x][y], maxWidth[x][y]); size++ {\n\t\tif maxHeight[x][y-size+1] >= size && maxWidth[x-size+1][y] >= size {\n\t\t\tmaxSize = size\n\t\t}\n\t}\n\treturn maxSize\n}\n\nfunc getRowsAndCols(matrix [][]int) (int, int) {\n\tif len(matrix) == 0 {\n\t\treturn 0, 0\n\t}\n\treturn len(matrix), len(matrix[0])\n}\n\nfunc get2DSlice(rows, column int) [][]int {\n\tslice := make([][]int, rows)\n\tfor i := 0; i < len(slice); i++ {\n\t\tslice[i] = make([]int, column)\n\t}\n\treturn slice\n}\n\nfunc min(arr ...int) int {\n\tif len(arr) == 1 {\n\t\treturn arr[0]\n\t}\n\ta, b := arr[0], min(arr[1:]...)\n\tif a > b {\n\t\treturn b\n\t}\n\treturn a\n}"
    }`

	var unit model.Unit
	if err := json.Unmarshal([]byte(unitJson),&unit);err!=nil{
		logs.Error(err)
		return
	}
	logs.Debug(
		MakeFilerContentToFile(
		`C:\Users\hasee\Desktop\arc\http_server\problem`,
		`C:\Users\hasee\Desktop\arc\http_server\template\unit_show_template`,
		&unit,
		),
	)
}


func TestGetContent(t *testing.T) {
	unitJson := `{
      "topic": "最大黑方阵",
      "link": "https://leetcode-cn.com/problems/max-black-square-lcci/",
      "main_body": "矩阵",
      "solution": "动态规划",
      "purpose": "求最大方阵位置",
      "comment": "方阵形成条件: 可以是空心的。",
      "code": "func findSquare(matrix [][]int) []int {\n\tx, y, size := getMaxBlankSquareCoordinateAndSize(matrix)\n\tif x == -1 {\n\t\treturn []int{}\n\t}\n\treturn []int{x - size + 1, y - size + 1, size}\n}\n\nfunc getMaxBlankSquareCoordinateAndSize(matrix [][]int) (int, int, int) {\n\tmaxWidth, maxHeight := getMaxWidthAndMaxHeight(matrix)\n\trows, cols := getRowsAndCols(matrix)\n\tmaxBlankSquareX, maxBlankSquareY, maxSize := -1, -1, 0\n\tfor i := 0; i < rows; i++ {\n\t\tfor t := 0; t < cols; t++ {\n\t\t\tcurBlankSquareSize := getBlankSquareSize(maxWidth, maxHeight, i, t)\n\t\t\tif curBlankSquareSize > maxSize {\n\t\t\t\tmaxBlankSquareX = i\n\t\t\t\tmaxBlankSquareY = t\n\t\t\t\tmaxSize = curBlankSquareSize\n\t\t\t}\n\t\t}\n\t}\n\treturn maxBlankSquareX, maxBlankSquareY, maxSize\n}\n\nfunc getMaxWidthAndMaxHeight(matrix [][]int) ([][]int, [][]int) {\n\trows, cols := getRowsAndCols(matrix)\n\tmaxWidth := get2DSlice(rows, cols)\n\tmaxHeight := get2DSlice(rows, cols)\n\tmaxWidth[0][0] = matrix[0][0] ^ 1\n\tmaxHeight[0][0] = matrix[0][0] ^ 1\n\tfor i := 1; i < rows; i++ {\n\t\tmaxWidth[i][0] = matrix[i][0] ^ 1\n\t\tif matrix[i][0] == 0 {\n\t\t\tmaxHeight[i][0] = maxHeight[i-1][0] + 1\n\t\t}\n\t}\n\tfor t := 1; t < cols; t++ {\n\t\tmaxHeight[0][t] = matrix[0][t] ^ 1\n\t\tif matrix[0][t] == 0 {\n\t\t\tmaxWidth[0][t] = maxWidth[0][t-1] + 1\n\t\t}\n\t}\n\tfor i := 1; i < rows; i++ {\n\t\tfor t := 1; t < cols; t++ {\n\t\t\tif matrix[i][t] == 0 {\n\t\t\t\tmaxHeight[i][t] = maxHeight[i-1][t] + 1\n\t\t\t\tmaxWidth[i][t] = maxWidth[i][t-1] + 1\n\t\t\t}\n\t\t}\n\t}\n\treturn maxWidth, maxHeight\n}\n\nfunc getBlankSquareSize(maxWidth, maxHeight [][]int, x, y int) int {\n\tmaxSize := 0\n\tfor size := 1; size <= min(maxHeight[x][y], maxWidth[x][y]); size++ {\n\t\tif maxHeight[x][y-size+1] >= size && maxWidth[x-size+1][y] >= size {\n\t\t\tmaxSize = size\n\t\t}\n\t}\n\treturn maxSize\n}\n\nfunc getRowsAndCols(matrix [][]int) (int, int) {\n\tif len(matrix) == 0 {\n\t\treturn 0, 0\n\t}\n\treturn len(matrix), len(matrix[0])\n}\n\nfunc get2DSlice(rows, column int) [][]int {\n\tslice := make([][]int, rows)\n\tfor i := 0; i < len(slice); i++ {\n\t\tslice[i] = make([]int, column)\n\t}\n\treturn slice\n}\n\nfunc min(arr ...int) int {\n\tif len(arr) == 1 {\n\t\treturn arr[0]\n\t}\n\ta, b := arr[0], min(arr[1:]...)\n\tif a > b {\n\t\treturn b\n\t}\n\treturn a\n}"
    }`

	var unit model.Unit
	if err := json.Unmarshal([]byte(unitJson),&unit);err!=nil{
		logs.Error(err)
		return
	}
	logs.Debug(GetContent(`C:\Users\hasee\Desktop\arc\http_server\template\unit_show_template`,unit))
}
