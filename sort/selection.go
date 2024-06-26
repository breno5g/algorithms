package sort

type Selection struct{}

func (s Selection) Sort(arr []int) []int {
	var min_index int

	for i := 0; i < len(arr); i++ {
		min_index = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}

		temp_arr := arr[i]
		arr[i] = arr[min_index]
		arr[min_index] = temp_arr
	}

	return arr
}
