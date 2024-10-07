package guessrange

import (
	"guess/calculations"
)

func Guessrange(numbers []float64) (float64, float64) {
	mean := calculations.Average(numbers)
	stdDev := calculations.StandardDev(numbers)
	// Use a prediction range of mean ± 2 * standard deviation
	return mean - 2*stdDev, mean + 2*stdDev
}
