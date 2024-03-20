package main

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
