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
	MAX_SQUARE_MATRIX_DIMENSION = 100
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

func NewSquareZeroMatrix(dimension int) Matrix {
	matrix := make([]Row, dimension)
	for i := range matrix {
		matrix[i] = make([]int, dimension)
	}
	return matrix
}

func NewSquareRandomMatrix(dimension int) Matrix {
	matrix := NewSquareZeroMatrix(dimension)
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
		for j := range m[i] {
			sb.WriteString(fmt.Sprintf("%4d", m[i][j]))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

type Row []int

func (r Row) MultiplyWithMatrix(rowIdx int, otherMatrix Matrix, output Matrix) {
	barrier := barrier.NewBarrier(len(r))
	wg := new(sync.WaitGroup)
	for otherColIdx := 0; otherColIdx < len(r); otherColIdx++ {
		wg.Add(1)
		go func(otherColIdx int) {
			barrier.Wait()
			num := 0
			for otherRowIdx := 0; otherRowIdx < len(r); otherRowIdx++ {
				num += r[otherRowIdx] * otherMatrix[otherRowIdx][otherColIdx]
			}
			output[rowIdx][otherColIdx] = num
			wg.Done()
		}(otherColIdx)
	}
	wg.Wait()
}

func (m Matrix) MultiplyWithBarrier(otherMatrix Matrix) Matrix {
	barrier := barrier.NewBarrier(len(m))
	res := NewSquareZeroMatrix(len(m))

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
