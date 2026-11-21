package main

import (
	"push-swap/operation"
	"testing"
)

func TestSortThree(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{"Cas 2 1 3", []int{2, 1, 3}},
		{"Cas 3 2 1", []int{3, 2, 1}},
		{"Cas 3 1 2", []int{3, 1, 2}},
		{"Cas 1 3 2", []int{1, 3, 2}},
		{"Cas 2 3 1", []int{2, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			stacks := operation.Stacks{
				A: operation.Stack{Items: inputCopy},
				B: operation.Stack{Items: []int{}},
			}

			sortThree(&stacks)

			items := stacks.A.Items
			if len(items) != 3 {
				t.Fatalf("La taille de la pile A a changé : got %d, want 3", len(items))
			}

			if !(items[0] < items[1] && items[1] < items[2]) {
				t.Errorf("Échec du tri pour l'entrée %v. Résultat obtenu : %v", tt.input, items)
			}
		})
	}
}