package q8

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "42",
			args: args{str: "42"},
			want: 42,
		},
		{
			name: "-42",
			args: args{str: "-42"},
			want: -42,
		},
		{
			name: "- 42",
			args: args{str: "- 42"},
			want: 0,
		},
		{
			name: "4 2",
			args: args{str: "4 2"},
			want: 4,
		},
		{
			name: " D 4212h",
			args: args{str: " D 4212h"},
			want: 0,
		},
		{
			name: "words and 987",
			args: args{str: "words and 987"},
			want: 0,
		},
		{
			name: "20000000000000000000",
			args: args{str: "20000000000000000000"},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
