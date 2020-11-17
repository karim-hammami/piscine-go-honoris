package piscine

func IsPrime(nb int) bool {
	if nb > 0 {
		if nb == 1 {
			return false
		}
		for d := 2; d*d <= nb; d++ {
			if nb%d == 0 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
