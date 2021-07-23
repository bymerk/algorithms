package arrays_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/bymerk/algorithms/leetcode/arrays"
)

func TestRotateWithAllocate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "rotate",
			args: args{nums: []int{0, 1, 2, 3}, k: 1},
			want: []int{3, 0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrays.RotateWithAllocate(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateWithAllocate() = %v, want %v", got, tt.want)
			}

			if got := arrays.Rotate(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateWithAllocate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRotateWithAllocate(b *testing.B) {
	nums := rand.Perm(1000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = arrays.RotateWithAllocate(nums, 100)
	}
}

func BenchmarkRotate(b *testing.B) {
	nums := rand.Perm(1000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = arrays.Rotate(nums, 100)
	}
}
