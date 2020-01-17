package main

import(
	"os"
	"fmt"
	"bufio"
	"strings"
)

func PrintWord(word string, asciiArt []string) {
	for i := 1; i <= 8; i++ {
		for _, letter := range word{
			index := int(rune(letter)-32)*9
			fmt.Print(asciiArt[index+i])
		}
		fmt.Println()
	}
}

func main() {
	var asciiArt []string

	file, err := os.Open("standard.txt")

	if err != nil {
		fmt.Println("error opening file : " + err.Error())
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
	asciiArt = append(asciiArt, line)
	}

	arguments := os.Args
	var input string
	for index, item := range arguments{
		if index == 0 {
			continue
		}
		if input == "" {
			input = item
			continue
		}
		input = input + " " + item
	}
	stringArr := strings.Split(input, "\\n")
	for _, word := range stringArr {
		PrintWord(word, asciiArt)
	}

}

							 