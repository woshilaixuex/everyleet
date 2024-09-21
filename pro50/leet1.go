package leetcodego

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: https://leetcode.cn/problems/spiral-matrix/?envType=study-plan-v2&envId=2024-spring-sprint-100 54.螺旋矩阵
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,3,6,9,8,7,4,5]
 * @Date: 2024-09-21 22:38
 */
func spiralOrder(matrix [][]int) []int { // 个人题解
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	ans := make([]int, 0, m*n)
	matrix_bool := make([][]bool, m)
	for i := range matrix_bool {
		matrix_bool[i] = make([]bool, n)
	}

	x := []int{0, 1, 0, -1}
	y := []int{1, 0, -1, 0}
	i, j, state := 0, 0, 0

	for len(ans) < m*n {
		if i >= 0 && i < m && j >= 0 && j < n && !matrix_bool[i][j] {
			ans = append(ans, matrix[i][j])
			matrix_bool[i][j] = true
		} else {
			i -= x[state%4]
			j -= y[state%4]
			state++
		}
		i += x[state%4]
		j += y[state%4]
	}

	return ans
}
func spiralOrder_(matrix [][]int) []int { // 官方题解
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		total          = rows * columns
		order          = make([]int, total)
		row, column    = 0, 0
		directions     = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		directionIndex = 0
	)

	for i := 0; i < total; i++ {
		order[i] = matrix[row][column]
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return order
}
