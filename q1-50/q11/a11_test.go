package q11

import (
	"testing"
)

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{height: []int{1, 99, 9}},
			want: 9,
		},
		{
			name: "2",
			args: args{height: []int{100, 100, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 1, 1, 1, 3, 3, 1, 1}},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}