package main

import (
	"fmt"
	"testing"
)

type test struct {
	data   []int
	answer int
}

// sum(i ...int) int

func TestSum(t *testing.T) {
	test := sum(3, 2, 1)
	result := 6
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

// failed version
/* func TestSum2(t *testing.T) {
	test := sum(3, 2, 1)
	result := 7
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
} */

// mutiple(i ...int) int

func TestMutiple(t *testing.T) {
	test := multiple(10, 10)
	result := 100
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

func TestSumInTable(t *testing.T) {
	tests := []test{
		{
			data:   []int{1, 2, 3},
			answer: 6,
		},
		{
			data:   []int{10, 11, 12},
			answer: 33,
		},
		{
			data:   []int{-5, 0, 5, 10},
			answer: 10,
		},
	}

	for _, v := range tests {
		z := sum(v.data...)
		if z != v.answer {
			t.Error("Expected:", v.answer, "Got:", z)
		}
	}
}

func Examplesum() {
	fmt.Println(sum(3, 2, 1))
	// Output: 7
}

func BenchmarkSum(b *testing.B) {
	tests := []test{
		{
			data:   []int{1, 2, 3},
			answer: 6,
		},
		{
			data:   []int{10, 11, 12},
			answer: 33,
		},
		{
			data:   []int{-5, 0, 5, 10},
			answer: 10,
		},
	}

	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			sum(v.data...)
		}
	}
}

func BenchmarkMutiple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiple(2, 1)
	}
}
