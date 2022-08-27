package fincalc

import "testing"

func TestPresentValue(t *testing.T) {
	c := new(CashFlow)
	c.rate = 0.08
	c.AddCashPayment(200, 1)
	c.AddCashPayment(300, 2)
	c.AddCashPayment(500, 3)
	c.AddCashPayment(-1000, 4)

	pv := c.PresentValue()

	if pv != 104.27309898935164 {
		t.Log("expected 104.27309898935164, got", pv)
		t.Fail()
	}
}
