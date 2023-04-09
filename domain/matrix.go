package matrix

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	MAX_SQUARE_MATRIX_DIMENSION = 3
	MAX_MATRIX_VAL              = 10
)

type Matrix [][]int

func NewZeroMatrix() Matrix {
	matrix := make([][]int, MAX_SQUARE_MATRIX_DIMENSION)
	for i := range matrix {
		matrix[i] = make([]int, MAX_SQUARE_MATRIX_DIMENSION)
	}
	return matrix
}

func NewRandomMatrix() Matrix {
	matrix := NewZeroMatrix()
	rand.Seed(time.Now().UnixNano())

	for i := range matrix {
		for j := range matrix[i] {
			num := rand.Intn(MAX_MATRIX_VAL)
			if rand.Intn(2) == 0 {
				matrix[i][j] = -num
			} else {
				matrix[i][j] = num
			}
		}
	}

	return matrix
}

func (m Matrix) String() string {
	sb := new(strings.Builder)

	for i := range m {
		fmt.Println(i)
		for j := range m[i] {
			sb.WriteString(fmt.Sprintf("%4d", m[i][j]))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (m Matrix) Multiply(otherMatrix Matrix) Matrix {
	// 1 2 3            // 1 4 7
	// 4 5 6            // 2 5 8
	// 7 8 9            // 3 6 9

	res := NewZeroMatrix()

	for rowIdx, row := range m {
		// multiply row with other matrix's columns
		for otherColIdx := 0; otherColIdx < MAX_SQUARE_MATRIX_DIMENSION; otherColIdx++ {
			num := 0
			for otherRowIdx := 0; otherRowIdx < MAX_SQUARE_MATRIX_DIMENSION; otherRowIdx++ {
				num += row[otherRowIdx] * otherMatrix[otherRowIdx][otherColIdx]
			}
			res[rowIdx][otherColIdx] = num
		}
	}

	return res

}
