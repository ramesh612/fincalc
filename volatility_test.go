package fincalc

import "testing"

func TestVolatility(t *testing.T) {
	v := new(Volatility)
	v.AddPrice(3)
	v.AddPrice(3.5)
	v.AddPrice(5)
	v.AddPrice(4.48)
	v.AddPrice(5.2)
	v.AddPrice(6)
	v.AddPrice(6.1)
	v.AddPrice(5.5)
	v.AddPrice(5.2)
	v.AddPrice(5.7)

	rv := v.RangeVolatility()
	adr := v.AvgDailyRange()
	sd := v.StdDev()
	m := v.Mean()

	if rv > 3.1 || adr > 0.62 || sd > 1.03 || m > 4.97 {
		t.Log("unexpected output from Volatility funcs")
		t.Fail()
	}
}
