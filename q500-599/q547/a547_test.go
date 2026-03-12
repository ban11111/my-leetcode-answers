package q547

import (
	"math/big"
	"testing"
)

func Test_findCircleNum(t *testing.T) {
	type args struct {
		isConnected [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{isConnected: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}},
			want: 2,
		},
		{
			name: "2",
			args: args{isConnected: [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNum(tt.args.isConnected); got != tt.want {
				t.Errorf("findCircleNum() = %v, want %v", got, tt.want)
			}
			if got := findCircleNum2(tt.args.isConnected); got != tt.want {
				t.Errorf("findCircleNum2() = %v, want %v", got, tt.want)
			}
			if got := findCircleNum3(tt.args.isConnected); got != tt.want {
				t.Errorf("findCircleNum3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_count1s(t *testing.T) {
	type args struct {
		num *big.Int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{num: big.NewInt(1)},
			want: 1,
		},
		{
			name: "1",
			args: args{num: big.NewInt(2)},
			want: 1,
		},
		{
			name: "1",
			args: args{num: big.NewInt(3)},
			want: 2,
		},
		{
			name: "1",
			args: args{num: big.NewInt(4)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := count1s(tt.args.num); got != tt.want {
				t.Errorf("count1s() = %v, want %v", got, tt.want)
			}
		})
	}
}
