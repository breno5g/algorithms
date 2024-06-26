package search

type Binary struct{}

// Search for a value in O(log n) time complexity
func (b *Binary) Search(list []int, value int) (int, int) {
	operations := 0
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		pitaco := list[mid]
		operations += 1

		if pitaco > value {
			high = mid - 1
		} else if pitaco < value {
			low = mid + 1
		} else {
			return mid, operations
		}
	}

	return -1, operations
}

// Search for a value in O(n) time complexity
func (b *Binary) IterativeSearch(list []int, value int) (int, int) {
	operations := 0

	for i, v := range list {
		if v == value {
			return i, operations
		}
		operations += 1
	}

	return -1, operations
}

// Search for a value in O(log n) time complexity
func (b *Binary) RecursiveSearch(list []int, value int, low, high, operations int) (int, int) {

	mid := (low + high) / 2
	operations += 1

	if low > high {
		return -1, operations
	}

	if list[mid] == value {
		return mid, operations
	} else if list[mid] > value {
		return b.RecursiveSearch(list, value, low, mid-1, operations)
	} else {
		return b.RecursiveSearch(list, value, mid+1, high, operations)
	}
}
