package find_replications

import (
	"reflect"
	"testing"
)

func TestFindReplications(t *testing.T) {
	type args struct {
		items []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{items: []string{"a", "b", "c", "b", "d", "a"}}, want: []int{0, 1, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindReplications(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindReplications() = %v, want %v", got, tt.want)
			}
		})
	}
}
