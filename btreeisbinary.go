package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var arr [50]string
	n := 0
	i := &n
	arr = InOrder(root, arr, i)
	for j, s := range arr {
		if j == 0 {
			continue
		}
		if s == "" {
			break
		}
		if atoi(arr[j]) <= atoi(arr[j-1]) {
			return false
		}
	}
	return true
}

func InOrder(root *TreeNode, arr [50]string, i *int) [50]string {
	if root != nil {
		arr = InOrder(root.Left, arr, i)
		arr[*i] = root.Data
		*i++
		arr = InOrder(root.Right, arr, i)
	}
	return arr
}

func atoi(s string) int {
	runes := []rune(s)
	len := 0
	result := 0
	for i := range runes {
		len = i + 1
	}
	if len == 0 {
		return 0
	}
	tens := 1
	for k := 0; k < len-1; k++ {
		if runes[k] == '+' || runes[k] == '-' {
			continue
		}
		tens *= 10
	}
	for i := range runes {
		if (runes[0] == '-' || runes[0] == '+') && i == 0 {
			continue
		}
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		numb := 0
		for j := '0'; j < runes[i]; j++ {
			numb++
		}
		result += numb * tens
		tens /= 10
	}
	if runes[0] == '-' {
		result *= -1
	}
	return result
}
