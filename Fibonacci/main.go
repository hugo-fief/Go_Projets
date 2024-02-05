package main

import (
	"fmt"
	"math"
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

func isInteger(num float64) bool {
	return num == math.Floor(num)
}

func printFibonacciSeries(n int) {
	term1 := 0
	term2 := 1
	nextTerm := 0

	fmt.Printf("Série Fibonacci pour les %d premiers éléments :", n)
	fmt.Println()

	for index := 1; index <= n; index++ {
		if index <= 1 {
			fmt.Print(term1)
		} else if index == 2 {
			fmt.Print(" ", term2)
		} else {
			nextTerm = term1 + term2
			term1 = term2
			term2 = nextTerm
			fmt.Print(" ", nextTerm)
		}
	}
	fmt.Println()
}

func main() {
	definedNumber := readIntFromInput()

	if definedNumber <= 0 || !isInteger(float64(definedNumber)) {
		fmt.Println("Nombre incorrect, veuillez saisir un nombre entier strictement positif.")
		return
	}

	printFibonacciSeries(definedNumber)
}
