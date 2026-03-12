package q85

import (
	"reflect"
	"testing"
)

func TestBytesAnd(t *testing.T) {
	type args struct {
		bytes1 []byte
		bytes2 []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				bytes1: []byte{0b110101},
				bytes2: []byte{0b110100},
			},
			want: []byte{0b110100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesAnd(tt.args.bytes1, tt.args.bytes2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersectionToAreaBytes(t *testing.T) {
	type args struct {
		rows         int
		startRow     []byte
		intersection []byte
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
				startRow:     []byte{'1', '1', '0', '1', '0', '1'},
				intersection: []byte{'1', '1', '0', '1', '0', '0'},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectionToAreaBytes(tt.args.rows, tt.args.startRow, tt.args.intersection); got != tt.want {
				t.Errorf("intersectionToAreaBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximalRectangle(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
