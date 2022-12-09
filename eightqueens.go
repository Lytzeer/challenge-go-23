package piscine

import "fmt"

const n = 8 // nombre de reines à placer

// cols[i] est la colonne où se trouve la reine en ligne i
var cols [n]int

// cherche une solution à partir de la ligne k
func cherche(k int) bool {
	if k == n {
		// une solution a été trouvée
		return true
	}

	// essaie chaque colonne pour la ligne k
	for c := 0; c < n; c++ {
		if estValide(k, c) {
			// place la reine à la ligne k et colonne c
			cols[k] = c

			// cherche une solution pour la ligne suivante
			if cherche(k + 1) {
				return true
			}
		}
	}

	// aucune solution pour cette ligne
	return false
}

// retourne vrai si la reine peut être placée à la ligne k et colonne c
func estValide(k, c int) bool {
	// vérifie si une reine se trouve déjà sur la même colonne
	for i := 0; i < k; i++ {
		if cols[i] == c {
			return false
		}
	}

	// vérifie si une reine se trouve sur la même diagonale
	for i, j := k-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if cols[i] == j {
			return false
		}
	}
	for i, j := k-1, c+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if cols[i] == j {
			return false
		}
	}

	return true
}

func main() {
	if cherche(0) {
		// affiche la première solution trouvée
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if cols[i] == j {
					fmt.Print("Q ")
				} else {
					fmt.Print(". ")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Aucune solution trouvée")
	}
}
