package arrays_test

import (
	"reflect"
	"testing"

	"github.com/bymerk/algorithms/leetcode/arrays"
)

func TestIntersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "has intersection",
			args: args{
				nums1: []int{0, 1, 2, 3, 4},
				nums2: []int{5, 6, 7, 3, 5, 4},
			},
			want: []int{3, 4},
		},
		{
			name: "hasn't intersection",
			args: args{
				nums1: []int{0, 1, 2, 3, 4},
				nums2: []int{5, 6, 7, 5},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrays.Intersection(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIntersection(b *testing.B) {
	nums1 := []int{0, 1, 2, 3, 4}
	nums2 := []int{5, 6, 7, 3, 5, 4}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = arrays.Intersection(nums1, nums2)
	}
}
