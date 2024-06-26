package main

import (
	"fmt"

	"github.com/breno5g/algorithms/search"
	"github.com/breno5g/algorithms/utils"
)

func main() {
	utils.Init()
	arr := utils.GetArray()
	binary := search.Binary{}

	fmt.Println(binary.Search(arr, 778))
	fmt.Println(binary.IterativeSearch(arr, 778))
	fmt.Println(binary.RecursiveSearch(arr, 778, 0, len(arr)-1, 0))
}
