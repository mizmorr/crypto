package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sync"
)

var (
	mutex = sync.RWMutex{}
)

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

func factorized(n int) []int {
	factorized := []int{}
	for i := 2; i <= n; i++ {

		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			factorized = append(factorized, i)
		}
	}

	return factorized
}

func binarize_pow(a, b, p int) int {
	return int(math.Pow(float64(a), float64(b))) % p
}

func prime_roots2(p int) {
	file, err := os.Create("result.txt")
	roots := []int{}
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	file.WriteString(fmt.Sprintln("Исходное значение - ", p))

	phi, factorized := p-1, factorized(p-1)
	file.WriteString(fmt.Sprintln("Значение функции Эйлера - ", phi))
	file.WriteString(fmt.Sprintln("Разложение - ", factorized))

	var wg sync.WaitGroup
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for g := 1; g <= p/2; g++ {
			mutex.Lock()
			file.WriteString(fmt.Sprintln("Текущее значение - ", g))
			ok := true
			file.WriteString("Проверка.. \n")
			mutex.Unlock()

			for _, v := range factorized {
				res := binarize_pow(g, phi/v, p)
				file.WriteString(fmt.Sprintln(g, "^", phi/v, " = ", math.Pow(float64(g), float64(phi/v)), "mod", p, "=", res))
				if res == 1 {
					ok = false
					mutex.Lock()
					file.WriteString("Не первообразный корень\n----------------------\n")
					mutex.Unlock()

					break
				}
			}
			if ok {
				mutex.Lock()
				roots = append(roots, g)
				file.WriteString("Первообразный корень\n----------------------\n")
				mutex.Unlock()

			}
		}

	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for g := p/2 + 1; g <= p; g++ {
			mutex.Lock()
			file.WriteString(fmt.Sprintln("Текущее значение - ", g))
			ok := true
			file.WriteString("Проверка.. \n")
			mutex.Unlock()

			for _, v := range factorized {
				res := binarize_pow(g, phi/v, p)
				file.WriteString(fmt.Sprintln(g, "^", phi/v, " = ", math.Pow(float64(g), float64(phi/v)), "mod", p, "=", res))
				if res == 1 {
					ok = false
					mutex.Lock()
					file.WriteString("Не первообразный корень\n----------------------\n")
					mutex.Unlock()

					break
				}
			}
			if ok {
				mutex.Lock()
				roots = append(roots, g)
				file.WriteString("Первообразный корень\n----------------------\n")
				mutex.Unlock()

			}
		}

	}(&wg)
	wg.Wait()
	file.WriteString(fmt.Sprintln("Результат - ", roots))

}
func prime_roots(p int) []int {

	file, err := os.Create("result.txt")
	roots := []int{}
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	file.WriteString(fmt.Sprintln("Исходное значение - ", p))

	phi, factorized := p-1, factorized(p-1)
	file.WriteString(fmt.Sprintln("Значение функции Эйлера - ", phi))
	file.WriteString(fmt.Sprintln("Разложение - ", factorized))
	for g := 1; g <= p; g++ {
		file.WriteString(fmt.Sprintln("Текущее значение - ", g))
		ok := true
		file.WriteString("Проверка.. \n")
		for _, v := range factorized {
			res := binarize_pow(g, phi/v, p)
			file.WriteString(fmt.Sprintln(g, "^", phi/v, " = ", math.Pow(float64(g), float64(phi/v)), "mod", p, "=", res))
			if res == 1 {
				ok = false
				file.WriteString("Не первообразный корень\n----------------------\n")
				break
			}
		}
		if ok {
			roots = append(roots, g)
			file.WriteString("Первообразный корень\n----------------------\n")
		}
	}

	file.WriteString(fmt.Sprintln("Результат - ", roots))
	return roots
}

func main() {
	prime_roots2(23)
}
