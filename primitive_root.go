package main

func gcd(a, b int) int {

	if a == b {
		return a
	} else {
		if a > b {
			return gcd(a-b, b)
		} else {
			return gcd(a, b-a)
		}
	}

}

func phi(n int) int {
	var result float64 = float64(n)

	for i := 2; i*i <= n; i++ {

		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			result *= (1.0 - (1.0 / float64(i)))
		}
	}
	if n > 1 {
		result -= result / float64(n)
	}
	return int(result)
}
func main() {

}
