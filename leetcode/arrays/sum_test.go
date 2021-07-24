package arrays_test

import (
	"reflect"
	"testing"

	"github.com/bymerk/algorithms/leetcode/arrays"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "exist",
			args: args{nums: []int{0, 1, 3, 5, 10, 9}, target: 8},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrays.TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}

			if got := arrays.TwoSumWithAllocate(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSumWithAllocate() = %v, want %v", got, tt.want)
			}
		})
	}
}
