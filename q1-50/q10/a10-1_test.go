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
		{
			name: "0",
			args: args{s: "a", p: "a"},
			want: true,
		},
		{
			name: "1",
			args: args{s: "aa", p: "aa"},
			want: true,
		},
		{
			name: "2",
			args: args{s: "a", p: "a*"},
			want: true,
		},
		{
			name: "3",
			args: args{s: "a", p: "a*a"},
			want: true,
		},
		{
			name: "4",
			args: args{s: "aa", p: "a*a"},
			want: true,
		},
		{
			name: "5",
			args: args{s: "aab", p: "a*a"},
			want: false,
		},
		{
			name: "6",
			args: args{s: "aab", p: "a*ab*"},
			want: true,
		},
		{
			name: "7",
			args: args{s: "ab", p: "a*ab*b"},
			want: true,
		},
		{
			name: "7-1",
			args: args{s: "aaaaaaaaabb", p: "a*ab*b"},
			want: true,
		},
		{
			name: "7-2",
			args: args{s: "abbbbbbbbbbbbbb", p: "a*ab*b"},
			want: true,
		},
		{
			name: "7-3",
			args: args{s: "aaaaabbbbb", p: "a*b*c*d*"},
			want: true,
		},
		{
			name: "7-4",
			args: args{s: "aaaaabbbbbddddddddddddd", p: "a*b*c*d*"},
			want: true,
		},
		{
			name: "8",
			args: args{s: "aa", p: "a*ab*x*"},
			want: true,
		},
		{
			name: "8-1",
			args: args{s: "aab", p: "a*ab*x*"},
			want: true,
		},
		{
			name: "8-1",
			args: args{s: "aaxx", p: "a*ab*x*"},
			want: true,
		},
		{
			name: "9",
			args: args{s: "aaxx", p: "aaxx"},
			want: true,
		},
		{
			name: "9-1",
			args: args{s: "aax", p: "aaxx"},
			want: false,
		},
		{
			name: "9-2",
			args: args{s: "aaxx", p: "aaxxh"},
			want: false,
		},
		{
			name: "9-3",
			args: args{s: "aaxxhxx", p: "aaxxh"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMatchWithDot(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "0",
			args: args{s: "aaxx", p: "...."},
			want: true,
		},
		{
			name: "0-0",
			args: args{s: "a", p: "."},
			want: true,
		},
		{
			name: "0-1",
			args: args{s: "aaxx", p: "..."},
			want: false,
		},
		{
			name: "0-2",
			args: args{s: "aaxx", p: "....."},
			want: false,
		},
		{
			name: "1",
			args: args{s: "aa", p: "..a*"},
			want: true,
		},
		{
			name: "1-1",
			args: args{s: "aa", p: "..a*a"},
			want: false,
		},
		{
			name: "2",
			args: args{s: "aabc", p: ".*a*a"},
			want: false,
		},
		{
			name: "2-1",
			args: args{s: "aabc", p: ".*a*c"},
			want: true,
		},
		{
			name: "2-2",
			args: args{s: "aabc", p: ".*a*bc"},
			want: true,
		},
		{
			name: "2-3",
			args: args{s: "aabc", p: ".*a*.c"},
			want: true,
		},
		{
			name: "2-4",
			args: args{s: "aabc", p: ".*a*.j"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFailed(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{s: "aa", p: "..a*a"},
			want: false,
		},
		{
			name: "2",
			args: args{s: "a", p: ""},
			want: false,
		},
		{
			name: "3",
			args: args{s: "aasdfasdfasdfasdfas", p: "aasdf.*asdf.*asdf.*asdf.*s"},
			want: true,
		},
		{
			name: "3-1",
			args: args{s: "aasdfasdfasdfasdfas", p: "aasdf.*asdf.*asdf.*s"},
			want: true,
		},
		{
			name: "3-2",
			args: args{s: "aasdfasdfasdfasdfas", p: "aasdf.*asdf.*asdf.*asdf.*asdf.*asdf.*s"},
			want: false,
		},
		{
			name: "3-3",
			args: args{s: "aasdfasdfasdfasdfasdfasdfas", p: "aasdf.*asdf.*asdf.*asdf.*asdf.*asdf.*s"},
			want: true,
		},
		{
			name: "4",
			args: args{s: "mississippi", p: "mis*is*ip*."},
			want: true,
		},
		{
			name: "4-1",
			args: args{s: "ssissipp", p: "s*is*ip*"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFailed2(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{s: "mississippi", p: "mis*is*p*."},
			want: false,
		},
		{
			name: "1-0",
			args: args{s: "ssissipp", p: "s*is*p*"},
			want: false,
		},
		{
			name: "2",
			args: args{s: "cbaacacaaccbaabcb", p: "c*b*b*.*ac*.*bc*a*"},
			want: true,
		},
		{
			name: "2-0",
			args: args{s: "XXxxXXaXXxaaxbXaababbbbXbxxXXb", p: "c*b*b*.*ac*.*bc*a*"},
			want: true,
		},
		{
			name: "3",                    // dotStar plain dotStar plain star plain star plain
			args: args{s: "abbaaaabaabbcba", p:"a*.*ba.*c*..a*.a*."},
			want: true,
		},
		{
			name: "4",
			args: args{s: "aaa", p:"ab*a"},
			want: false,
		},
		{
			name: "4-1",
			args: args{s: "aaa", p:"ab*a*c*a"},
			want: true,
		},
		{
			name: "5",
			args: args{s: "a", p:"aa*"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_plainToken_findFirst(t1 *testing.T) {
	newS := func(s string) *string {
		return &s
	}
	type fields struct {
		pattern string
	}
	type args struct {
		s    *string
		left int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  int
	}{
		{
			name:   "1",
			fields: fields{"aba"},
			args:   args{s: newS("aaabaa"), left: 0},
			want:   true,
			want1:  4,
		},
		{
			name:   "2",
			fields: fields{"acd"},
			args:   args{s: newS("aacdjkl"), left: 0},
			want:   true,
			want1:  3,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &plainToken{
				pattern: tt.fields.pattern,
			}
			got, got1 := t.findFirst(tt.args.s, tt.args.left)
			if got != tt.want {
				t1.Errorf("findFirst() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t1.Errorf("findFirst() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
