package q84

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "one",
			args: args{
				heights: []int{1, 2, 1},
			},
			want: 3,
		},
		{
			name: "two",
			args: args{
				heights: []int{2, 1, 5, 6, 2, 3},
			},
			want: 10,
		},
		{
			name: "three",
			args: args{
				heights: []int{2, 1, 2},
			},
			want: 3,
		},
		{
			name: "four",
			args: args{
				heights: []int{3, 1, 2, 3, 1, 2, 3, 1, 2, 3},
			},
			want: 10,
		},
		{
			name: "five",
			args: args{
				heights: []int{3, 2, 3, 2, 3, 2, 3},
			},
			want: 14,
		},
		{
			name: "six",
			args: args{
				heights: []int{5, 4, 1, 2},
			},
			want: 8,
		},
		{
			name: "seven",
			args: args{
				heights: []int{4, 2},
			},
			want: 4,
		},
		{
			name: "eight",
			args: args{
				heights: []int{4, 2, 0, 3, 1},
			},
			want: 4,
		},
		{
			name: "eight-2",
			args: args{
				heights: []int{4, 2, 0, 3, 2, 5},
			},
			want: 6,
		},
		{
			name: "nine",
			args: args{
				heights: []int{3, 6, 5, 7, 4, 8, 1, 0},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
