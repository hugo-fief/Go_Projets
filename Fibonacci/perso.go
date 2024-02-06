package main

import (
	"fmt"
)

func fibonacciGenerator(numb int) {
	term1, term2 := 0, 1
	total := 0

	fmt.Printf("Affichage de la suite Fibonacci pour les %d premiers éléments : ", numb)

	for index := 0; index < numb; index++ {
		if index > 0 {
			fmt.Print(", ")
		}

		fmt.Print(term1)
		total += term1
		term1, term2 = term2, term2+term1
	}

	fmt.Printf("\nTotal de la suite Fibonacci %d", total)
}

func main() {
	var definedNumber int
	fmt.Print("Veuillez insérer un nombre entier : ")
	fmt.Scan(&definedNumber)

	fibonacciGenerator(definedNumber)

}
