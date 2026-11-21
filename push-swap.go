package main

import (
	"fmt"
	"os"
	"strconv"

	"push-swap/operation"
)

// Trie spécifiquement 3 éléments dans la pile A
func sortThree(st *operation.Stacks) {
	a := st.A.Items
	if len(a) != 3 {
		return
	}

	n1, n2, n3 := a[0], a[1], a[2]

	if n1 > n2 && n2 < n3 && n1 < n3 {
		operation.Sa(st)
		fmt.Println("sa")
	} else if n1 > n2 && n2 > n3 {
		operation.Sa(st)
		fmt.Println("sa")
		operation.Rra(st)
		fmt.Println("rra")
	} else if n1 > n2 && n2 < n3 && n1 > n3 {
		operation.Ra(st)
		fmt.Println("ra")
	} else if n1 < n2 && n2 > n3 && n1 < n3 {
		operation.Sa(st)
		fmt.Println("sa")
		operation.Ra(st)
		fmt.Println("ra")
	} else if n1 < n2 && n2 > n3 && n1 > n3 {
		operation.Rra(st)
		fmt.Println("rra")
	}
}

// Trouve l'index de la plus petite valeur dans une tranche d'entiers
func findMinIndex(items []int) int {
	minIdx := 0
	for i := 1; i < len(items); i++ {
		if items[i] < items[minIdx] {
			minIdx = i
		}
	}
	return minIdx
}

// Trie 4 ou 5 éléments
func sortFourOrFive(st *operation.Stacks) {
	for len(st.A.Items) > 3 {
		minIdx := findMinIndex(st.A.Items)

		if minIdx <= len(st.A.Items)/2 {
			for i := 0; i < minIdx; i++ {
				operation.Ra(st)
				fmt.Println("ra")
			}
		} else {
			for i := 0; i < len(st.A.Items)-minIdx; i++ {
				operation.Rra(st)
				fmt.Println("rra")
			}
		}

		operation.Pb(st)
		fmt.Println("pb")
	}

	sortThree(st)

	for len(st.B.Items) > 0 {
		operation.Pa(st)
		fmt.Println("pa")
	}
}

// Étape 1 : Indexer les nombres (remplace les valeurs par leur rang de 0 à N-1)
func indexValues(stacks *operation.Stacks) {
	n := len(stacks.A.Items)
	indexed := make([]int, n)

	// Créer une copie triée pour trouver le rang de chaque élément
	sorted := make([]int, n)
	copy(sorted, stacks.A.Items)
	
	// Tri à bulle simple pour le tableau temporaire
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if sorted[i] > sorted[j] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	// Assigner l'index (le rang) à chaque élément d'origine
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if stacks.A.Items[i] == sorted[j] {
				indexed[i] = j
				break
			}
		}
	}
	stacks.A.Items = indexed
}

// Étape 2 : Le tri Radix binaire pour les grands nombres (> 5)
func sortBigNumbers(st *operation.Stacks) {
	indexValues(st)
	n := len(st.A.Items)

	// Trouver le nombre maximum de bits nécessaires
	maxBits := 0
	for (n-1) >> maxBits != 0 {
		maxBits++
	}

	for i := 0; i < maxBits; i++ {
		size := len(st.A.Items)
		for j := 0; j < size; j++ {
			// Vérifie si le i-ème bit est 0
			if ((st.A.Items[0] >> i) & 1) == 0 {
				operation.Pb(st)
				fmt.Println("pb")
			} else {
				operation.Ra(st)
				fmt.Println("ra")
			}
		}
		// Remet tout de B vers A
		for len(st.B.Items) > 0 {
			operation.Pa(st)
			fmt.Println("pa")
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	var numbers []int
	seen := make(map[int]bool) // Pour détecter les doublons facilement

	for _, arg := range os.Args[1:] {
		// Parse en int64 avec une limite de 32 bits (attrape INT_MAX / INT_MIN et les lettres)
		val64, err := strconv.ParseInt(arg, 10, 32)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
		val := int(val64)

		// Vérification des doublons
		if seen[val] {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
		seen[val] = true

		numbers = append(numbers, val)
	}

	stacks := operation.Stacks{
		A: operation.Stack{Items: numbers},
		B: operation.Stack{Items: []int{}},
	}

	switch len(stacks.A.Items) {
	case 2:
		if stacks.A.Items[0] > stacks.A.Items[1] {
			operation.Sa(&stacks)
			fmt.Println("sa")
		}
	case 3:
		sortThree(&stacks)
	case 4, 5:
		sortFourOrFive(&stacks)
	default:
		sortBigNumbers(&stacks)
	}
}