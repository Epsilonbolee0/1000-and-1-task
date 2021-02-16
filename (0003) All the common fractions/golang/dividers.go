package main

import (
	"fmt"
	"math"
)
// Найти все несократимые дроби от 0 до 1, где знаменатель не превышает n

type Fraction struct {
	numerator, denominator int
}

func (pair *Fraction) print() {
	fmt.Printf("%d/%d\n", pair.numerator, pair.denominator)
}

func printNumbers(pairs []Fraction) {
	for _, pair := range pairs {
		pair.print()
	}
}

func findAllCommonFractions(upperBorder int) (commonFractions []Fraction) {
	for i := 2; i <= upperBorder; i++ {
		if numberIsSimple(i) {
			commonFractions = findCommonFractionsPrimeDenomenator(commonFractions, i)
		} else {
			commonFractions = findCommonFractionsCompositeDenominator(commonFractions, i)
		}
	}

	return
}

func numberIsSimple(number int) bool {
	check := true
	for i := 2; i < checkUpperBorder(number); i++ {
		if number % i == 0 {
			check = false
			break
		}
	}

	return check
}

func checkUpperBorder(number int) int {
	return (int)(math.Sqrt((float64)(number)))
}



func findCommonFractionsPrimeDenomenator(commonFractions []Fraction, divisor int) []Fraction {
	for i := 1; i < divisor; i++ {
		fraction := Fraction{i, divisor}
		commonFractions = append(commonFractions, fraction)
	}

	return commonFractions
}

func findCommonFractionsCompositeDenominator(commonFractions []Fraction, divisor int) []Fraction {
	for i := 1; i < divisor; i++ {
		if greatestCommonDivisor(divisor, i) == 1 {
			fraction := Fraction{i, divisor}
			commonFractions = append(commonFractions, fraction)
		}
	}

	return commonFractions
}



func greatestCommonDivisor(divisor, divisible int) int {
	for divisible != 0 {
		divisor, divisible = divisible, divisor % divisible
	}

	return divisor
}

func main() {
	numbers := findAllCommonFractions(1700)
	printNumbers(numbers)
}

