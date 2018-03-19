package main

import (
	"math/rand"
	"testing"
)

func generateRandomArray(arrayLen int) []int {
	var a []int
	for i := 0; i < arrayLen; i++ {
		a = append(a, rand.Intn(MAX_RAND_LIMIT))
	}
	return a
}

func benchmarkStableSort(i int, b *testing.B) {
	a := generateRandomArray(i)
	stableSort(a)
}

func benchmarkNotInPlaceSort(i int, b *testing.B) {
	a := generateRandomArray(i)
	notInPlaceSort(a)
}

func benchmarkMergeSort(i int, b *testing.B) {
	a := generateRandomArray(i)
	MergeSort(a)
}

func BenchmarkStableSort10(b *testing.B) {
	benchmarkStableSort(10, b)
}

func BenchmarkStableSort100(b *testing.B) {
	benchmarkStableSort(100, b)
}

func BenchmarkStableSort1000(b *testing.B) {
	benchmarkStableSort(1000, b)
}

func BenchmarkStableSort10000(b *testing.B) {
	benchmarkStableSort(10000, b)
}

func BenchmarkStableSort100000(b *testing.B) {
	benchmarkStableSort(100000, b)
}

func BenchmarkStableSort1000000(b *testing.B) {
	benchmarkStableSort(1000000, b)
}

func BenchmarkStableSort10000000(b *testing.B) {
	benchmarkStableSort(10000000, b)
}

func BenchmarkNotInPlaceSort1(b *testing.B) {
	benchmarkNotInPlaceSort(1, b)
}

func BenchmarkNotInPlaceSort10(b *testing.B) {
	benchmarkNotInPlaceSort(10, b)
}

func BenchmarkNotInPlaceSort100(b *testing.B) {
	benchmarkNotInPlaceSort(100, b)
}

func BenchmarkNotInPlaceSort1000(b *testing.B) {
	benchmarkNotInPlaceSort(1000, b)
}

func BenchmarkNotInPlaceSort10000(b *testing.B) {
	benchmarkNotInPlaceSort(10000, b)
}

func BenchmarkNotInPlaceSort100000(b *testing.B) {
	benchmarkNotInPlaceSort(100000, b)
}

func BenchmarkNotInPlaceSort1000000(b *testing.B) {
	benchmarkNotInPlaceSort(1000000, b)
}

func BenchmarkNotInPlaceSort10000000(b *testing.B) {
	benchmarkNotInPlaceSort(10000000, b)
}

func BenchmarkMergeSort1(b *testing.B) {
	benchmarkMergeSort(1, b)
}

func BenchmarkMergeSort10(b *testing.B) {
	benchmarkMergeSort(10, b)
}

func BenchmarkMergeSort100(b *testing.B) {
	benchmarkMergeSort(100, b)
}

func BenchmarkMergeSort1000(b *testing.B) {
	benchmarkMergeSort(1000, b)
}

func BenchmarkMergeSort10000(b *testing.B) {
	benchmarkMergeSort(10000, b)
}

func BenchmarkMergeSort100000(b *testing.B) {
	benchmarkMergeSort(100000, b)
}

func BenchmarkMergeSort1000000(b *testing.B) {
	benchmarkMergeSort(1000000, b)
}

func BenchmarkMergeSort10000000(b *testing.B) {
	benchmarkMergeSort(10000000, b)
}
