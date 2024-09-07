package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func affichage() {
	fmt.Println("#################################")
	fmt.Println("\tBonjour, bienvenue au jeu du morpion !")
	fmt.Println("#################################")
}

func affichagemorpion(morpionparam [3][3]string) {
	fmt.Println("-------------")
	fmt.Println("|", morpionparam[0][0], "|", morpionparam[0][1], "|", morpionparam[0][2], "|")
	fmt.Println("-------------")
	fmt.Println("|", morpionparam[1][0], "|", morpionparam[1][1], "|", morpionparam[1][2], "|")
	fmt.Println("-------------")
	fmt.Println("|", morpionparam[2][0], "|", morpionparam[2][1], "|", morpionparam[2][2], "|")
	fmt.Println("-------------")
}

func main() {
	affichage()

	var morpion [3][3]string
	morpion = initmorpion(morpion)

	affichagemorpion(morpion)
	var joueur int = 1
	var caractere string = "X"
	var passage int = 1
	var pasnul int
	var varpartiegagne int

	for true {

		scanner3 := bufio.NewScanner(os.Stdin)
		fmt.Println("Joueur ", joueur, " entrez un nombre compris entre 1 à 9 : ")
		scanner3.Scan()
		choix, err := strconv.Atoi(scanner3.Text()) // Récupération du choix user

		if err != nil {
			// Si la conversion n'a pas fonctionné alors on affiche un message d'erreur
			fmt.Println("Entrez un nombre et non autre chose !")
		}

		if choix > 10 { // choix (case choisi)
			fmt.Println("Votre nombre doit être compris entre 0 à 9 !")
		} else { // sinon on l'accepte
			//modifier la case du morpion
			switch choix {
			case 1: // [0,0]
				if morpion[0][0] != "X" && morpion[0][0] != "O" {
					morpion[0][0] = caractere
					passage = 1
				} else {
					fmt.Println("Cette case est déjà prise !")
					passage = 0
				}
			case 2:
				if morpion[0][1] != "X" && morpion[0][1] != "O" {
					morpion[0][1] = caractere
					passage = 1
				} else {
					fmt.Println("Cette case est déjà prise !")
					passage = 0

				}
			case 3:
				if morpion[0][2] != "X" && morpion[0][2] != "O" {
					morpion[0][2] = caractere
					passage = 1
				} else {
					fmt.Println("Cette case est déjà prise !")
					passage = 0

				}
			case 4:
				if morpion[1][0] != "X" && morpion[1][2] != "O" {
					morpion[1][0] = caractere
					passage = 1
				} else {
					fmt.Println("Cette case est déjà prise !")
					passage = 0

				}
			case 5:
				if morpion[1][1] != "X" && morpion[1][1] != "O" {
					morpion[1][1] = caractere
					passage = 1
				} else {
					fmt.Println("Cette case est déjà prise !")
					passage = 0

				}
			case 6:
				if morpion[1][2] != "X" && morpion[1][2] != "O" {
					morpion[1][2] = caractere
					passage = 1
				} else {
					fmt.Println("Cette case est déjà prise !")
					passage = 0

				}
			case 7:
				if morpion[2][0] != "X" && morpion[2][0] != "O" {
					morpion[2][0] = caractere
					passage = 1
				} else {
					fmt.Println("Cette case est déjà prise !")
					passage = 0

				}
			case 8:
				if morpion[2][1] != "X" && morpion[2][1] != "O" {
					morpion[2][1] = caractere
					passage = 1
				} else {
					fmt.Println("Cette case est déjà prise !")
					passage = 0

				}
			case 9:
				if morpion[2][2] != "X" && morpion[2][2] != "O" {
					morpion[2][2] = caractere
					passage = 1
				} else {
					fmt.Println("Cette case est déjà prise !")
					passage = 0

				}
			default:
				println("Mauvais choix !") // Valeur par défaut
			}
		}

		if passage == 1 {

			//Partie nulle !
			pasnul = partienulle(morpion)
			if pasnul == 0 {
				println("Partie nulle!")
				break
			}

			//Partie gagné !
			varpartiegagne = partiegagne(morpion)
			if varpartiegagne == 1 {
				println("Joueur ", joueur, " vous avez gagné !")
				break
			}

			if joueur == 1 {
				joueur = 2
				caractere = "O"
			} else {
				joueur = 1
				caractere = "X"
			}

			affichagemorpion(morpion)
		}

	}

}

func initmorpion(morpioninit [3][3]string) [3][3]string {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			morpioninit[i][j] = strconv.Itoa((3 * i) + 1 + j)
		}
	}
	return morpioninit
}

func partienulle(morpioninit [3][3]string) int {
	var pasnul int = 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if morpioninit[i][j] != "X" && morpioninit[i][j] != "O" {
				pasnul = 1
			}
		}
	}
	return pasnul
}

func partiegagne(morpioninit [3][3]string) int {
	var partiegagne int = 0
	//Lignes
	if morpioninit[0][0] == morpioninit[0][1] && morpioninit[0][1] == morpioninit[0][2] { // Première ligne même caractères
		partiegagne = 1
	}
	if morpioninit[1][0] == morpioninit[1][1] && morpioninit[1][1] == morpioninit[1][2] { // Deuxième ligne même caractères
		partiegagne = 1
	}
	if morpioninit[2][0] == morpioninit[2][1] && morpioninit[2][1] == morpioninit[2][2] { // Troisième ligne même caractères
		partiegagne = 1
	}
	//Colonnes
	if morpioninit[0][0] == morpioninit[1][0] && morpioninit[1][0] == morpioninit[2][0] { // Première colonne même caractères
		partiegagne = 1
	}
	if morpioninit[0][1] == morpioninit[1][1] && morpioninit[1][1] == morpioninit[2][1] { // Deuxième colonne même caractères
		partiegagne = 1
	}
	if morpioninit[0][2] == morpioninit[1][2] && morpioninit[1][2] == morpioninit[2][2] { // Troisième colonne même caractères
		partiegagne = 1
	}
	//Diagonales
	if morpioninit[0][0] == morpioninit[1][1] && morpioninit[1][1] == morpioninit[2][2] { // Diago1
		partiegagne = 1
	}
	if morpioninit[2][0] == morpioninit[1][1] && morpioninit[1][1] == morpioninit[0][2] { // Diago2
		partiegagne = 1
	}
	return partiegagne
}
