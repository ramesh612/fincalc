package fincalc

import "math"

type IntRate struct {
	rate float64
}

func (i *IntRate) SinglePeriod(value float64) float64 {
	return value * (1 + i.rate)
}

func (i *IntRate) MultiplePeriod(value float64, numPeriods int) float64 {
	return value * math.Pow(1+i.rate, float64(numPeriods))
}

func (i *IntRate) ContinuousCompounding(value float64, numPeriods int) float64 {
	return value * math.Exp(i.rate*float64(numPeriods))
}
