package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	// question 1

	var tailHasVisitedRepeated [][2]int
	headPosition := [2]int{0, 0}
	tailPosition := [2]int{0, 0}
	tailHasVisitedRepeated = append(tailHasVisitedRepeated, tailPosition)
	for scanner.Scan() {
		move := string(scanner.Text()[0])
		numberOfMove, _ := strconv.Atoi(strings.Split(string(scanner.Text()), " ")[1])
		switch move {
		case "R":
			for moveNo := 1; moveNo <= numberOfMove; moveNo++ {
				headPosition[0] = headPosition[0] + 1
				if !canTailRemain(headPosition, tailPosition) {
					tailPosition[0] = tailPosition[0] + 1
					if tailPosition[1] == headPosition[1]-1 {
						tailPosition[1] = tailPosition[1] + 1
					} else if tailPosition[1] == headPosition[1]+1 {
						tailPosition[1] = tailPosition[1] - 1
					}
				}
				tailHasVisitedRepeated = append(tailHasVisitedRepeated, tailPosition)
			}
		case "L":
			for moveNo := 1; moveNo <= numberOfMove; moveNo++ {
				headPosition[0] = headPosition[0] - 1
				if !canTailRemain(headPosition, tailPosition) {
					tailPosition[0] = tailPosition[0] - 1
					if tailPosition[1] == headPosition[1]-1 {
						tailPosition[1] = tailPosition[1] + 1
					} else if tailPosition[1] == headPosition[1]+1 {
						tailPosition[1] = tailPosition[1] - 1
					}
				}
				tailHasVisitedRepeated = append(tailHasVisitedRepeated, tailPosition)
			}
		case "U":
			for moveNo := 1; moveNo <= numberOfMove; moveNo++ {
				headPosition[1] = headPosition[1] + 1
				if !canTailRemain(headPosition, tailPosition) {
					tailPosition[1] = tailPosition[1] + 1
					if tailPosition[0] == headPosition[0]-1 {
						tailPosition[0] = tailPosition[0] + 1
					} else if tailPosition[0] == headPosition[0]+1 {
						tailPosition[0] = tailPosition[0] - 1
					}
				}
				tailHasVisitedRepeated = append(tailHasVisitedRepeated, tailPosition)
			}
		case "D":
			for moveNo := 1; moveNo <= numberOfMove; moveNo++ {
				headPosition[1] = headPosition[1] - 1
				if !canTailRemain(headPosition, tailPosition) {
					tailPosition[1] = tailPosition[1] - 1
					if tailPosition[0] == headPosition[0]-1 {
						tailPosition[0] = tailPosition[0] + 1
					} else if tailPosition[0] == headPosition[0]+1 {
						tailPosition[0] = tailPosition[0] - 1
					}
				}
				tailHasVisitedRepeated = append(tailHasVisitedRepeated, tailPosition)
			}
		}
	}
	var tailsHasVisited [][2]int
	var alreadyExists bool
	for _, element := range tailHasVisitedRepeated {
		for _, check := range tailsHasVisited {
			if element == check {
				alreadyExists = true
			}
		}
		if !alreadyExists {
			tailsHasVisited = append(tailsHasVisited, element)
		}
		alreadyExists = false
	}
	fmt.Println(len(tailsHasVisited))

	// question 2

	// 	var positions [10][2]int
	// 	var ninthTailHasVisitedRepeated [][2]int

	// 	for scanner.Scan() {
	// 		move := string(scanner.Text()[0])
	// 		numberOfMove, _ := strconv.Atoi(strings.Split(string(scanner.Text()), " ")[1])
	// 		switch move {
	// 		case "R":
	// 			for moveNo := 1; moveNo <= numberOfMove; moveNo++ {
	// 				for index := range positions {
	// 					if index == 0 {
	// 						positions[0][0] = positions[0][0] + 1
	// 					}
	// 					if index == 9 {
	// 						break
	// 					}
	// 					positions[index+1] = findMove(positions[index], positions[index+1])
	// 				}
	// 				ninthTailHasVisitedRepeated = append(ninthTailHasVisitedRepeated, positions[9])
	// 			}
	// 		case "L":
	// 			for moveNo := 1; moveNo <= numberOfMove; moveNo++ {
	// 				for index := range positions {
	// 					if index == 0 {
	// 						positions[0][0] = positions[0][0] - 1
	// 					}
	// 					if index == 9 {
	// 						break
	// 					}
	// 					positions[index+1] = findMove(positions[index], positions[index+1])
	// 				}
	// 				ninthTailHasVisitedRepeated = append(ninthTailHasVisitedRepeated, positions[9])
	// 			}
	// 		case "U":
	// 			for moveNo := 1; moveNo <= numberOfMove; moveNo++ {
	// 				for index := range positions {
	// 					if index == 0 {
	// 						positions[0][1] = positions[0][1] + 1
	// 					}
	// 					if index == 9 {
	// 						break
	// 					}
	// 					positions[index+1] = findMove(positions[index], positions[index+1])
	// 				}
	// 				ninthTailHasVisitedRepeated = append(ninthTailHasVisitedRepeated, positions[9])
	// 			}
	// 		case "D":
	// 			for moveNo := 1; moveNo <= numberOfMove; moveNo++ {
	// 				for index := range positions {
	// 					if index == 0 {
	// 						positions[0][1] = positions[0][1] - 1
	// 					}
	// 					if index == 9 {
	// 						break
	// 					}
	// 					positions[index+1] = findMove(positions[index], positions[index+1])
	// 				}
	// 				ninthTailHasVisitedRepeated = append(ninthTailHasVisitedRepeated, positions[9])
	// 			}
	// 		}
	// 	}
	// 	var ninthTailHasVisited [][2]int
	// 	var alreadyExists bool
	// 	for _, element := range ninthTailHasVisitedRepeated {
	// 		for _, check := range ninthTailHasVisited {
	// 			if element == check {
	// 				alreadyExists = true
	// 			}
	// 		}
	// 		if !alreadyExists {
	// 			ninthTailHasVisited = append(ninthTailHasVisited, element)
	// 		}
	// 		alreadyExists = false
	// 	}
	// 	fmt.Println(len(ninthTailHasVisited))
	// }

	// func findMove(head [2]int, tail [2]int) [2]int {
	// 	if tail[1] == head[1] && tail[0] == head[0]-2 {
	// 		tail[0] = tail[0] + 1
	// 	}
	// 	if tail[1] == head[1] && tail[0] == head[0]+2 {
	// 		tail[0] = tail[0] - 1
	// 	}
	// 	if tail[0] == head[0] && tail[1] == head[1]+2 {
	// 		tail[1] = tail[1] - 1
	// 	}
	// 	if tail[0] == head[0] && tail[1] == head[1]-2 {
	// 		tail[1] = tail[1] + 1
	// 	}
	// 	if tail[0] == head[0]-1 && tail[1] == head[1]-2 {
	// 		tail[0] = tail[0] + 1
	// 		tail[1] = tail[1] + 1
	// 	}
	// 	if tail[0] == head[0]-1 && tail[1] == head[1]+2 {
	// 		tail[0] = tail[0] + 1
	// 		tail[1] = tail[1] - 1
	// 	}
	// 	if tail[0] == head[0]+1 && tail[1] == head[1]-2 {
	// 		tail[0] = tail[0] - 1
	// 		tail[1] = tail[1] + 1
	// 	}
	// 	if tail[0] == head[0]+1 && tail[1] == head[1]+2 {
	// 		tail[0] = tail[0] - 1
	// 		tail[1] = tail[1] - 1
	// 	}
	// 	if tail[0] == head[0]-2 && tail[1] == head[1]+1 {
	// 		tail[0] = tail[0] + 1
	// 		tail[1] = tail[1] - 1
	// 	}
	// 	if tail[0] == head[0]-2 && tail[1] == head[1]-1 {
	// 		tail[0] = tail[0] + 1
	// 		tail[1] = tail[1] + 1
	// 	}
	// 	if tail[0] == head[0]+2 && tail[1] == head[1]-1 {
	// 		tail[0] = tail[0] - 1
	// 		tail[1] = tail[1] + 1
	// 	}
	// 	if tail[0] == head[0]+2 && tail[1] == head[1]+1 {
	// 		tail[0] = tail[0] - 1
	// 		tail[1] = tail[1] - 1
	// 	}
	// 	if tail[0] == head[0]-2 && tail[1] == head[1]+2 {
	// 		tail[0] = tail[0] + 1
	// 		tail[1] = tail[1] - 1
	// 	}
	// 	if tail[0] == head[0]-2 && tail[1] == head[1]-2 {
	// 		tail[0] = tail[0] + 1
	// 		tail[1] = tail[1] + 1
	// 	}
	// 	if tail[0] == head[0]+2 && tail[1] == head[1]-2 {
	// 		tail[0] = tail[0] - 1
	// 		tail[1] = tail[1] + 1
	// 	}
	// 	if tail[0] == head[0]+2 && tail[1] == head[1]+2 {
	// 		tail[0] = tail[0] - 1
	// 		tail[1] = tail[1] - 1
	// 	}
	// 	return tail
}

// question 1
func canTailRemain(headPosition [2]int, tailPosition [2]int) bool {
	if headPosition == tailPosition {
		return true
	} else if headPosition[0] == tailPosition[0] {
		if headPosition[1] == tailPosition[1]+1 || headPosition[1] == tailPosition[1]-1 {
			return true
		}
	} else if headPosition[1] == tailPosition[1] {
		if headPosition[0] == tailPosition[0]-1 || headPosition[0] == tailPosition[0]+1 {
			return true
		}
	} else if headPosition[0] == tailPosition[0]+1 {
		if headPosition[1] == tailPosition[1]-1 || headPosition[1] == tailPosition[1]+1 {
			return true
		}
	} else if headPosition[0] == tailPosition[0]-1 {
		if headPosition[1] == tailPosition[1]-1 || headPosition[1] == tailPosition[1]+1 {
			return true
		}
	} else {
		return false
	}
	return false
}
