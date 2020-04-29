package q10

import (
	"testing"
)

func Test_isMatch(t *testing.T) {
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
			name: "test-1",
			args: args{s: "aaaaaa", p: "a*a*."},
			want: true,
		},
		{
			name: "test-2",
			args: args{s: "aaabaaa", p: "a*abaa*."},
			want: true,
		},
		{
			name: "test-3",
			args: args{s: "abcd", p: "ab*c*d"},
			want: true,
		},
		{
			name: "test-4",
			args: args{s: "acbd", p: "ab*c*d"},
			want: false,
		},
		{
			name: "test-5",
			args: args{s: "ababcd", p: ".*a*b*cd"},
			want: true,
		},
		{
			name: "test-6",
			args: args{s: "ababcd", p: "a*b*cd"},
			want: false,
		},
		{
			name: "test-7",
			args: args{s: "vcxvaaaccd", p: "a*b*.*..cd"},
			want: true,
		},
		{
			name: "test-8",
			args: args{s: "xxxccc", p: "....cc"},
			want: true,
		},
		{
			name: "test-9",
			args: args{s: "xxxccc", p: "..*..cc"},
			want: true,
		},
		{
			name: "test-10",
			args: args{s: "aaabbaabb", p: ".*bb"},
			want: true,
		},
		{
			name: "test-11",
			args: args{s: "aaabbaabb", p: ".*bb.*bb"},
			want: true,
		},
		{
			name: "test-12",
			args: args{s: "aaabbaabb", p: ".*bb..*.bb"},
			want: true,
		},
		{
			name: "test-13",
			args: args{s: "aaabbaabb", p: ".*bb...*bb"},
			want: true,
		},
		{
			name: "test-14",
			args: args{s: "aaabbaabb", p: ".*bb.*..bb"},
			want: true,
		},
		{
			name: "test-15",
			args: args{s: "aaabbaabb", p: ".*bb.*.bb"},
			want: true,
		},
		{
			name: "test-16",
			args: args{s: "aaabbaabbcc", p: ".*bb.*cc"},
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

func Test_combinePlainTokens(t *testing.T) {
	original := &original{}
	type args struct {
		tokens []token
	}
	tests := []struct {
		name          string
		args          args
		wantTokenSets tokenSets
	}{
		{
			name:          "1",
			args:          args{tokens: []token{original, &star{}, &star{}, &dotStar{}, original}},
			wantTokenSets: tokenSets{original, &dotStarOriginal{}},
		},
		{
			name:          "2",
			args:          args{tokens: []token{&star{}, &star{}, &dotStar{}, original}},
			wantTokenSets: tokenSets{&dotStarOriginal{}},
		},
		{
			name:          "3",
			args:          args{tokens: []token{&star{}, &star{}, original, &dotStar{}, original, &star{}}},
			wantTokenSets: tokenSets{&starOriginal{}, &dotStarOriginal{}, &starOriginal{}},
		},
		{
			name:          "4",
			args:          args{tokens: []token{&star{}, &star{}, original}},
			wantTokenSets: tokenSets{&starOriginal{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTokenSets := combinePlainTokens(tt.args.tokens)
			if len(gotTokenSets) != len(tt.wantTokenSets) {
				t.Errorf("combinePlainTokens() len not equal = %v, want %v", len(gotTokenSets), len(tt.wantTokenSets))
				return
			}
			for i := 0; i < len(gotTokenSets); i++ {
				if gotTokenSets[i].Type() != tt.wantTokenSets[i].Type() {
					t.Errorf("combinePlainTokens().[%d] = %v, want %v", i, gotTokenSets, tt.wantTokenSets)
				}
			}
		})
	}
}

func Test_isMatch1(t *testing.T) {
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
		//	args: args{s:"aaaaaaa", p:"a*b*"},
		//	want: true,
		//},
		//{
		//	name: "1",
		//	args: args{s:"ababcd", p:"a*b*cd"},
		//	want: false,
		//},
		//{
		//	name: "2",
		//	args: args{s:"aaabbbbabcd", p:"a*b*a*b*cd"},
		//	want: true,
		//},
		//{
		//	name: "3",
		//	args: args{s:"abcabcd", p:"a*b*a*b*cd"},
		//	want: false,
		//},
		//{
		//	name: "4",
		//	args: args{s:"abcd", p:"a*b*a*b*cd"},
		//	want: true,
		//},
		//{
		//	name: "5",
		//	args: args{s:"cd", p:"a*b*a*b*cd"},
		//	want: true,
		//},
		//{
		//	name: "6",
		//	args: args{s:"babacd", p:"a*b*a*b*cd"},
		//	want: false,
		//},
		{
			name: "7",
			args: args{s: "acd", p: "a*acd"},
			want: true,
		},
		//{
		//	name: "8",
		//	args: args{s:"aacdjkl", p:"a*acdj*kl"},
		//	want: true,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseToken(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name          string
		args          args
		wantTokenSets tokenSets
	}{
		{
			name:          "1",
			args:          args{p: "a*b*cd"},
			wantTokenSets: tokenSets{&starOriginal{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTokenSets := parseToken(tt.args.p)
			if len(gotTokenSets) != len(tt.wantTokenSets) {
				t.Errorf("combinePlainTokens() len not equal = %v, want %v", len(gotTokenSets), len(tt.wantTokenSets))
				return
			}
			for i := 0; i < len(gotTokenSets); i++ {
				if gotTokenSets[i].Type() != tt.wantTokenSets[i].Type() {
					t.Errorf("combinePlainTokens().[%d] = %v, want %v", i, gotTokenSets, tt.wantTokenSets)
				}
			}
		})
	}
}
