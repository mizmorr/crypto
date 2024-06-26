package main

import (
	"fmt"

	"gopkg.in/dedis/kyber.v2"
	"gopkg.in/dedis/kyber.v2/group/edwards25519"
)

var curve = edwards25519.NewBlakeSHA256Ed25519()
var sha256 = curve.Hash()

type Signature struct {
	r kyber.Point
	s kyber.Scalar
}

func Hash(s string) kyber.Scalar {
	sha256.Reset()
	sha256.Write([]byte(s))

	return curve.Scalar().SetBytes(sha256.Sum(nil))
}

// m: Message
// x: Private key
func Sign(m string, x kyber.Scalar) Signature {
	// Get the base of the curve.
	g := curve.Point().Base()

	// Pick a random k from allowed set.
	k := curve.Scalar().Pick(curve.RandomStream())

	// r = k * G (a.k.a the same operation as r = g^k)
	r := curve.Point().Mul(k, g)

	// Hash(m || r)
	e := Hash(m + r.String())

	// s = k - e * x
	s := curve.Scalar().Sub(k, curve.Scalar().Mul(e, x))

	return Signature{r: r, s: s}
}

// m: Message
// S: Signature
func PublicKey(m string, S Signature) kyber.Point {
	// Create a generator.
	g := curve.Point().Base()

	// e = Hash(m || r)
	e := Hash(m + S.r.String())

	// y = (r - s * G) * (1 / e)
	y := curve.Point().Sub(S.r, curve.Point().Mul(S.s, g))
	y = curve.Point().Mul(curve.Scalar().Div(curve.Scalar().One(), e), y)

	return y
}

// m: Message
// s: Signature
// y: Public key
func Verify(m string, S Signature, y kyber.Point) bool {
	// Create a generator.
	g := curve.Point().Base()

	// e = Hash(m || r)
	e := Hash(m + S.r.String())

	// Attempt to reconstruct 's * G' with a provided signature; s * G = r - e * y
	sGv := curve.Point().Sub(S.r, curve.Point().Mul(e, y))

	// Construct the actual 's * G'
	sG := curve.Point().Mul(S.s, g)

	// Equality check; ensure signature and public key outputs to s * G.
	return sG.Equal(sGv)
}

func (S Signature) String() string {
	return fmt.Sprintf("(r=%s, s=%s)", S.r, S.s)
}

func KeyPair() (privateKey kyber.Scalar, publicKey kyber.Point) {
	privateKey = curve.Scalar().Pick(curve.RandomStream())
	publicKey = curve.Point().Mul(privateKey, curve.Point().Base())

	return
}

func main() {
	//privateKey := curve.Scalar().Pick(curve.RandomStream())
	//publicKey := curve.Point().Mul(privateKey, curve.Point().Base())
	privateKey, publicKey := KeyPair()

	fmt.Printf("Generated private key: %s\n", privateKey)
	fmt.Printf("Derived public key: %s\n\n", publicKey)

	message := "We're gonna be signing this!"

	signature := Sign(message, privateKey)
	fmt.Printf("Signature %s\n\n", signature)

	derivedPublicKey := PublicKey(message, signature)
	fmt.Printf("Derived public key: %s\n", derivedPublicKey)
	fmt.Printf("Are the original and derived public keys the same? %t\n", publicKey.Equal(derivedPublicKey))
	fmt.Printf("Is the signature legit w.r.t the original public key? %t\n\n", Verify(message, signature, publicKey))

	fakePublicKey := curve.Point().Mul(curve.Scalar().Neg(curve.Scalar().One()), publicKey)
	fmt.Printf("Fake public key: %s\n", fakePublicKey)
	fmt.Printf("Is the signature legit w.r.t a fake public key? %t\n", Verify(message, signature, fakePublicKey))
}
