package main

import (
	"fmt"
	"matrix_mult/domain/barrier"
	"matrix_mult/domain/matrix"
	"time"
)

const (
	MAX_SQUARE_MATRIX_DIMENSION = 1000
	BARRIER_SIZE                = 2
)

func main() {
	for _, m := range []struct {
		title  string
		option int
	}{
		{
			title:  "vanilla",
			option: 0,
		},
		{
			title:  "barrier",
			option: 1,
		},
	} {

		start := time.Now()
		var (
			workStart    = barrier.NewBarrier(BARRIER_SIZE)
			workComplete = barrier.NewBarrier(BARRIER_SIZE)
		)

		a := matrix.NewSquareZeroMatrix(MAX_SQUARE_MATRIX_DIMENSION)
		b := matrix.NewSquareZeroMatrix(MAX_SQUARE_MATRIX_DIMENSION)

		go func() {
			for {

				a = matrix.NewSquareRandomMatrix(MAX_SQUARE_MATRIX_DIMENSION)
				b = matrix.NewSquareRandomMatrix(MAX_SQUARE_MATRIX_DIMENSION)

				workStart.Wait()
				workComplete.Wait()
			}
		}()

		for i := 0; i < 10; i++ {

			workStart.Wait()

			var res matrix.Matrix
			if m.option == 0 {
				res = a.Multiply(b)
			} else {
				res = a.MultiplyWithBarrier(b)
			}

			if res.IsZero() {
				panic("matrix product should not be zero matrix")
			}

			workComplete.Wait()
		}

		elapsed := time.Since(start)

		fmt.Println(fmt.Sprintf("%s took %.3v", m.title, elapsed.Seconds()))

	}

}
