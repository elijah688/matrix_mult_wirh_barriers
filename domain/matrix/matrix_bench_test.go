package matrix

import "testing"

func BenchmarkMiltiplyVanilla(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewRandomMatrix().Miltiply(NewRandomMatrix())
	}
}
func BenchmarkMiltiplyBarrier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewRandomMatrix().MultiplyWithBarrier(NewRandomMatrix())
	}
}
