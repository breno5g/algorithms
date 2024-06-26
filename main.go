package main

import (
	"fmt"

	"github.com/breno5g/algorithms/search"
)

func main() {
	binary := search.Binary{}
	arr := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		arr[i] = i
	}
	// fmt.Println(binary.Search(arr, 7435))
	fmt.Println(binary.IterativeSearch(arr, 7435))
}
