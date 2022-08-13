package mergesort

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	first := mergeSort(arr[:len(arr)/2])
	second := mergeSort(arr[len(arr)/2:])
	return merge(first, second)
}

func merge(a, b []int) []int {
	final := []int{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
