package matrix

import "testing"

func BenchmarkMiltiplyVanilla(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSquareRandomMatrix(MAX_SQUARE_MATRIX_DIMENSION).Multiply(NewSquareRandomMatrix(MAX_SQUARE_MATRIX_DIMENSION))
	}
}
func BenchmarkMiltiplyBarrier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSquareRandomMatrix(MAX_SQUARE_MATRIX_DIMENSION).MultiplyWithBarrier(NewSquareRandomMatrix(MAX_SQUARE_MATRIX_DIMENSION))
	}
}
