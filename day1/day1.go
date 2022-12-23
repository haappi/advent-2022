package main

import (
	"fmt"
	"os"
	"strconv"
	_ "strconv"
	"strings"
)

func main() {
	fileBytesData, err := os.ReadFile("day1/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	splittedData := strings.Split(string(fileBytesData), "\n")

	var perElf []int

	var tempCounter int

	for _, line := range splittedData {
		line = strings.ReplaceAll(line, "\r", "")
		i, err := strconv.Atoi(line)
		if err != nil {
			perElf = append(perElf, tempCounter)
			tempCounter = 0
			continue
		}
		tempCounter += i
	}

	// print the largest one
	var largest int
	for _, i := range perElf {
		if i > largest {
			largest = i
		}
	}

	var secondLargest int
	for _, i := range perElf {
		if i > secondLargest && i < largest {
			secondLargest = i
		}
	}

	var thirdLargest int
	for _, i := range perElf {
		if i > thirdLargest && i < secondLargest {
			thirdLargest = i
		}
	}

	fmt.Println(largest + secondLargest + thirdLargest)

}
