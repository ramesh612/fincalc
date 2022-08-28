package fincalc

import "reflect"
import "testing"

func TestMACalculator(t *testing.T) {
	macalc := new(MovingAvg)
	macalc.numPeriods = 5
	macalc.AddPriceQuote(10)
	macalc.AddPriceQuote(11)
	macalc.AddPriceQuote(22)
	macalc.AddPriceQuote(12)
	macalc.AddPriceQuote(13)
	macalc.AddPriceQuote(23)
	macalc.AddPriceQuote(12)
	macalc.AddPriceQuote(32)
	macalc.AddPriceQuote(12)
	macalc.AddPriceQuote(3)
	macalc.AddPriceQuote(2)
	macalc.AddPriceQuote(22)
	macalc.AddPriceQuote(32)

	expected_ma := []float64{18.2, 18.6, 22.8, 20.8, 19.0, 16.8, 16.6, 20.6}
	expected_ema := []float64{18.2, 16.133333333333333, 21.422222222222224, 18.281481481481485, 13.187654320987658, 9.458436213991773, 13.638957475994516, 19.759304983996344}

	if !reflect.DeepEqual(macalc.CalculateMA(), expected_ma) {
		t.Log("CalculateMA failed")
		t.Fail()
	}

	if !reflect.DeepEqual(macalc.CalculateEMA(), expected_ema) {
		t.Log("CalculateEMA failed")
		t.Fail()
	}

	if !reflect.DeepEqual(macalc.CalculateEMA2(), expected_ema) {
		t.Log("CalculateEMA2 failed")
		t.Fail()
	}
}
