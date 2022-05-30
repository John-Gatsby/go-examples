package sorting

import (
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func getUnsorted() []int {
	unsorted := make([]int, 1<<10)
	for i := range unsorted {
		unsorted[i] = i ^ 0x2cc
	}
	return unsorted
}

func isSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	data := ints
	BubbleSort(data[:])
	if !isSorted(data[:]) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	b.StopTimer()
	unsorted := getUnsorted()
	data := make([]int, len(unsorted))
	for i := 0; i < b.N; i++ {
		copy(data, unsorted)
		b.StartTimer()
		BubbleSort(data)
		b.StopTimer()
	}
}
