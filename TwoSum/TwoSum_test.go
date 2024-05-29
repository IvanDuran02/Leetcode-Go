package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
		{nums: []int{1, 2, 3, 4, 5}, target: 8, want: []int{2, 4}},
		{nums: []int{0, 4, 3, 0}, target: 0, want: []int{0, 3}},
	}

	for _, tt := range tests {
		got := TwoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("TwoSum(%v, %v) = %v; want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
