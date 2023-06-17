package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func streamLines(path string) bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return *bufio.NewScanner(file)
}

func main() {

	scanner := streamLines("./input.txt")
	arr := [][]string{}
	count := 0
	reverseCount := 0

	for scanner.Scan() {
		arr = append(arr, strings.Split(scanner.Text(), "\n"))
	}

	for index, i := range arr {
		arr[index] = strings.Split(i[0], ",")
		firstArg := strings.Split(arr[index][0], "-")
		secondArg := strings.Split(arr[index][1], "-")
		firstLower, _ := strconv.Atoi(firstArg[0])
		firstUpper, _ := strconv.Atoi(firstArg[1])
		secondLower, _ := strconv.Atoi(secondArg[0])
		secondUpper, _ := strconv.Atoi(secondArg[1])

		// Part 1

		if (firstLower <= secondLower && firstUpper >= secondUpper) ||
			(secondLower <= firstLower && secondUpper >= firstUpper) {
			count++
		}

		// Part 2

		if (firstUpper < secondLower && firstLower < secondLower) ||
			(secondUpper < firstLower && secondLower < firstLower) {
			reverseCount++
		}
	}
	// Part 1
	fmt.Println(count)
	// Part 2
	fmt.Println(len(arr) - reverseCount)

}
