package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func get_primes() (res []int) {

	readFile, err := os.Open("prime_num.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		curr, err := strconv.Atoi(strings.Trim((fileScanner.Text()), " "))
		if err != nil {
			log.Fatal(err)
		}

		res = append(res, curr)
	}
	readFile.Close()
	return
}
func Fermat_check(n int) bool {

	primes := get_primes()
	for _, prime := range primes {
		if prime != n && Fast_mod(prime, n-1, n) != 1 {
			return false
		}
	}
	return true
}

// need to use fermat check

func Task14() {
	args := []int{100, 110, 130, 150, 200, 250, 271, 341, 561}
	for _, n := range args {
		fmt.Println(n, "prime:", Fermat_check(n))
	}
}
