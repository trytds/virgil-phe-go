package swu

import (
	"math/big"
)

// GF represents galois field over prime
type GF struct {
	P *big.Int
}

var (
	one   = big.NewInt(1)
	two   = big.NewInt(2)
	three = big.NewInt(3)
	four  = big.NewInt(4)
)

func (g *GF) Neg(a *big.Int) *big.Int {
	return new(big.Int).Sub(g.P, a)
}

func (g *GF) NegBytes(a []byte) *big.Int {
	return new(big.Int).Sub(g.P, new(big.Int).SetBytes(a))
}

func (g *GF) Square(a *big.Int) *big.Int {
	return new(big.Int).Exp(a, two, g.P)
}

func (g *GF) Cube(a *big.Int) *big.Int {
	return new(big.Int).Exp(a, three, g.P)
}

func (g *GF) Pow(a, b *big.Int) *big.Int {
	return new(big.Int).Exp(a, b, g.P)
}

func (g *GF) Inv(a *big.Int) *big.Int {
	return new(big.Int).ModInverse(a, g.P)
}

func (g *GF) InvBytes(a []byte) *big.Int {
	return new(big.Int).ModInverse(new(big.Int).SetBytes(a), g.P)
}

func (g *GF) Add(a, b *big.Int) *big.Int {
	add := new(big.Int).Add(a, b)
	return add.Mod(add, g.P)
}

func (g *GF) AddBytes(a []byte, b *big.Int) *big.Int {
	add := new(big.Int).Add(new(big.Int).SetBytes(a), b)
	return add.Mod(add, g.P)
}

func (g *GF) Sub(a, b *big.Int) *big.Int {

	negB := new(big.Int).Sub(a, b)
	return negB.Mod(negB, g.P)
}

func (g *GF) Mul(a, b *big.Int) *big.Int {
	mul := new(big.Int).Mul(a, b)
	return mul.Mod(mul, g.P)
}

func (g *GF) MulBytes(a []byte, b *big.Int) *big.Int {
	mul := new(big.Int).Mul(new(big.Int).SetBytes(a), b)
	return mul.Mod(mul, g.P)
}

func (g *GF) Div(a, b *big.Int) *big.Int {
	invB := g.Inv(b)
	t := g.Mul(a, invB)
	return t
}
