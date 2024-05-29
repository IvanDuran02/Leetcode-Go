package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   bool
	}{
		{[]int{1, 2, 3, 4, 5}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 6, false},
		{[]int{1, 2, 3, 4, 5}, 1, true},
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{}, 1, false},
		{[]int{1}, 1, true},
		{[]int{1}, 2, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := BinarySearch(tt.arr, tt.target)
			if got != tt.want {
				t.Errorf("BinarySearch(%v, %d) = %v; want %v", tt.arr, tt.target, got, tt.want)
			}
		})
	}
}
