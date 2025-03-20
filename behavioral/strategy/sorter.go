package main

type SortStrategy interface {
	Sort([]int)
}

type Sorter struct {
	strategy SortStrategy
}

func (s *Sorter) Sort(data []int) {
	s.strategy.Sort(data)
}

func (s *Sorter) SetSortStrategy(strategy SortStrategy) {
	s.strategy = strategy
}
