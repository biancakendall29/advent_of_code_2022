package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func streamLines(path string) bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return *bufio.NewScanner(file)
}

type Position struct {
	row    int
	column int
}

func main() {
	grid := getGrid()
	length := len(grid)
	var blocker bool
	visibilityMap := map[Position]int{}

	checkVisRows(grid, length, visibilityMap, blocker)
	fmt.Println(len(visibilityMap) + (3 * length) + (length - 1))

}

func checkVisRows(grid [][]int, length int, visibilityMap map[Position]int, blocker bool) {
	for rowInd, row := range grid {
		if rowInd == 0 {
			continue
		} else if rowInd == length-1 {
			break
		} else {
			// check rows
			for colInd, item := range row {
				blockerArr := []bool{}
				if colInd == 0 {
					continue
				} else if colInd == length-1 {
					break
				} else {
					// check each row to the right
					checkLeft(colInd, length, row, item, rowInd, blockerArr, visibilityMap)
					// check each row to the left
					checkRight(colInd, length, row, item, rowInd, blockerArr, visibilityMap)
					// check each column, downwards
					checkDown(colInd, length, grid, item, rowInd, blockerArr, visibilityMap)
					// check each column, upwards
					checkUp(colInd, length, grid, item, rowInd, blockerArr, visibilityMap)
				}
			}
		}
	}
}

func checkLeft(colInd int, length int, row []int, item int, rowInd int, blockerArr []bool, visibilityMap map[Position]int) {
	for i := colInd + 1; i < length-1; i++ {
		if row[i] >= item {
			blockerArr = append(blockerArr, true)
		} else {
			blockerArr = append(blockerArr, false)
		}
		// fmt.Println(row[i], item, i, blockerArr)
		if contains(blockerArr, true) == false && i == length-2 {
			visibilityMap[Position{row: rowInd, column: colInd}]++
		}
	}
}

func checkRight(colInd int, length int, row []int, item int, rowInd int, blockerArr []bool, visibilityMap map[Position]int) {
	for i := colInd - 1; i > 0; i-- {
		if row[i] >= item {
			blockerArr = append(blockerArr, true)
		} else {
			blockerArr = append(blockerArr, false)
		}
		if contains(blockerArr, true) == false && i == 1 {
			visibilityMap[Position{row: rowInd, column: colInd}]++
		}
	}
}

func checkDown(colInd int, length int, grid [][]int, item int, rowInd int, blockerArr []bool, visibilityMap map[Position]int) {
	for i := rowInd; i < length-1; i++ {
		if grid[i+1][colInd] >= item {
			blockerArr = append(blockerArr, true)
		} else {
			blockerArr = append(blockerArr, false)
		}
		if contains(blockerArr, true) == false && i == length-2 {
			visibilityMap[Position{row: rowInd, column: colInd}]++
		}
	}
}

func checkUp(colInd int, length int, grid [][]int, item int, rowInd int, blockerArr []bool, visibilityMap map[Position]int) {
	for i := rowInd; i > 0; i-- {
		if grid[i-1][colInd] >= item {
			blockerArr = append(blockerArr, true)
		} else {
			blockerArr = append(blockerArr, false)
		}
		if contains(blockerArr, true) == false && i == 1 {
			visibilityMap[Position{row: rowInd, column: colInd}]++
		}
	}
}

func contains(arr []bool, e bool) bool {
	for _, i := range arr {
		if i == e {
			return true
		}
	}
	return false
}

func getGrid() [][]int {
	scanner := streamLines("./input.txt")
	arr := [][]string{}
	grid := [][]int{}

	for scanner.Scan() {
		arr = append(arr, strings.Split(scanner.Text(), "\n"))
	}

	for _, i := range arr {
		for _, j := range i {
			rowInt := []int{}
			row := []rune(j)
			for _, k := range row {
				rowInt = append(rowInt, int(k)-48)

			}
			grid = append(grid, rowInt)
		}
	}
	return grid
}
