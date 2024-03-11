package main

import (
	"fmt"
	"math"
)

func SQUARE(n int) {
	for k := 1; k < n; k++ {
		fmt.Println(k, "^2 mod =", n, Fast_mod(k, 2, n))
	}
}

func ferma_factorize(n int) {

	sqr := int(math.Sqrt(float64(n)))
	for k := 0; k < n/2; k++ {
		curr := (k + sqr) * (k + sqr)
		fmt.Println(sqr, "+", k, "^2 -", n, " = ", curr-n, math.Sqrt(float64(curr-n)))
	}
}
func main() {
	// field := generate_field(4)
	// fmt.Println(primitivity_check([]byte{1, 0, 0, 1, 0, 1}, []byte{1, 0, 0, 0, 0, 1, 1}))
	// t := field[3]
	// mark_field(field, []byte{1, 1, 0, 1})
	// fmt.Println(field)
	ferma_factorize(1219)
	// fmt.Println(Ferma_check(23))
}
