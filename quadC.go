package main

import "github.com/01-edu/z01"

func main() {
	QuadA(1, 1)
}

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		z01.PrintRune('P')
		z01.PrintRune('A')
		z01.PrintRune('S')
		z01.PrintRune('O')
		z01.PrintRune('U')
		z01.PrintRune('F')
		z01.PrintRune('\n')
		return
	}

	if x == 1 {
		for i := 0; i < y; i++ { // J'ai mis un -1 après le y, car problème d'index, autres solutions faire commencer i à 1 et non 0, pu faire x-- y-- décale l'index
			if i == 0 {
				z01.PrintRune('A')
			} else if i == y-1 {
				z01.PrintRune('C') 
			} else {
				z01.PrintRune('B')
			}
			z01.PrintRune('\n')
		}
		return
	}

	// Ligne du haut
	z01.PrintRune('A')
	for i := 0; i < x-2; i++ {
		z01.PrintRune('B')
	}
	z01.PrintRune('A')
	z01.PrintRune('\n')

	if y <= 1 {
		return
	}

	// Lignes du milieu
	for i := 0; i < y-2; i++ {
		z01.PrintRune('B')
		for j := 0; j < x-2; j++ {
			z01.PrintRune(' ')
		}
		z01.PrintRune('B')
		z01.PrintRune('\n')
	}

	// Ligne du bas
	if y > 1 {
		z01.PrintRune('C')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
		z01.PrintRune('\n')
	}
}
