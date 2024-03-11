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

func factorized(n int64) []int64 {
	factorized := []int64{}
	for i := int64(2); i*i <= n; i++ {

		if n%i == 0 {
			factorized = append(factorized, i)

			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		factorized = append(factorized, n)
	}

	return factorized
}

func binarize_pow(a, b, p int64) int64 {
	return int64(math.Pow(float64(a), float64(b))) % p
}

func prime_roots2(p int64) {
	file, err := os.Create("result.txt")
	roots := []int64{}
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
		for g := int64(1); g <= p/2; g++ {
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
func prime_roots(p int64) []int64 {

	file, err := os.Create("result.txt")
	roots := []int64{}
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	file.WriteString(fmt.Sprintln("Исходное значение - ", p))

	phi, factorized := p-1, factorized(p-1)
	file.WriteString(fmt.Sprintln("Значение функции Эйлера - ", phi))
	file.WriteString(fmt.Sprintln("Разложение - ", factorized))
	for g := int64(2); g < p; g++ {
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

func pp(a, mod int) {

	for i := 2; i < mod; i++ {
		fmt.Println(a, "^", i, "mod", mod, "=", int(math.Pow(float64(a), float64(i)))%mod)
	}
}
