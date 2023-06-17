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

type turn struct {
	antagonist  string
	protagonist string
}

func main() {

	scanner := streamLines("./input2.txt")
	arr := []turn{}
	antagonistPoints := map[string]int{"A": 1, "B": 2, "C": 3}

	for scanner.Scan() {
		arr = append(arr, turn{antagonist: string(scanner.Text()[0]), protagonist: string(scanner.Text()[2])})
	}

	// Part 1
	protagonistPoints := map[string]int{"X": 1, "Y": 2, "Z": 3}

	points := 0

	for _, i := range arr {
		result := antagonistPoints[i.antagonist] - protagonistPoints[i.protagonist]
		switch result {
		case 0:
			points += 3
		case -1, 2:
			points += 6
		case -2, 1:
			points += 0
		}
		points += protagonistPoints[i.protagonist]
	}
	fmt.Println(points)

	//Part 2
	protagonistPoints2 := map[string]int{"X": 0, "Y": 3, "Z": 6}

	points2 := 0

	for _, i := range arr {
		if protagonistPoints2[i.protagonist] == 0 {
			switch antagonistPoints[i.antagonist] {
			case 1:
				points2 += 3
			case 2:
				points2 += 1
			case 3:
				points2 += 2
			}
		} else if protagonistPoints2[i.protagonist] == 6 {
			switch antagonistPoints[i.antagonist] {
			case 1:
				points2 += 2
			case 2:
				points2 += 3
			case 3:
				points2 += 1
			}
		} else {
			points2 += antagonistPoints[i.antagonist]
		}
		points2 += protagonistPoints2[i.protagonist]
	}
	fmt.Println(points2)
}
