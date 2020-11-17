package piscine

func Max(arr []int) int {
	maxResult := 0
	for _, i := range arr {
		if i > maxResult {
			maxResult = i
		}
	}
	return maxResult
}
