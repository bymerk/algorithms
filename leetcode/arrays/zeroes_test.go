package arrays_test

import (
	"reflect"
	"testing"

	"github.com/bymerk/algorithms/leetcode/arrays"
)

func TestMoveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "has zeroes",
			args: args{nums: []int{0, 1, 1, 1, 0}},
			want: []int{1, 1, 1, 0, 0},
		},
		{
			name: "hasn't zeroes",
			args: args{nums: []int{1, 1, 1, 2, 3, 4}},
			want: []int{1, 1, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrays.MoveZeroes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveZeroes() = %v, want %v", got, tt.want)
			}

			if got := arrays.MoveZeroesWithAllocation(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveZeroesWithAllocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMoveZeroes(b *testing.B) {
	nums := []int{0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 0}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = arrays.MoveZeroes(nums)
	}
}

func BenchmarkMoveZeroesWithAllocation(b *testing.B) {
	nums := []int{0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 0}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = arrays.MoveZeroesWithAllocation(nums)
	}
}
