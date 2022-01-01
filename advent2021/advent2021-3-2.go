package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(arr []int) int {
	tot := 0
	for i := range arr {
		tot += i
	}
	return tot
}

const bin_length int = 12
const set_length int = 1000

func main() {
	var lines [][]string
	var output []int
	var gammaRateArray []string
	var epsilonRateArray []string

	file, err := os.Open("advent2021/resources/advent2021-3.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		int_arr := strings.Split(val, "")
		lines = append(lines, int_arr)
	}

	tot := 0
	for index1 := 0; index1 < bin_length; index1++ {
		tot = 0
		for index2 := 0; index2 < set_length; index2++ {
			integerVal, _ := strconv.Atoi(lines[index2][index1])
			tot += integerVal
		}
		output = append(output, tot)
	}
	for index := range output {
		if output[index] >= 500 {
			gammaRateArray = append(gammaRateArray, "1")
			epsilonRateArray = append(epsilonRateArray, "0")
		} else {
			gammaRateArray = append(gammaRateArray, "0")
			epsilonRateArray = append(epsilonRateArray, "1")
		}
	}
	epsilonString := strings.Join(epsilonRateArray, "")
	gammaString := strings.Join(gammaRateArray, "")
	epsilonInt, err := strconv.ParseInt(epsilonString, 2, 64)
	if err != nil {
		fmt.Println("o no")
	}
	gammaInt, err := strconv.ParseInt(gammaString, 2, 64)
	if err != nil {
		fmt.Println("o no")
	}
	fmt.Println("Epsilon Rate: ", epsilonInt)
	fmt.Println("Gamma Rate: ", gammaInt)
	fmt.Println("Power Consumption: ", epsilonInt*gammaInt)
}
