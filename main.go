package main

import (
	"sort"
)

const (
	MAX_RAND_LIMIT = 1000
)

type arr []int

func (a arr) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a arr) Len() int {
	return len(a)
}

func notInPlaceSort(array arr) []int {
	sort.Sort(array)
	return array
}

func stableSort(array arr) []int {
	sort.Stable(array)
	return array
}

func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func Merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
