package sort_test

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/bymerk/algorithms/sort"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty slice",
			args: args{nums: []int{}},
			want: []int{},
		},
		{
			name: "even - sorted",
			args: args{nums: []int{1, 2, 3, 0, -1, 10, 20, -10}},
			want: []int{-10, -1, 0, 1, 2, 3, 10, 20},
		},
		{
			name: "odd - sorted",
			args: args{nums: []int{1, 2, 3, 0, -1, 20, -10}},
			want: []int{-10, -1, 0, 1, 2, 3, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort.MergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	rand.Seed(time.Now().Unix())
	nums := rand.Perm(100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sort.MergeSort(nums)
	}
}
