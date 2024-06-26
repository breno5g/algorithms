package main

import (
	"fmt"

	"github.com/breno5g/algorithms/sort"
)

func main() {
	// utils.Init()
	// arr := utils.GetArray()
	// binary := search.Binary{}

	// fmt.Println(binary.Search(arr, 778))
	// fmt.Println(binary.IterativeSearch(arr, 778))
	// fmt.Println(binary.RecursiveSearch(arr, 778, 0, len(arr)-1, 0))

	unordered := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	selection := sort.Selection{}
	fmt.Println(selection.Sort(unordered))

}
