package q8

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	"testing"
)

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

func TestName(t *testing.T) {
	const (
		MarProfileSwitchMandatory = "Mandatory" // 默认?
		MqrProfileSwitchOptional  = "Optional"
		MqrProfileSwitchHidden    = "Hidden"
	)

	// 各项信息(展示/必填)开关
	type MqrProfileSwitch struct {
		IdInfoSwitch           string `json:"id_info_switch"`
		WorkInfoSwitch         string `json:"work_info_switch"`
		ContactInfoSwitch      string `json:"contact_info_switch"`
		AccountInfoSwitch      string `json:"account_info_switch"`
		PersonalInfoSwitch     string `json:"personal_info_switch"`
		SupplementaryDocSwitch string `json:"supplementary_doc_switch"`
	}
	fmt.Println(json.StringifyJson(MqrProfileSwitch{
		IdInfoSwitch:           MarProfileSwitchMandatory,
		WorkInfoSwitch:         MarProfileSwitchMandatory,
		ContactInfoSwitch:      MarProfileSwitchMandatory,
		AccountInfoSwitch:      MarProfileSwitchMandatory,
		PersonalInfoSwitch:     MarProfileSwitchMandatory,
		SupplementaryDocSwitch: MarProfileSwitchMandatory,
	}))
}

func TestSomething(t *testing.T) {
	type test struct {
		s string
	}
	var slice []interface{}
	slice = append(slice, &test{"fff"})
	fmt.Println(slice)
	for _, x := range slice {
		x = &test{"ccc"}
		fmt.Println(x)
	}
	x := slice[0]
	x = &test{"ccc"}
	fmt.Println(x)
	fmt.Println(slice, slice[0].(*test).s)
}
