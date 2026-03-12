package q85

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_rowToInt(t *testing.T) {
	type args struct {
		row []byte
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{row: []byte{'1', '0', '1', '1'}},
			want: 0b1101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rowToInt(tt.args.row); got != tt.want {
				t.Errorf("rowToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersectionToArea(t *testing.T) {
	type args struct {
		rows         int
		startRow     uint64
		intersection uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				rows:         2,
				startRow:     0b110101,
				intersection: 0b110100,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectionToArea(tt.args.rows, tt.args.startRow, tt.args.intersection); got != tt.want {
				t.Errorf("intersectionToArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximalRectangleBigInt(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			}},
			want: 6,
		},
		{
			name: "test2",
			args: args{matrix: [][]byte{
				{'0', '1'},
				{'1', '0'},
			}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rowToBigInt(t *testing.T) {
	type args struct {
		row []byte
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{row: []byte{'1', '0', '1', '1'}},
			want: big.NewInt(0b1101),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rowToBigInt(tt.args.row); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rowToBigInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersectionToAreaBigInt(t *testing.T) {
	type args struct {
		rows         int
		startRow     *big.Int
		intersection *big.Int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				rows:         2,
				startRow:     big.NewInt(0b110101),
				intersection: big.NewInt(0b110100),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectionToAreaBigInt(tt.args.rows, tt.args.startRow, tt.args.intersection); got != tt.want {
				t.Errorf("intersectionToAreaBigInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
