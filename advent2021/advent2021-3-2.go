package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const binaryLength int = 12
const dataSetLength int = 1000

// link: https://adventofcode.com/2021/day/3
func main() {
	var inputLines [][]string
	var parsedOutput []int
	var gammaRateArray []string
	var epsilonRateArray []string

	// read input data provided into an array of arrays of strings.
	file, err := os.Open("advent2021/resources/advent2021-3.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		int_arr := strings.Split(val, "")
		inputLines = append(inputLines, int_arr)
	}

	// get sum of VERTICAL bits
	tot := 0
	for index1 := 0; index1 < binaryLength; index1++ {
		tot = 0
		for index2 := 0; index2 < dataSetLength; index2++ {
			integerVal, err := strconv.Atoi(inputLines[index2][index1])
			if err != nil {
				fmt.Println(err.Error())
			}
			tot += integerVal
		}
		parsedOutput = append(parsedOutput, tot)
	}

	// read each of the output numbers and determine the new binary value of epsilon/gamma rates as per day 3 part 2
	// instructions.

	for index := range parsedOutput {
		if parsedOutput[index] >= 501 {
			gammaRateArray = append(gammaRateArray, "1")
			epsilonRateArray = append(epsilonRateArray, "0")
		} else {
			gammaRateArray = append(gammaRateArray, "0")
			epsilonRateArray = append(epsilonRateArray, "1")
		}
	}

	// convert above string arrays into binary numbers, then convert to integer
	epsilonString := strings.Join(epsilonRateArray, "")
	gammaString := strings.Join(gammaRateArray, "")
	epsilonInt, err := strconv.ParseInt(epsilonString, 2, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	gammaInt, err := strconv.ParseInt(gammaString, 2, 64)
	if err != nil {
		fmt.Println(err.Error())
	}

	// print output
	fmt.Println("Epsilon Rate: ", epsilonInt)
	fmt.Println("Gamma Rate: ", gammaInt)
	fmt.Println("Power Consumption: ", epsilonInt*gammaInt)
}
