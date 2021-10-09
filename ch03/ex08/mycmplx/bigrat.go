package mycmplx

import (
	"math"
	"math/big"
)

type BigRat struct {
	Real *big.Rat
	Imag *big.Rat
}

func NewBigRat(r, i float64) *BigRat {
	return &BigRat{Real: new(big.Rat).SetFloat64(r), Imag: new(big.Rat).SetFloat64(i)}
}

// α + β = (a + ib) + (c + id)
func (alpha *BigRat) Add(beta *BigRat) *BigRat {
	a, b := alpha.Real, alpha.Imag
	c, d := beta.Real, beta.Imag

	r := new(big.Rat).Add(a, c)
	i := new(big.Rat).Add(b, d)

	return &BigRat{r, i}
}

// αβ = (a + ib)(c + id)
func (alpha *BigRat) Mul(beta *BigRat) *BigRat {
	a, b := alpha.Real, alpha.Imag
	c, d := beta.Real, beta.Imag

	ac := new(big.Rat).Mul(a, c)
	bd := new(big.Rat).Mul(b, d)
	ad := new(big.Rat).Mul(a, d)
	bc := new(big.Rat).Mul(b, c)

	r := new(big.Rat).Sub(ac, bd)
	i := new(big.Rat).Add(ad, bc)

	return &BigRat{r, i}
}

func (z *BigRat) Abs() float64 {
	r, _ := z.Real.Float64()
	i, _ := z.Imag.Float64()

	return math.Hypot(r, i)
}

func (z *BigRat) ToComplex128() complex128 {
	r, _ := z.Real.Float64()
	i, _ := z.Imag.Float64()
	return complex(r, i)
}
