package checker

import (
	"strconv"
)

// Fonction pour parser et valider les arguments reçus
func ParseArgs(args []string) ([]int, error) {
	var numbers []int
	for _, arg := range args[1:] {
		val, err := strconv.Atoi(arg)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, val)
	}
	return numbers, nil
}

// Fonction utilitaire pour vérifier si la slice est triée
func IsSorted(stack []int) bool {
	for i := 0; i < len(stack)-1; i++ {
		if stack[i] > stack[i+1] {
			return false
		}
	}
	return true
}