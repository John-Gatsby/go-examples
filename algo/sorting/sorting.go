package sorting

func BubbleSort(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func ShakerSort(a []int) {
	left := 0
	right := len(a) - 1
	pivot := 0

	for left < right {

		for i := left; i < right; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				pivot = i
			}
		}
		right = pivot

		for i := right; i > left; i-- {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
				pivot = i - 1
			}
		}
		left = pivot
	}
}

func SelectSort(a []int) {
	for i := len(a) - 1; i > 1; i-- {
		max := i
		for j := 0; j < i; j++ {
			if a[j] > a[max] {
				max = j
			}
		}
		a[i], a[max] = a[max], a[i]
	}
}

func InsertionSort(a []int) {

}

func QuickSort(a []int) {
	left := 0
	right := len(a) - 1
	pivot := a[len(a)>>1]

	for {
		for ; a[left] < pivot; left++ {
		}
		for ; a[right] > pivot; right-- {
		}

		if left > right {
			break
		}

		a[left], a[right] = a[right], a[left]
		left++
		right--

		if left > right {
			break
		}
	}

	if right > 0 {
		QuickSort(a[:right+1])
	}
	if left < len(a) {
		QuickSort(a[left:])
	}
}
