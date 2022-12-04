package one

import (
	"reflect"
	"testing"
)

func TestPick(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{[]string{"a", "b", "c"}, "a"},
		{[]string{"", "b", "c"}, "b"},
		{[]string{"", "", "c"}, "c"},
		{[]string{"", "", ""}, ""},
		{[]string{}, ""},
		{[]string{""}, ""},
		{[]string{"a"}, "a"},
	}
	for _, tt := range tests {
		if got := Pick(tt.args...); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Pick() = %v, want %v", got, tt.want)
		}
	}

	numTests := []struct {
		args []int
		want int
	}{
		{[]int{1, 2, 3}, 1},
		{[]int{0, 2, 3}, 2},
		{[]int{0, 0, 3}, 3},
		{[]int{0, 0, 0}, 0},
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{0, 1}, 1},
	}
	for _, tt := range numTests {
		if got := Pick(tt.args...); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Pick() = %v, want %v", got, tt.want)
		}
	}
}
