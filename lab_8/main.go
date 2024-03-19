package main

import "fmt"

type LFSR struct {
	init  uint32
	polyn []uint32
}

func (l *LFSR) Len() int      { return len(l.polyn) }
func (l *LFSR) Grade() uint32 { return l.polyn[0] }

func LSFR_init(init uint32, polyn []uint32) *LFSR { return &LFSR{init: init, polyn: polyn} }

func (l *LFSR) next() uint32 {
	current := l.init
	for i := 1; i < l.Len(); i++ {
		in_cicle := l.init >> l.polyn[i]
		current = current ^ in_cicle
	}
	current &= 1
	new := l.init>>1 | current<<(l.Grade()-1)
	l.init = new
	return l.init & 1
}

func main() {
	var init uint32 = 0b00000100
	// var s uint32 = 0b00000100

	polyn := []uint32{8, 0, 2, 3, 4}
	lfsr := LSFR_init(init, polyn)
	for {
		fmt.Print(lfsr.next())
	}
	// b := ((s >> 0) ^ (s >> 2) ^ (s >> 3) ^ (s >> 4)) & 1
	// fmt.Println((s >> 0), (s >> 2), (s >> 3), (s >> 4))
	// fmt.Println(b)

	// k := (s >> 1) | (b << 7)
	// new := s & 1
	// fmt.Println(k, new)
}
