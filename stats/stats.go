package stats

import "errors"

const (
	Minimum = "minimum"
	Average = "average"
	Maximum = "maximum"
)

func min(values ...int) int {
	var minValue int = values[0]
	for _, value := range values {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}

func max(values ...int) int {
	var maxValue int = values[0]
	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func avg(values ...int) int {
	var sum int = 0
	for _, value := range values {
		sum += value
	}
	return sum / len(values)
}

func Operation(op string)  (func(values ...int) int, error) {
	switch op {
	case Minimum:
		return min, nil
	case Average:
		return avg, nil
	case Maximum:
		return max, nil
	default:
		return nil, errors.New("Invalid Operation")
	}
}