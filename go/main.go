package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var input = bufio.NewReader(os.Stdin)

func printBox() {
	println("+" + strings.Repeat("-", 78) + "+\n" +
		"|" + strings.Repeat(" ", 33) + "\033[7mKarlKulator\033[0m" + strings.Repeat(" ", 34) + "|\n" +
		"+" + strings.Repeat("-", 78) + "+")
}

func prompt(text string) string {
	print(text + ":\n> ")
	text, err := input.ReadString('\n')
	if err != nil {
		log.Fatalf("You fucked up")
	}
	return strings.TrimSuffix(text, "\n")
}

func printSolution(number1 float64, operation string, number2 float64, solution float64) {
	fmt.Printf("\n%.5f %s %.5f = \033[1m\033[93m%.5f\033[0m", number1, operation, number2, solution)
}

func parseNumbers(number1 string, number2 string) (float64, float64, error) {
	res1, err := strconv.ParseFloat(number1, 64)
	if err != nil {
		return 0, 0, err
	}

	res2, err := strconv.ParseFloat(number2, 64)
	if err != nil {
		return 0, 0, err
	}

	return res1, res2, nil
}

func main() {
	printBox()

	input1 := prompt("Enter number 1")
	operation := prompt("Enter operation")
	input2 := prompt("Enter number 2")

	number1, number2, err := parseNumbers(input1, input2)
	if err != nil {
		log.Fatalln("Error", err)
	}

	switch operation {
	case "+":
		printSolution(number1, operation, number2, number1+number2)
	case "-":
		printSolution(number1, operation, number2, number1-number2)
	case "*":
		printSolution(number1, operation, number2, number1*number2)
	case "/":
		printSolution(number1, operation, number2, number1/number2)
	default:
		fmt.Printf("%+v", operation)
		log.Fatalln("Error")
	}
}
