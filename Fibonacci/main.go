package main

import (
	"fmt"
)

func readIntFromInput() int {
	var num int
	fmt.Print("Veuillez saisir un nombre entier : ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'entrée. Veuillez saisir un nombre entier.")
		return 0
	}
	return num
}

func printFibonacciSeries(n int) {
	term1, term2 := 0, 1

	fmt.Printf("Série Fibonacci pour les %d premiers éléments : ", n)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(term1)
		term1, term2 = term2, term1+term2
	}
	fmt.Println()
}

func main() {
	definedNumber := readIntFromInput()

	if definedNumber <= 0 {
		fmt.Println("Nombre incorrect, veuillez saisir un nombre entier strictement positif.")
		return
	}

	printFibonacciSeries(definedNumber)
}
