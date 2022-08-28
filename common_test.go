package fincalc

import "testing"
import "math"

func TestCorrelation1(t *testing.T) {
	valuesA := []float64{1.2, 2, 2.5, 4, 3, 6, 5.5, 6.3, 7.1, 5.4}
	valuesB := []float64{3.4, 3.3, 3, 5.5, 1.2, 2.4, 3.2, 3.1, 2.9, 3.2}

	corr, err := Correlation(valuesA, valuesB)

	if err == nil {
		if corr != -0.05060099158767409 {
			t.Log("unexpected correlation value", corr)
			t.Fail()
		}
	}
}

func TestCorrelation2(t *testing.T) {
	valuesA := []float64{1, 2, 3, 4, 5, 6, 7}
	valuesB := []float64{10, 9, 8, 7, 6, 5, 4}

	corr, err := Correlation(valuesA, valuesB)

	if err == nil {
		if math.Round(corr) != -1 {
			t.Log("unexpected correlation value", math.Round(corr))
			t.Fail()
		}
	}
}
