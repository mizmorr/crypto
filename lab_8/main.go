package main

import (
	"fmt"
	"time"
)

func main() {
	// var init_8 uint32 = 0b00000100

	// polyn_8 := []uint32{8, 0, 2, 3, 4}
	// var init_10 uint32 = 0b0100010100

	// polyn_10 := []uint32{10, 0, 3}

	// lfsr := LSFR_init(init_10, polyn_10)
	t := time.Now()

	MT19937(100000, 1)
	// i := 0
	// for i < 100000 {
	// 	// fmt.Print(lfsr.next())
	// 	lfsr.next()
	// 	i++
	// }
	fmt.Println()
	fmt.Println(time.Since(t).Minutes())
	// b := ((s >> 0) ^ (s >> 2) ^ (s >> 3) ^ (s >> 4)) & 1
	// fmt.Println((s >> 0), (s >> 2), (s >> 3), (s >> 4))
	// fmt.Println(b)

	// k := (s >> 1) | (b << 7)
	// new := s & 1
	// fmt.Println(k, new)
}
