package calculations

import (
	"math"
	"sort"
)

func Average(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

func Median(data []float64) float64 {
	sort.Float64s(data)
	midpoint := len(data) / 2
	if len(data)%2 == 0 {
		return (data[midpoint-1] + data[midpoint]) / 2.0
	} else {
		return data[midpoint]
	}
}

func Variance(numbers []float64) float64 {
	var squaredsum float64
	mean := 0.00
	n := float64(len(numbers))
	sum := 0.00
	for _, number := range numbers {
		sum += number
	}
	mean = float64(sum) / n
	for _, number := range numbers {
		deviation := float64(number) - mean
		squaredsum += (math.Pow(deviation, 2))
	}
	return squaredsum / n
}
func StandardDev(numbers []float64) float64 {
	var squaredsum float64
	mean := 0.00
	n := float64(len(numbers))
	sum := 0.00
	for _, number := range numbers {
		sum += number
	}
	mean = float64(sum) / n
	for _, number := range numbers {
		deviation := float64(number) - mean
		squaredsum += (math.Pow(deviation, 2))
	}
	return math.Sqrt(squaredsum / n)
}
