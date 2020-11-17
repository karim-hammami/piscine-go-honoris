package piscine

func Map(f func(int) bool, arr []int) []bool {
	arrLength := 0
	for range arr {
		arrLength++
	}
	result := make([]bool, arrLength)
	for i, num := range arr {
		result[i] = f(num)
	}
	return result
}
