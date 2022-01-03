package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
	z int
}

// link: https://adventofcode.com/2021/day/2
func main() {
	file, err := os.Open("advent2021/resources/advent2021-2.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		lines = append(lines, val)
	}
	pos := position{0, 0, 0}
	for index := range lines {
		line := strings.Split(lines[index], " ")
		value, _ := strconv.Atoi(line[1])
		switch line[0] {
		case "down":
			pos.z = pos.z + value
		case "up":
			pos.z = pos.z - value
		case "forward":
			pos.x = pos.x + value
		}
		fmt.Println(pos.x, pos.y, pos.z)
	}
	fmt.Println(pos.x * pos.z)
}
