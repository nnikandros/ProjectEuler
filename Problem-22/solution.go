package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var alphabet = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

func main() {
	res, err := os.ReadFile("/ec/local/shared2/user/nikanni/my_go_files/names.txt")
	if err != nil {
		fmt.Println("coundn't read ")
	}

	slice_of_strings := strings.Split(string(res), ",")

	slices.Sort(slice_of_strings)

	for index, element := range slice_of_strings {
		slice_of_strings[index] = strings.Trim(element, "\"")
	}

	totalSum := 0

	for index, name := range slice_of_strings {
		totalSum += calculateWorthwithIndex(name, index)
	}

	fmt.Println(totalSum)

}

func calculateWorthwithIndex(s string, index int) int {

	sumOfLetters := 0
	for _, element := range s {
		sumOfLetters += alphabet[string(element)]
	}

	return sumOfLetters * (index + 1)

}
