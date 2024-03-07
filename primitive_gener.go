package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func divider(a, b []byte) []byte {
	for len(a) >= len(b) {
		temp := []byte{}
		for i := 0; i < len(a); i++ {

			if len(b)-i-1 >= 0 {
				temp = append(temp, (a[i]+b[i])%2)
			} else {
				temp = append(temp, a[i])
			}

		}
		for temp[0] != 1 {
			temp = temp[1:]
			if len(temp) == 0 {
				break
			}
		}
		a = temp
	}
	return a
}

func string_to_bytes(s string) (b []byte) {

	for k := 0; k < len(s); k++ {
		num, _ := strconv.Atoi(string(s[k]))
		b = append(b, byte(num))
	}
	return
}

func generate_field(n int) (res [][]byte) {
	for i := int64(1); i < int64(math.Pow(2, float64(n))); i++ {
		x_hex := strconv.FormatInt(i, 2)
		// if i > 3 && strings.TrimFunc(x_hex, func(a rune) bool { return a == '0' }) == "1" {
		res = append(res, string_to_bytes(x_hex))

	}
	return
}

func unpack_power(a, res []byte, degree int, number int) []byte {

	if degree == number {
		return res
	} else {
		return unpack_power(a, multiply(a, res), degree, number+1)
	}
}

func power(a []byte, degree int) []byte {
	if degree == 0 {
		return []byte{1}
	}
	return unpack_power(a, a, degree, 1)
}

func multiply(a, b []byte) (res []byte) {
	res = make([]byte, len(a)+len(b)-1)
	for i := range a {
		for j := range b {
			res[i+j] = (res[i+j] + a[i]*b[j]) % 2
		}
	}
	return
}

// func interpolation_search(arr [][]byte, low, high int, search []byte) bool {

// 	if low <= high && bytes.Compare(search, arr[low]) >= 0 && bytes.Compare(search, arr[high]) <= 0 {

// 		if bytes.Compare(arr[high], arr[low]) == 0 {
// 			switch {
// 			case bytes.Compare(arr[len(arr)-1], search) == 0:
// 				return true
// 			default:
// 				return false
// 			}
// 		}
// 		pos := low + (((high - low) / (arr[high] - arr[low])) * (search - arr[low]))
// 		switch {
// 		case arr[pos] == search:
// 			return pos
// 		case arr[pos] < search:
// 			return interpolation_search(arr, pos+1, high, search)
// 		case arr[pos] > search:
// 			return interpolation_search(arr, low, pos-1, search)
// 		}
// 	}
// 	return -1
// }

func find(array [][]byte, search []byte) (bool, int) {
	for index, elem := range array {
		if bytes.Equal(elem, search) {
			return true, index
		}
	}

	return false, -1

}
func mark_field(field [][]byte, polynom []byte) {

	file, err := os.Create("primitive_result.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	file.WriteString(fmt.Sprint("Markup GF(", len(field)+1, ")", "\n"))

	for _, subject := range field {
		counter := 0
		file.WriteString(fmt.Sprintln("subject: ", subject))
		var result string

		if prime(subject) {
			current := subject
			result = "prime, "
			for j := 0; j < len(field); j++ {
				current = divider(power(subject, j), polynom)
				if ok, index := find(field, current); ok {
					counter++
					file.WriteString(fmt.Sprintln(subject, "^", j, "=", current, "- ", index, " number in field"))
				}
				//test
				if j > 4 && counter < 3 {
					break
				}
			}
		} else {
			result = "not prime so - "
		}
		if counter == len(field) {
			result = "primitive"
		} else {
			result += "non primitive"
		}
		file.WriteString(fmt.Sprint("Result: ", result, "\n", "-------------", "\n", "\n"))
	}
}

func get_irr(degree int) (res [][]byte) {

	readFile, err := os.Open("irreducible.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() && len(fileScanner.Text())-1 <= degree/2 {
		res = append(res, string_to_bytes(fileScanner.Text()))
	}
	readFile.Close()
	return
}

func prime(polyn []byte) bool {
	if bytes.Equal(polyn, []byte{1}) {
		return false
	}
	irreducible_p := get_irr(len(polyn) - 1)
	for _, tested := range irreducible_p {
		if bytes.Equal(divider(polyn, tested), []byte{}) {
			return false
		}
	}
	return true
}

func main() {
	field := generate_field(4)
	// t := field[3]
	mark_field(field, []byte{1, 1, 0, 1})
}
