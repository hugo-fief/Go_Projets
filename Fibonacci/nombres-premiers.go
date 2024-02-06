package main

import (
	"fmt"
)

func sieveOfEratosthenes(numb int) []int {
	prime := make([]bool, numb+1)

	for i := range prime {
		prime[i] = true
	}

	for p := 2; p*p <= numb; p++ {
		if prime[p] == true {
			for i := p * p; i <= numb; i += p {
				prime[i] = false
			}
		}
	}

	var primes []int
	for p := 2; p <= numb; p++ {
		if prime[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

func main() {
	var definedNumber int
	fmt.Print("Veuillez insérer un nombre entier : ")
	fmt.Scan(&definedNumber)

	if definedNumber < 2 {
		fmt.Print("Nombre incorrect. Veuillez saisir un nombre entier positif supérieur ou égal à 2")
		return
	}

	primes := sieveOfEratosthenes(definedNumber)
	fmt.Print("Nombres premiers : ", primes)
}
