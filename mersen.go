package main

import (
	"fmt"
	"math"
)

func get_mersenn_number(n int) int { return int(math.Pow(2, float64(n)) - 1) }

func Check_Mnumb(n int) {

	M_n := get_mersenn_number(n)
	fmt.Println(Is_prime(M_n), " - primitive algorithm")
	fmt.Println("check dividers:")
	Rho_factorize(M_n)
}
