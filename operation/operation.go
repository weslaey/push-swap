package operation

import "errors"

// Stack représente une pile (a ou b) avec une slice d'entiers.
type Stack struct {
	Items []int
}

// Stacks représente l'ensemble des deux piles du projet.
type Stacks struct {
	A Stack
	B Stack
}

// --- SWAP (Échange les 2 premiers éléments) ---

func swap(s *Stack) error {
	if len(s.Items) < 2 {
		return errors.New("pas assez d'éléments pour effectuer l'échange")
	}
	s.Items[0], s.Items[1] = s.Items[1], s.Items[0]
	return nil
}

func Sa(st *Stacks) {
	_ = swap(&st.A)
}

func Sb(st *Stacks) {
	_ = swap(&st.B)
}

func Ss(st *Stacks) {
	_ = swap(&st.A)
	_ = swap(&st.B)
}

// --- PUSH (Déplace le premier élément d'une pile vers l'autre) ---

func push(src *Stack, dest *Stack) error {
	if len(src.Items) == 0 {
		return errors.New("la pile est vide")
	}
	// Récupère le premier élément de src
	elem := src.Items[0]
	// Retire le premier élément de src
	src.Items = src.Items[1:]
	// Ajoute l'élément au début de dest
	dest.Items = append([]int{elem}, dest.Items...)
	return nil
}

func Pa(st *Stacks) {
	_ = push(&st.B, &st.A) // Met le premier de B en haut de A
}

func Pb(st *Stacks) {
	_ = push(&st.A, &st.B) // Met le premier de A en haut de B
}

// ROTATE (Décale vers le haut le 1er devient le dernier) 

func rotate(s *Stack) error {
	if len(s.Items) < 2 {
		return nil
	}
	first := s.Items[0]
	s.Items = append(s.Items[1:], first)
	return nil
}

func Ra(st *Stacks) {
	_ = rotate(&st.A)
}

func Rb(st *Stacks) {
	_ = rotate(&st.B)
}

func Rr(st *Stacks) {
	_ = rotate(&st.A)
	_ = rotate(&st.B)
}

// --- REVERSE ROTATE (Décale vers le bas : le dernier devient le 1er) ---

func reverseRotate(s *Stack) error {
	if len(s.Items) < 2 {
		return nil
	}
	lastIdx := len(s.Items) - 1
	last := s.Items[lastIdx]
	s.Items = append([]int{last}, s.Items[:lastIdx]...)
	return nil
}

func Rra(st *Stacks) {
	_ = reverseRotate(&st.A)
}

func Rrb(st *Stacks) {
	_ = reverseRotate(&st.B)
}

func Rrr(st *Stacks) {
	_ = reverseRotate(&st.A)
	_ = reverseRotate(&st.B)
}