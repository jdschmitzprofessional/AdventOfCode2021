package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type chunk struct {
	depth1, depth2, depth3 int
}

func (c *chunk) sum() int {
	return c.depth3 + c.depth2 + c.depth1
}

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
	currentDepth := chunk{0, 0, 0}
	lastDepth := chunk{0, 0, 0}
	increases := 0
	fmt.Println(currentDepth.sum())
	for i := range lines {
		lastDepth = currentDepth
		if i+2 == 2000 {
			break
		}
		currentDepth = chunk{depth1: lines[i], depth2: lines[i+1], depth3: lines[i+2]}
		if i == 0 {
			continue
		}

		currentDepth = chunk{depth1: lines[i], depth2: lines[i+1], depth3: lines[i+2]}
		fmt.Println(currentDepth.sum(), lastDepth.sum())
		if currentDepth.sum() > lastDepth.sum() {
			increases += 1
		}
	}
	fmt.Println(increases)
}
