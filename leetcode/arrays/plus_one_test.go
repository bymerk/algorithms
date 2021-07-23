package arrays_test

import (
	"reflect"
	"testing"

	"github.com/bymerk/algorithms/leetcode/arrays"
)

func TestPlusOne(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "without allocate",
			args: args{nums: []int{1, 0, 9}},
			want: []int{1, 1, 0},
		},
		{
			name: "with allocate",
			args: args{nums: []int{9, 9, 9}},
			want: []int{1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrays.PlusOne(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
