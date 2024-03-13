package main

import (
	"fmt"
	"math"
	"math/rand"
)

func Ferma_factorize(n int) {

	sqr := int(math.Sqrt(float64(n)))
	for k := 0; k < n/2; k++ {
		curr := (k + sqr) * (k + sqr)
		fmt.Println(sqr, "+", k, "^2 -", n, " = ", curr-n, math.Sqrt(float64(curr-n)))
	}
}

func Polard_p_1(n, B int) int {
	a := 2
	e := 2
	for e <= B {
		a = Fast_mod(a, e, n)
		e++
	}
	p := Gcd(a-1, n)
	if 1 < p && p < n {
		return p
	} else {
		return -1
	}
}

func pollard_p_2(n int) int {
	i := 2
	a := 2
	d := GCD2(a, n)

	for d == 1 {
		a = Fast_mod(a, i, n)
		d = GCD2(a-1, n)
		i++
	}

	return d
}
func P_1_factorize(n int) {

	if Is_prime(n) {
		fmt.Println(n, "is prime")
		return
	}
	divider := pollard_p_2(n)
	fmt.Print(divider, " ")
	n /= divider
	for n != 1 || Is_prime(n) {
		divider = pollard_p_2(n)
		n /= divider
		fmt.Print(divider, " ")

	}
	fmt.Println()
}
func g(x int) int { return int(math.Pow(float64(x), 2) + 1) }

func Polard_rho(n int) {
	x_1 := 2
	x_2 := x_1
	d := 1
	for d == 1 {
		x_1 = g(x_1) % n
		x_2 = g(x_1) % n
		d = Gcd(int(math.Abs(float64(x_1-x_2))), n)
	}
	if d == n {
		fmt.Println("failure")
	} else {
		fmt.Println(d)
	}
}

func pollard_rho(n int) int {

	if n == 1 {
		return n
	}

	if n%2 == 0 {
		return 2
	}

	var x_1 int = rand.Int()%(n-2) + 2
	x_2 := x_1

	var c int = rand.Int()%(n-1) + 1

	d := 1
	// fmt.Println("Сгенерированные параметры:", x_1, c)
	/* until the prime factor isn't obtained.
	   If n is prime, return n */
	for d == 1 {
		/* Tortoise Move: x(i+1) = f(x(i)) */
		x_1 = (Fast_mod(x_1, 2, n) + c + n) % n

		/* Hare Move: y(i+1) = f(f(y(i))) */
		x_2 = (Fast_mod(x_2, 2, n) + c + n) % n
		x_2 = (Fast_mod(x_2, 2, n) + c + n) % n

		/* check gcd of |x-y| and n */
		d = GCD2(int(math.Abs(float64(x_1-x_2))), n)

		/* retry if the algorithm fails to find prime factor
		 * with chosen x and c */
		if d == n {
			// fmt.Println("one more try")
			return pollard_rho(n)
		}
	}

	return d
}

func Rho_factorize(n int) {

	// result := []int{}
	if Is_prime(n) {
		fmt.Println(n)
		return
	} else {
		current := pollard_rho(n)
		n /= current
		if Is_prime(current) {
			fmt.Println(current)
			Rho_factorize(n)
		} else {
			Rho_factorize(current)
			fmt.Println(n)
		}
	}

}
