package main

import (
	"fmt"
)

func SQUARE(n int) {
	for k := 1; k < n; k++ {
		fmt.Println(k, "^2 mod =", n, Fast_mod(k, 2, n))
	}
}

func main() {
	// field := generate_field(4)
	// fmt.Println(primitivity_check([]byte{1, 0, 0, 1, 0, 1}, []byte{1, 0, 0, 0, 0, 1, 1}))
	// t := field[3]
	// mark_field(field, []byte{1, 1, 0, 1})
	// fmt.Println(field)
	// Ferma_factorize(1219)
	// Polard_rho(1219)
	// fmt.Println(Ferma_check(23))
	fmt.Println(PollardRho(1207))
}
