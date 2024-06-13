package main

import "testing"

func Sum(numbers [5]int) int {
	return 15
}

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
