package matrix

import (
	"fmt"
	"math/rand"
	"matrix_mult/domain/barrier"
	"strings"
	"sync"
	"time"
)

const (
	MAX_SQUARE_MATRIX_DIMENSION = 3
	MAX_MATRIX_VAL              = 10
)

type Row []int

func (r Row) MultiplyWithMatrix(rowIdx int, otherMatrix Matrix, output Matrix) {
	for otherColIdx := 0; otherColIdx < len(r); otherColIdx++ {
		num := 0
		for otherRowIdx := 0; otherRowIdx < len(r); otherRowIdx++ {
			num += r[otherRowIdx] * otherMatrix[otherRowIdx][otherColIdx]
		}
		output[rowIdx][otherColIdx] = num
	}
}

type Matrix []Row

func NewZeroMatrix() Matrix {
	matrix := make([]Row, MAX_SQUARE_MATRIX_DIMENSION)
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
			matrix[i][j] = rand.Intn(MAX_MATRIX_VAL) - 5
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

type Row []int

func (r Row) MultiplyWithMatrix(rowIdx int, otherMatrix Matrix, output Matrix) {
	barrier := barrier.NewBarrier(MAX_SQUARE_MATRIX_DIMENSION)
	wg := new(sync.WaitGroup)
	for otherColIdx := 0; otherColIdx < MAX_SQUARE_MATRIX_DIMENSION; otherColIdx++ {
		wg.Add(1)
		go func(otherColIdx int) {
			barrier.Wait()
			num := 0
			for otherRowIdx := 0; otherRowIdx < MAX_SQUARE_MATRIX_DIMENSION; otherRowIdx++ {
				num += r[otherRowIdx] * otherMatrix[otherRowIdx][otherColIdx]
			}
			output[rowIdx][otherColIdx] = num
			wg.Done()
		}(otherColIdx)
	}
	wg.Wait()
}

func (m Matrix) Multiply(otherMatrix Matrix) Matrix {
	// 1 2 3            // 1 4 7
	// 4 5 6            // 2 5 8
	// 7 8 9            // 3 6 9

	barrier := barrier.NewBarrier(MAX_SQUARE_MATRIX_DIMENSION)

	res := NewZeroMatrix()

	wg := new(sync.WaitGroup)
	for rowIdx, row := range m {
		wg.Add(1)
		go func(row Row, rowIdx int) {
			barrier.Wait()
			Row(row).MultiplyWithMatrix(rowIdx, otherMatrix, res)
			wg.Done()
		}(row, rowIdx)
	}

	wg.Wait()

	return res

}

func (m Matrix) Multiply(otherMatrix Matrix) Matrix {

	res := NewSquareZeroMatrix(len(m))

	for rowIdx, row := range m {
		for otherColIdx := 0; otherColIdx < len(m); otherColIdx++ {
			num := 0
			for otherRowIdx := 0; otherRowIdx < len(m); otherRowIdx++ {
				num += row[otherRowIdx] * otherMatrix[otherRowIdx][otherColIdx]
			}
			res[rowIdx][otherColIdx] = num
		}
	}

	return res

}
