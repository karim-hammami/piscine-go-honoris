package piscine

func SortWordArr(array []string) {
	strLen := 0
	for range array {
		strLen++
	}

	for i := 0; i < strLen-1; i++ {
		for j := 0; j < strLen-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}
