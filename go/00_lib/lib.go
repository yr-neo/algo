package _0_lib

func GetMaxMin(a []int) (max int, min int) {
	if len(a) <= 0 {
		return
	}

	max, min = a[0], a[0]
	for i := range a {
		if max < a[i] {
			max = a[i]
		}
		if min > a[i] {
			min = a[i]
		}
	}
	return max, min
}
