package q10

import (
	"testing"
)

func Test_isMatch2(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//{
		//	name: "0",
		//	args: args{s: "a", p: "a"},
		//	want: true,
		//},
		//{
		//	name: "1",
		//	args: args{s: "aa", p: "aa"},
		//	want: true,
		//},
		//{
		//	name: "2",
		//	args: args{s: "a", p: "a*"},
		//	want: true,
		//},
		//{
		//	name: "3",
		//	args: args{s: "a", p: "a*a"},
		//	want: true,
		//},
		//{
		//	name: "4",
		//	args: args{s: "aa", p: "a*a"},
		//	want: true,
		//},
		//{
		//	name: "5",
		//	args: args{s: "aab", p: "a*a"},
		//	want: false,
		//},
		//{
		//	name: "6",
		//	args: args{s: "aab", p: "a*ab*"},
		//	want: true,
		//},
		//{
		//	name: "7",
		//	args: args{s: "ab", p: "a*ab*b"},
		//	want: true,
		//},
		{
			name: "7-1",
			args: args{s: "aaaab", p: "a*ab*b"},
			want: true,
		},
		//{
		//	name: "8",
		//	args: args{s: "aab", p: "a*ab*bx*"},
		//	want: true,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch2(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch2() = %v, want %v", got, tt.want)
			}
		})
	}
}
