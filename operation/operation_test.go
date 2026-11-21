package operation

import (
	"reflect"
	"testing"
)

func TestSa(t *testing.T) {
	st := Stacks{
		A: Stack{Items: []int{2, 1, 3}},
		B: Stack{Items: []int{}},
	}

	Sa(&st)

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(st.A.Items, expected) {
		t.Errorf("Expected %v, got %v", expected, st.A.Items)
	}
}