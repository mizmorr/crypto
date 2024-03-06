package main

import "math"

func legendre(a, p int) int {

	if a%p == 0 {
		return 0
	}
	if int(math.Pow(float64(a), float64(p/2)))%p == 1 {
		return 1
	}
	return -1
}

func is_prime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func quadratic_comparisons(a, b, m int) (results []int) {

	if legendre(a, m) == 1 {
		if is_prime(m) {
			temp_a := int(math.Pow(float64(a), float64((m+1)/4))) % m
			x_1 := temp_a * b % m
			results = append(results, x_1)
			results = append(results, m-x_1)
			return
		} else {
			if len(factorized(int64(m))) != 1 {

			} else {

			}
			for j := 0; j < m; j++ {
				if (a*j*j)%m == b {
					results = append(results, j)
				}
			}
			return
		}
	}
	return
}