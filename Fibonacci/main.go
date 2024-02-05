package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func capitalizeFirstLetter(s string) string {
	if s == "" {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func readFromInput() (int, error) {
	var input string
	fmt.Print("Veuillez saisir un nombre entier : ")
	_, err := fmt.Scan(&input)

	if err != nil {
		return 0, fmt.Errorf(capitalizeFirstLetter("erreur lors de la lecture de l'entrée"))
	}

	if strings.Contains(input, ".") || strings.Contains(input, ",") {
		return 0, fmt.Errorf(capitalizeFirstLetter("votre nombre contient une virgule, veuillez saisir un nombre entier"))
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf(capitalizeFirstLetter("veuillez saisir un nombre entier valide"))
	}

	return num, nil
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
}

func main() {
	definedNumber, err := readFromInput()

	if err != nil {
		fmt.Print(err)
		return
	}

	if definedNumber <= 0 {
		fmt.Print("Nombre incorrect, veuillez saisir un nombre entier strictement positif")
		return
	}

	printFibonacciSeries(definedNumber)
}
