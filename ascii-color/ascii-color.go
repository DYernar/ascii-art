package main

import(
	"bufio"
	"os"
	"strings"
	"fmt"
)

func rgbGenerator(color string) string {
	switch color{
		case "black": return "0"
		case "red": return "1"
		case "green": return "2"
		case "yellow": return "3"
		case "blue": return "4"
		case "magenta": return "5"
		case "cyan": return "6"
		case "orange": return "2"
	}	
	return "7"
}

func PrintWord(word string, asciiArt []string, color string, slice string, startIndexToColor int) {
	// var printStr string
	for i := 1; i <= 8; i++ {
		for indexOfWord, letter := range word{
			index := int(rune(letter)-32)*9
			if startIndexToColor == -1 {
				if slice == "" {
					genColor := rgbGenerator(color)
					fmt.Printf("\x1b[3"+ genColor +"m"+asciiArt[index+i])
				} else {
					fmt.Printf("\x1b[37m"+asciiArt[index+i])	
				}

				continue			
			}

			if indexOfWord >= startIndexToColor && indexOfWord < len(slice)+startIndexToColor || startIndexToColor == 0 && slice==""{
				genColor := rgbGenerator(color)
				fmt.Printf("\x1b[3"+ genColor +"m"+asciiArt[index+i])
			} else {
				fmt.Printf("\x1b[37m"+asciiArt[index+i])	
			}


		}
		fmt.Println()
	}

}


func main() {
	file, _ := os.Open("standard.txt")
	scanner := bufio.NewScanner(file)
	var asciiSymb []string
	for scanner.Scan() {
		asciiSymb = append(asciiSymb, scanner.Text())
	}
	var slice string
	if len(os.Args) > 3 {
		for index, item := range os.Args[3] {
			if index > 5 {
				slice += string(item)
			}
		}
	}
	strToConvert := os.Args[1]
	var color string
	for index, item := range os.Args[2] {
		if index >= 8 {
			color += string(item)
		}
	}
	startIndexToColor := strings.Index(strToConvert, slice)
	PrintWord(strToConvert, asciiSymb, color, slice, startIndexToColor)
}