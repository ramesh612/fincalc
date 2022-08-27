package fincalc

import "testing"

func TestBondIRR(t *testing.T) {
	b := new(Bond)
	b.institution = "XYZ"
	b.principal = 10000
	b.coupon = 5000
	b.numPeriods = 20
	irr := b.InterestRate()

	if irr != 0.5 {
		t.Log("expected 0.5, got", irr)
		t.Fail()
	}
}
