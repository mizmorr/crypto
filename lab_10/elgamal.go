package elgamal

import (
	"crypto/rand"
	"crypto/subtle"
	"errors"
	"io"
	"math/big"
)

type PublicKey struct {
	G, P, Y *big.Int
}

type PrivateKey struct {
	PublicKey
	X *big.Int
}

func Encrypt(random io.Reader, pub *PublicKey, msg []byte) (a, b *big.Int, err error) {
	pLen := (pub.P.BitLen() + 7) / 8
	if len(msg) > pLen-11 {
		err = errors.New("elgamal: message too long")
		return
	}

	// EM = 0x02 || PS || 0x00 || M
	em := make([]byte, pLen-1)
	em[0] = 2
	ps, mm := em[1:len(em)-len(msg)-1], em[len(em)-len(msg):]
	err = nonZeroRandomBytes(ps, random)
	if err != nil {
		return
	}
	em[len(em)-len(msg)-1] = 0
	copy(mm, msg)

	m := new(big.Int).SetBytes(em)

	k, err := rand.Int(random, pub.P)
	if err != nil {
		return
	}

	a = new(big.Int).Exp(pub.G, k, pub.P)
	s := new(big.Int).Exp(pub.Y, k, pub.P)
	b = s.Mul(s, m)
	b.Mod(b, pub.P)

	return
}

func Decrypt(priv *PrivateKey, c1, c2 *big.Int) (msg []byte, err error) {
	s := new(big.Int).Exp(c1, priv.X, priv.P)
	if s.ModInverse(s, priv.P) == nil {
		return nil, errors.New("elgamal: invalid private key")
	}
	s.Mul(s, c2)
	s.Mod(s, priv.P)
	em := s.Bytes()

	firstByteIsTwo := subtle.ConstantTimeByteEq(em[0], 2)

	var lookingForIndex, index int
	lookingForIndex = 1

	for i := 1; i < len(em); i++ {
		equals0 := subtle.ConstantTimeByteEq(em[i], 0)
		index = subtle.ConstantTimeSelect(lookingForIndex&equals0, i, index)
		lookingForIndex = subtle.ConstantTimeSelect(equals0, 0, lookingForIndex)
	}

	if firstByteIsTwo != 1 || lookingForIndex != 0 || index < 9 {
		return nil, errors.New("elgamal: decryption error")
	}

	return em[index+1:], nil
}

func nonZeroRandomBytes(s []byte, rand io.Reader) (err error) {
	_, err = io.ReadFull(rand, s)
	if err != nil {
		return
	}

	for i := 0; i < len(s); i++ {
		for s[i] == 0 {
			_, err = io.ReadFull(rand, s[i:i+1])
			if err != nil {
				return
			}
		}
	}

	return
}
