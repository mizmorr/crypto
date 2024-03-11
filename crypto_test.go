package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {

	args := [][]int{{15, 13}, {17, 12}, {27, 18}, {30, 10}, {40, 20}, {150, 25}}
	expected := []int{1, 1, 9, 10, 20, 25}
	for i, arg := range args {
		actual := Gcd(arg[0], arg[1])

		if expected[i] != actual {
			t.Errorf("Result is incorrect, got: %v, expected: %v", actual, expected)
		}
	}
}

// func TestEulerFunc(t *testing.T) {

// 	args := []int64{6, 8, 9, 10, 11, 12, 13, 14, 15, 16}
// 	expected := []int64{2, 4, 6, 4, 10, 4, 12, 6, 8, 8}

// 	for i, arg := range args {
// 		phi_value, _ := phi(arg)
// 		assert.Equal(t, phi_value, expected[i])
// 	}
// }

func TestBinarize(t *testing.T) {
	args := [][]int64{
		{4, 1},
		{4, 2},
		{7, 5},
		{8, 6},
	}
	module := []int64{11, 123, 15, 17}
	expected := []int64{4, 16, 7, 4}
	for i, arg := range args {
		assert.Equal(t, binarize_pow(arg[0], arg[1], module[i]), expected[i])
	}
}

func TestPrime(t *testing.T) {

	args := [][]byte{
		{1, 0, 1},
		{1, 0, 1, 0},
		{1, 0, 1, 1, 1, 1, 0},
		{1, 1, 1},
		{1, 0, 0, 1, 1},
	}
	expected := []bool{false, false, false, true, true}
	for i, arg := range args {
		assert.Equal(t, prime(arg), expected[i])
	}
}
