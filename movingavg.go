package fincalc

type MovingAvg struct {
	numPeriods int
	prices     []float64
}

func (m *MovingAvg) AddPriceQuote(close float64) {
	m.prices = append(m.prices, close)
}

func (m *MovingAvg) CalculateMA() []float64 {
	var ma []float64
	sum := 0.0
	for i := 0; i < len(m.prices); i++ {
		sum += m.prices[i]
		if i >= m.numPeriods {
			ma = append(ma, sum/float64(m.numPeriods))
			sum -= m.prices[i-m.numPeriods]
		}
	}
	return ma
}

func (m *MovingAvg) CalculateEMA2() []float64 {
	var ema []float64
	sum := 0.0
	multiplier := 2 / float64(m.numPeriods+1)
	for i := 0; i < len(m.prices); i++ {
		sum += m.prices[i]
		if i == m.numPeriods {
			ema = append(ema, sum/float64(m.numPeriods))
			sum -= m.prices[i-m.numPeriods]
		} else if i > m.numPeriods {
			val := (1-multiplier)*ema[len(ema)-1] + multiplier*m.prices[i]
			ema = append(ema, val)
		}
	}
	return ema
}

func (m *MovingAvg) CalculateEMA() []float64 {
	var ema []float64
	multiplier := 2 / float64(m.numPeriods+1)

	ma := m.CalculateMA()
	ema = append(ema, ma[0])

	for i := m.numPeriods + 1; i < len(m.prices); i++ {
		val := (1-multiplier)*ema[len(ema)-1] + multiplier*m.prices[i]
		ema = append(ema, val)
	}
	return ema
}
