package fincalc

import (
	"math"
)

type CashFlow struct {
	rate         float64
	timePeriods  []int
	cashPayments []float64
}

func (c *CashFlow) AddCashPayment(value float64, timePeriod int) {
	c.cashPayments = append(c.cashPayments, value)
	c.timePeriods = append(c.timePeriods, timePeriod)
}

func (c *CashFlow) PresentValue() float64 {
	var total float64
	for i := 0; i < len(c.cashPayments); i++ {
		total += c.presentValue(c.cashPayments[i], c.timePeriods[i])
	}
	return total
}

func (c *CashFlow) presentValue(futureValue float64, timePeriod int) float64 {
	return futureValue / math.Pow(1+c.rate, float64(timePeriod))
}
