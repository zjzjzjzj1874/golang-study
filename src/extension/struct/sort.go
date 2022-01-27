package _struct

import "sort"

// People example with people
type People struct {
	Name   string
	Age    int
	Height int
}

// By is the type of "less" function that defines the ordering of its People arguments.
type By func(p1, p2 *People) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(p []People) {
	ps := &peopleSorter{
		people: p,
		by:     by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// peopleSorter joins a By function and a slice of People to be sorted.
type peopleSorter struct {
	people []People
	by     func(p1, p2 *People) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *peopleSorter) Len() int {
	return len(s.people)
}

// Swap is part of sort.Interface.
func (s *peopleSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *peopleSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}