package sort

import "testing"

func BenchmarkSelectionSort(b *testing.B) {
	unordered := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	selection := Selection{}

	for i := 0; i < b.N; i++ {
		selection.Sort(unordered)
	}
}
