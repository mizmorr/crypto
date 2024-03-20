package main

import "fmt"

var (
	w         uint32 = 32
	n                = 624
	m                = 397
	a         uint32 = 0x9908B0DF
	u         uint32 = 11
	d         uint32 = 0xFFFFFFFF
	s         uint32 = 7
	b         uint32 = 0x9D2C5680
	t         uint32 = 15
	c         uint32 = 0xEFC60000
	l                = 18
	f                = uint32(1812433253)
	MT               = make([]uint32, n)
	index            = n + 1
	lowerMask        = uint32(0x7FFFFFFF)
	upperMask        = uint32(0x80000000)
)

func mtSeed(seed uint32) {
	MT[0] = seed
	for i := 1; i < n; i++ {
		temp := f*(MT[i-1]^(MT[i-1]>>(w-2))) + uint32(i)
		MT[i] = temp & 0xffffffff
	}
}

func extractNumber() uint32 {
	if index >= n {
		twist()
		index = 0
	}
	y := MT[index]
	y = y ^ ((y >> u) & d)
	y = y ^ ((y << s) & b)
	y = y ^ ((y << t) & c)
	y = y ^ (y >> l)
	index++
	return y & 0xffffffff
}

func twist() {
	for i := 0; i < n; i++ {
		x := (MT[i] & upperMask) + (MT[(i+1)%n] & lowerMask)
		xA := x >> 1
		if x%2 != 0 {
			xA = xA ^ a
		}
		MT[i] = MT[(i+m)%n] ^ xA
	}
}

func MT19937(n int, a uint32) {
	mtSeed(a)
	for i := 0; i < n; i++ {

		fmt.Println(extractNumber())
	}
}
