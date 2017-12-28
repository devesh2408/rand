package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Printf("usage :\n %s count\n", os.Args[0])
		return
	}
	totalRandomWords := 1
	var err error
	if len(os.Args) == 2 {
		totalRandomWords, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error : Invalid Argument ", err)
			return
		}
	}

	rand.Seed(time.Now().UTC().UnixNano())
	wordArray, err := readFile("../words_alpha.txt")
	if err != nil {
		fmt.Println("Error: can't read the file ", err)
		return
	}

	totalWords := len(wordArray)
	if totalWords < totalRandomWords {
		fmt.Println("Warning: not enough word in file")
	}
	for i := 0; i < totalRandomWords; i++ {
		fmt.Printf("%s\n", wordArray[randInt(0, totalWords)])
	}
}

//read a space sepearted word file and return the array of words
func readFile(fileName string) (lines []string, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return lines, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	// scan all words from the file.
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
