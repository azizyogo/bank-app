package float

import "math"

func RoundToDecimalPlaces(value float64, decimalPlaces int) float64 {
	shift := math.Pow(10, float64(decimalPlaces))
	return math.Ceil(value*shift) / shift
}
