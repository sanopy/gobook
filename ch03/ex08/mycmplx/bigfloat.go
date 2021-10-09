package mycmplx

import (
	"math/big"
)

type BigFloat struct {
	Real *big.Float
	Imag *big.Float
}

func NewBigFloat(r, i float64) *BigFloat {
	return &BigFloat{Real: big.NewFloat(r), Imag: big.NewFloat(i)}
}

// α + β = (a + ib) + (c + id)
func (alpha *BigFloat) Add(beta *BigFloat) *BigFloat {
	a, b := alpha.Real, alpha.Imag
	c, d := beta.Real, beta.Imag

	r := new(big.Float).Add(a, c)
	i := new(big.Float).Add(b, d)

	return &BigFloat{r, i}
}

// αβ = (a + ib)(c + id)
func (alpha *BigFloat) Mul(beta *BigFloat) *BigFloat {
	a, b := alpha.Real, alpha.Imag
	c, d := beta.Real, beta.Imag

	ac := new(big.Float).Mul(a, c)
	bd := new(big.Float).Mul(b, d)
	ad := new(big.Float).Mul(a, d)
	bc := new(big.Float).Mul(b, c)

	r := new(big.Float).Sub(ac, bd)
	i := new(big.Float).Add(ad, bc)

	return &BigFloat{r, i}
}

func (z *BigFloat) Abs() *big.Float {
	a, b := z.Real, z.Imag

	a2 := new(big.Float).Mul(a, a)
	b2 := new(big.Float).Mul(b, b)

	a2_b2 := new(big.Float).Add(a2, b2)

	return new(big.Float).Sqrt(a2_b2)
}

func (z *BigFloat) ToComplex128() complex128 {
	r, _ := z.Real.Float64()
	i, _ := z.Imag.Float64()
	return complex(r, i)
}
