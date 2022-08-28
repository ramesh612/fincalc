package fincalc

import (
	"errors"
	"math"
)

func Mean(values []float64) float64 {
	var sum float64

	for _, value := range values {
		sum += value
	}

	return sum / float64(len(values))
}

func StdDev(values []float64) float64 {
	var sum float64

	mean := Mean(values)
	for _, value := range values {
		diff := value - mean
		sum += (diff * diff)
	}

	return math.Sqrt(sum / float64(len(values)-1))
}

func Correlation(valuesA []float64, valuesB []float64) (float64, error) {
	if len(valuesA) != len(valuesB) {
		return 0, errors.New("unequal number of observations between two sets of values")
	}

	var sum float64

	meanA := Mean(valuesA)
	meanB := Mean(valuesB)

	for i := 0; i < len(valuesA); i++ {
		value := (valuesA[i] - meanA) * (valuesB[i] - meanB)
		sum += value
	}

	stdDevA := StdDev(valuesA)
	stdDevB := StdDev(valuesB)

	sum /= (stdDevA * stdDevB)

	return sum / float64(len(valuesB)-1), nil
}
