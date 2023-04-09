package matrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandomMatrix(t *testing.T) {

	for i := 0; i < 10; i++ {
		totalInitialZeros := MAX_SQUARE_MATRIX_DIMENSION * MAX_SQUARE_MATRIX_DIMENSION
		countedZeros := 0
		m := NewSquareRandomMatrix(MAX_SQUARE_MATRIX_DIMENSION)

		for i, row := range m {
			for j := range row {
				if m[i][j] == 0 {
					countedZeros++
				}
			}
		}

		if !assert.NotEqual(t, countedZeros, totalInitialZeros) {
			t.Fatalf("all matrix should have distinct values")
		}
	}
}

func TestMult(t *testing.T) {
	m := Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	otherMartrix := Matrix{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}

	expected := Matrix{
		{14, 32, 50},
		{32, 77, 122},
		{50, 122, 194},
	}

	t.Run("multiply vanilla", func(t *testing.T) {

		actual := m.Multiply(otherMartrix)
		if !assert.Equal(t, actual, expected) {
			t.Fatalf(fmt.Sprintf("\n%s is not the correct result when multiplying \n%s with \n%s", actual, m.String(), otherMartrix.String()))
		}
	})

	t.Run("multiply with barrier", func(t *testing.T) {

		actual := m.MultiplyWithBarrier(otherMartrix)
		if !assert.Equal(t, actual, expected) {
			t.Fatalf(fmt.Sprintf("\n%s is not the correct result when multiplying \n%s with \n%s", actual, m.String(), otherMartrix.String()))
		}
	})
}
