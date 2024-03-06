package main

import (
	"bytes"
	"fmt"
	"math"
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

func prime(a []byte) bool {
	if a[len(a)-1] == 0 {
		return false
	}
	sum := byte(0)
	for _, k := range a {
		sum += k
	}
	if sum%2 == 0 {
		return false
	}
	if len(a) > 3 {
		if bytes.Equal(divider(a, []byte{1, 1, 1}), []byte{}) {
			return false
		}
	}
	if len(a) > 5 {
		if bytes.Equal(divider(a, []byte{1, 1, 0, 1}), []byte{}) {
			return false
		}
		if bytes.Equal(divider(a, []byte{1, 0, 1, 1}), []byte{}) {
			return false
		}
	}
	if len(a) > 7 {
		if bytes.Equal(divider(a, []byte{1, 1, 1, 1, 1}), []byte{}) {
			return false
		}
		if bytes.Equal(divider(a, []byte{1, 1, 0, 0, 1}), []byte{}) {
			return false
		}
		if bytes.Equal(divider(a, []byte{1, 1, 0, 0, 1}), []byte{}) {
			return false
		}
	}

	return true
}
func string_to_bytes(s string) (bytes []byte) {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error:", err)
		return []byte{}
	}
	for num > 0 {
		bytes = append(bytes, byte(num%10))
		num /= 10
	}
	return
}

func generate_field(n int) (res [][]byte) {
	for i := int64(1); i < int64(math.Pow(2, float64(n))); i++ {
		x_hex := strconv.FormatInt(i, 2)
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

func find(array [][]byte, search []byte) bool {

	go func() bool{
		for i:=0;i<len(array)/2;i++ {
			if bytes.Equal(search,array[i]){
				return true
			}
	}

	return false

}
func get_primitives(field [][]byte, polynom []byte, degree int) {

	for i, subject := range field {
		counter := 1
		// fmt.Println(subject, prime(subject))
		if prime(subject) {
			for j := 1; j < degree; j++ {
				subject = divider(power(subject, degree), polynom)
				for k, subj_second := range field {
					if i != k {
						if bytes.Equal(subject, subj_second) {
							counter++
							break
						}
					}
				}
			}
		}
		if counter == degree {
			fmt.Println(subject)
		}
	}
}
func main() {

	var gf_8 = generate_field(3)

	get_primitives(gf_8, []byte{1, 0, 1, 1}, 3)
}
