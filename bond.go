package fincalc

import "strconv"

type Bond struct {
	institution string
	principal   float64
	coupon      float64
	numPeriods  int
}

func (b *Bond) InterestRate() float64 {
	return b.coupon / b.principal
}

func (b *Bond) String() string {
	return b.institution + "," +
		strconv.FormatFloat(b.principal, 'f', 2, 64) + "," +
		strconv.FormatFloat(b.coupon, 'f', 2, 64) + "," +
		strconv.Itoa(b.numPeriods)
}
