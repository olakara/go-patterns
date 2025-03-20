package main

func main() {
	sorter := Sorter{}
	sorter.SetSortStrategy(MergeSort{})
	sorter.Sort([]int{10, 2, 7, 5, 6})

	sorter.SetSortStrategy(QuickSort{})
	sorter.Sort([]int{8, 3, 33, 1, 66, 5})
}
