package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))

}

func Sqrt(x float64) (float64 error) {
	if x < 0 {
		return ErrNegativeSqrt(x)
	}
	return nil

}

func main() {

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}