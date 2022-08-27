package fincalc

import "testing"

func TestSinglePeriod(t *testing.T) {
	i := new(IntRate)
	i.rate = 0.08
	j := i.SinglePeriod(10000)
	if j != 10800 {
		t.Log("unexpected return value from SinglePeriod. expected", j)
		t.Fail()
	}
}

func TestMultiplePeriod(t *testing.T) {
	i := new(IntRate)
	i.rate = 0.05
	j := i.MultiplePeriod(1000, 4)
	if j != 1215.50625 {
		t.Log("unexpected return value from MultiplePeriod. expected", j)
		t.Fail()
	}
}

func TestContinouosCompounding(t *testing.T) {
	i := new(IntRate)
	i.rate = 0.05
	j := i.ContinuousCompounding(1000, 4)
	if j != 1221.40275816017 {
		t.Log("unexpected return value from ContinuousCompounding. expected", j)
		t.Fail()
	}
}
