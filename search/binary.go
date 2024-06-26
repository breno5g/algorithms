package search

type Binary struct{}

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
