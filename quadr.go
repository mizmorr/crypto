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

func Is_prime(n int) bool {
	if n <= 1 {
		return false
	}
	if n%2 == 0 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func quadratic_comparisons(a, b, m int) (results []int) {

	if legendre(a, m) == 1 {
		if Is_prime(m) {
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
