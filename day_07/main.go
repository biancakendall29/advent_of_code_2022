// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func streamLines(path string) bufio.Scanner {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return *bufio.NewScanner(file)
// }

// func main() {

// 	arr := [][]string{}
// 	values := map[string]int{}
// 	items := map[string][][]string{}
// 	directories := []string{}
// 	// upCount := 1
// 	scanner := streamLines("./input.txt")

// 	for scanner.Scan() {
// 		arr = append(arr, strings.Split(scanner.Text(), "\n"))
// 	}

// 	for _, i := range arr {
// 		dirName := changeDirectory(i)
// 		// if i[0][0:1] == "$" {
// 		// listingItems = false
// 		// if dirName != "/" {
// 		// 	addFileSizes(contents, directories[len(directories)-1], values)
// 		// }
// 		// contents = [][]string{}
// 		if dirName != "" {
// 			// if dirName == ".." {
// 			// 	upCount++
// 			// 	continue
// 			// }  else {
// 			directories = append(directories, dirName)
// 			// }
// 		} else if dirName == "" && !listContents(i) {
// 			if len(directories) > 0 {
// 				items[directories[len(directories)-1]] = append(items[directories[len(directories)-1]], i)
// 			}
// 		} else {
// 			// directory == directories[len(directories)-1]
// 			continue
// 		}
// 	}
// 	addSubDirectorySizes(values, items)
// 	fmt.Println(items)
// 	fmt.Println(values)
// 	findLessThanLimit(values)
// }

// func findLessThanLimit(values map[string]int) {
// 	count := 0
// 	total := 0
// 	for in, v := range values {
// 		if in != "/" {

// 			total += v
// 		}
// 		if v < 100000 {
// 			count += v
// 		}
// 		// fmt.Println(v)
// 	}
// 	fmt.Println(count)
// 	fmt.Println(total)
// }

// func addSubDirectorySizes(values map[string]int, items map[string][][]string) {
// 	m := map[string][]string{}
// 	for k, i := range items {
// 		if len(i) > 0 {
// 			for _, j := range i {
// 				// fmt.Println(k, j)
// 				if j[0][0:3] == "dir" {
// 					m[k] = append(m[k], j[0][4:])
// 				} else {
// 					size := strings.Split(j[0], " ")
// 					sizeInt, err := strconv.Atoi(size[0])
// 					if err != nil {
// 						fmt.Println("Error: can't convert to int")
// 					} else {
// 						// values[directory] += sizeInt
// 						// fmt.Println(k)
// 						values[k] += sizeInt
// 					}
// 				}
// 			}
// 		}
// 		// fmt.Println(m)

// 	}
// 	for key, val := range m {
// 		// fmt.Println(key, val)
// 		for _, val2 := range val {
// 			values[key] += values[val2]
// 			// fmt.Println(values)
// 			// fmt.Println(key, values[key], val, values[val2], val2)
// 		}
// 	}
// }

// // func addFileSizes(contents [][]string, directory string, values map[string]int) {
// // 	for _, i := range contents {
// // 		if i[0][0:3] == "dir" {
// // 			continue
// // 		} else {
// // 			size := strings.Split(i[0], " ")
// // 			sizeInt, err := strconv.Atoi(size[0])
// // 			if err != nil {
// // 				fmt.Println("Error: can't convert to int")
// // 			} else {
// // 				values[directory] += sizeInt
// // 			}
// // 		}
// // 	}
// // }

// func changeDirectory(i []string) string {
// 	if i[0][0:1] == "$" {
// 		if i[0][2:4] == "cd" {
// 			return i[0][5:]
// 		}
// 	}
// 	return ""
// }

// func listContents(i []string) bool {
// 	if i[0][0:1] == "$" {
// 		if i[0][2:4] == "ls" {
// 			return true
// 		}
// 	}
// 	return false
// }

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

type Directory struct {
	name     string
	files    []File
	children []*Directory
	parent   *Directory
}

type File struct {
	name string
	size int
}

var totalUnderLimit int
var deletables = make([]int, 0)
var trackSpace int

const DiskSpace = 70000000
const FreeSpace = 30000000

func main() {

	arr := [][]string{}
	lines := [][]string{}
	var root *Directory
	var current *Directory
	scanner := streamLines("./input.txt")

	for scanner.Scan() {
		arr = append(arr, strings.Split(scanner.Text(), "\n"))
	}
	for _, i := range arr {
		lines = append(lines, strings.Split(i[0], " "))
	}

	for _, j := range lines {
		if j[0] == "$" {
			if j[1] == "cd" {
				if j[2] == ".." {
					current = current.parent
				} else if j[2] == "/" {
					current = &Directory{
						name:     j[2],
						files:    []File{},
						children: []*Directory{},
						parent:   &Directory{},
					}
					root = current
				} else {
					current = current.changeDir(j[2])
				}
			}
		} else {
			if j[0] == "dir" {
				current.addDir(j[1])
			} else {
				current.addFile(j[0], j[1])
			}
		}

	}
	if root != nil {
		// Part 1
		root.getSize()
		fmt.Println(totalUnderLimit)

		// Part 2
		trackSpace = FreeSpace - (DiskSpace - root.getSize())
		root.getSize()

		chosen := deletables[0]
		for i, e := range deletables {
			if i == 0 || e < chosen {
				chosen = e
			}
		}
		fmt.Println(chosen)
	}

}

func (dir *Directory) addDir(name string) {
	if dir.children == nil {
		var dirs []*Directory
		dirs = append(dirs, &Directory{
			name:   name,
			parent: dir,
		})
	}
	dir.children = append(dir.children, &Directory{
		name:   name,
		parent: dir,
	})
}

func (dir *Directory) addFile(size string, name string) {
	num, err := strconv.Atoi(size)
	if err != nil {
		fmt.Println("Could not convert to number")
	} else if dir.files == nil {
		files := []File{}
		files = append(files, File{
			name: name,
			size: num,
		})
		dir.files = files
	} else {
		dir.files = append(dir.files, File{
			name: name,
			size: num,
		})
	}
}

func (dir *Directory) changeDir(name string) *Directory {
	for _, dir := range dir.children {
		if dir.name == name {
			return dir
		}
	}
	return nil
}

func (dir *Directory) getSize() int {
	sum := 0
	for _, file := range dir.files {
		sum += file.size
	}

	for _, dir := range dir.children {
		sum += dir.getSize()
	}

	// Part 1
	if sum <= 100000 {
		totalUnderLimit += sum
	}

	// Part 2

	if sum >= trackSpace && trackSpace > 0 {
		deletables = append(deletables, sum)
	}

	return sum
}
