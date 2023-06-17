package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func streamLines(path string) bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return *bufio.NewScanner(file)
}

func Equal(a, b []rune) rune {
	for _, x := range a {
		for _, y := range b {
			if x != y {
				continue
			} else {
				return x
			}
		}
	}
	return rune(0)
}

func Equal3(a, b []rune) []rune {
	arr := []rune{}
	for _, x := range a {
		for _, y := range b {
			if x != y {
				continue
			} else {
				arr = append(arr, x)
			}
		}
	}
	return arr
}

func ConvertPoints(arr []rune) int {
	points := 0
	for _, i := range arr {
		if i <= 'z' && i >= 'a' {
			points += int(i) - 96
		} else if i <= 'Z' && i >= 'A' {
			points += int(i) - 38
		}
	}
	return points
}

func main() {

	scanner := streamLines("./input.txt")
	arr := []string{}

	for scanner.Scan() {
		arr = append(arr, string(scanner.Text()))
	}

	// Part 1
	isEven := true
	common := []rune{}
	for _, i := range arr {
		length := len(i)
		firstHalf := []rune{}
		secondHalf := []rune{}
		if length%2 == 0 {
			isEven = true
		}
		for index, j := range i {
			if isEven == true && index < length/2 {
				firstHalf = append(firstHalf, j)
			} else if isEven == false {
				fmt.Println("Odd length")
			} else {
				secondHalf = append(secondHalf, j)
			}
		}
		commonRune := Equal(firstHalf, secondHalf)
		if commonRune != ' ' {
			common = append(common, commonRune)
		}
	}
	fmt.Println(ConvertPoints(common))

	// Part 2
	common3 := []rune{}
	for index, i := range arr {
		if index+1 == 1 || index+1 == 2 {
			continue
		} else if index+1 == 3 {
			compareArr := Equal3([]rune(i), []rune(arr[index-1]))

			common3 = append(common3, Equal(compareArr, []rune(arr[index-2])))

		} else if index+1 >= 3 && (index+1)%3 == 0 {
			compareArr := Equal3([]rune(i), []rune(arr[index-1]))

			common3 = append(common3, Equal(compareArr, []rune(arr[index-2])))

		}
	}
	fmt.Println(ConvertPoints(common3))
}
