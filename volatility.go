package fincalc

import "math"

type Volatility struct {
	prices []float64
}

func (v *Volatility) AddPrice(price float64) {
	v.prices = append(v.prices, price)
}

func (v *Volatility) RangeVolatility() float64 {
	if len(v.prices) < 1 {
		return 0
	}

	min := v.prices[0]
	max := min

	for _, p := range v.prices {
		if p < min {
			min = p
		}

		if p > max {
			max = p
		}
	}
	return max - min
}

func (v *Volatility) StdDev() float64 {
	m := v.Mean()
	sum := 0.0
	for _, v := range v.prices {
		val := v - m
		sum += (val * val)
	}
	return math.Sqrt(sum / float64(len(v.prices)-1))
}

func (v *Volatility) Mean() float64 {
	sum := 0.0
	for _, v := range v.prices {
		sum += v
	}
	return sum / float64(len(v.prices))
}

func (v *Volatility) AvgDailyRange() float64 {
	n := len(v.prices)
	if n < 2 {
		return 0
	}

	previous := v.prices[0]
	sum := 0.0

	for i := 1; i < n; i++ {
		r := math.Abs(v.prices[i] - previous)
		sum += r
		previous = v.prices[i]
	}
	return sum / float64(n-1)
}
