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

func PollardRho(n int) int {

	/* no prime divisor for 1 */
	if n == 1 {
		return n
	}

	/* even number means one of the divisors is 2 */
	if n%2 == 0 {
		return 2
	}

	var x_1 int = rand.Int()%(n-2) + 2
	x_2 := x_1

	var c int = rand.Int()%(n-1) + 1

	/* Initialize candidate divisor (or result) */
	d := 1

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
			return PollardRho(n)
		}
	}

	return d
}
