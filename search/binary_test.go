package search

import (
	"testing"

	"github.com/breno5g/algorithms/utils"
)

func BenchmarkBinarySearch(b *testing.B) {
	utils.Init()

	for i := 0; i < b.N; i++ {
		arr := utils.GetArray()
		binary := Binary{}
		binary.Search(arr, 5)
	}
}

func BenchmarkIterativeSearch(b *testing.B) {
	utils.Init()

	for i := 0; i < b.N; i++ {
		arr := utils.GetArray()
		binary := Binary{}
		binary.IterativeSearch(arr, 5)
	}
}

func BenchmarkRecursiveSearch(b *testing.B) {
	utils.Init()

	for i := 0; i < b.N; i++ {
		arr := utils.GetArray()
		binary := Binary{}
		binary.RecursiveSearch(arr, 5, 0, len(arr)-1, 0)
	}
}
