package util

func SumInt(values []int) int {
	sum := 0

	for _, val := range values {
		sum += val
	}

	return sum
}
