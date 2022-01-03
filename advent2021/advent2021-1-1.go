package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// link: https://adventofcode.com/2021/day/1
func main() {
	// instantiate variables
	increases := 0
	current_depth := 0
	last_depth := 0

	// open file and read in puzzle input
	file, err := os.Open("advent2021/resources/advent2021-1.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		int_val, _ := strconv.Atoi(val)
		lines = append(lines, int_val)
	}

	// parse input, increment increases when the number increases
	for i := range lines {
		if i == 0 {
			current_depth = lines[i]
			continue
		}
		last_depth = current_depth
		current_depth = lines[i]
		if current_depth > last_depth {
			increases += 1
		}
	}

	// print output
	fmt.Println(increases)
}
