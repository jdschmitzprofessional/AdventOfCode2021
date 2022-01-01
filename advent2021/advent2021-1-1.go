package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
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
	increases := 0
	current_depth := 0
	last_depth := 0
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
	fmt.Println(increases)
}
