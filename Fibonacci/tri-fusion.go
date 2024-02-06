package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var definedSeries string
	fmt.Print("Veuillez saisir une série de nombres séparés par des espaces : ")
	reader := bufio.NewReader(os.Stdin)
	definedSeries, err := reader.ReadString('\n')

	if err != nil {
		fmt.Print("Série de nombre incorrects, veuillez séparés vos nombre par des espaces")
		return
	}

	definedSeries = definedSeries[:len(definedSeries)-1]

	numbers := convertInputToSlice(definedSeries)
	sortedNumbers := mergeSort(numbers)
	fmt.Print("Nombres triés : ", sortedNumbers)
}

func convertInputToSlice(definedSeries string) []int {
	stringSlice := strings.Fields(definedSeries)

	var intSlice []int
	for _, s := range stringSlice {
		num, err := strconv.Atoi(s)

		if err == nil {
			intSlice = append(intSlice, num)
		}
	}

	return intSlice
}

func mergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	mid := len(slice) / 2
	left := mergeSort(slice[:mid])
	right := mergeSort(slice[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}
