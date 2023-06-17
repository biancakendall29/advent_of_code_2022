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

func main() {

	arr := ""
	scanner := streamLines("./input.txt")

	for scanner.Scan() {
		arr = scanner.Text()
	}

	// Part 1

	fourPack := [4]rune{}
	lastMarker4 := 0

	for index, i := range arr {

		if index < 4 {
			fourPack[index] = i
		} else if checkPacket4(fourPack) == true {
			lastMarker4 = index
			break
		} else {
			for j := 0; j < 3; j++ {
				fourPack[j] = fourPack[j+1]
			}
			fourPack[3] = rune(arr[index])
		}
	}
	fmt.Println("Packet (4)", fourPack)
	fmt.Println("Marker for 4 pack:", lastMarker4)

	// Part 2

	fourteenPack := [14]rune{}
	lastMarker14 := 0

	for index, i := range arr {

		if index < 14 {
			fourteenPack[index] = i
		} else if checkPacket14(fourteenPack) == true {
			lastMarker14 = index
			break
		} else {
			for j := 0; j < 13; j++ {
				fourteenPack[j] = fourteenPack[j+1]
			}
			fourteenPack[13] = rune(arr[index])
		}
	}
	fmt.Println("Packet (14)", fourteenPack)
	fmt.Println("Marker for 14 pack:", lastMarker14)
}

func checkPacket4(arr [4]rune) bool {
	m := map[rune]int{}
	for k := 0; k < 4; k++ {
		m[arr[k]]++
	}
	if len(m) == 4 {
		return true

	} else {
		return false
	}
}

func checkPacket14(arr [14]rune) bool {
	m := map[rune]int{}
	for k := 0; k < 14; k++ {
		m[arr[k]]++
	}
	if len(m) == 14 {
		return true

	} else {
		return false
	}
}
