package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	forward int
	depth   int
	aim     int
}

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
			pos.aim = pos.aim + value
		case "up":
			pos.aim = pos.aim - value
		case "forward":
			pos.forward = pos.forward + value
			pos.depth = pos.depth + int(pos.aim*value)
		}
		fmt.Println(pos.forward, pos.depth, pos.aim)
	}
	fmt.Println(pos.forward * pos.depth)
}
