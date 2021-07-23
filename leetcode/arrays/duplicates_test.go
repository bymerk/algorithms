package arrays_test

import (
	"testing"

	"github.com/bymerk/algorithms/leetcode/arrays"
)

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name:      "empty slice",
			args:      args{nums: []int{}},
			wantCount: 0,
		},
		{
			name:      "non duplicates",
			args:      args{nums: []int{1}},
			wantCount: 1,
		},
		{
			name:      "has duplicates",
			args:      args{nums: []int{0, 0, 0, 1, 1, 1, 1, 2, 2, 3, 3, 4, 5}},
			wantCount: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := arrays.RemoveDuplicates(tt.args.nums); gotCount != tt.wantCount {
				t.Errorf("RemoveDuplicates() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = arrays.RemoveDuplicates([]int{0, 0, 0, 1, 1, 1, 1, 2, 2, 3, 3, 4, 5})
	}
}
