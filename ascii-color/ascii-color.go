package main

import(
	"bufio"
	"os"
	"strings"
	"fmt"
)

func rgbGenerator(color string) (string,string,string) {
	switch color{
		case "black": return "0","0","0"
		case "red": return "2","255","0"
		case "green": return "0","255","0"
		case "yellow": return "255","255","0"
		case "blue": return "0","0","255"
		case "magenta": return "255","0","255"
		case "cyan": return "0","255","255"
		case "orange": return "0","0","0"
	}	
	return "255","255","255 "
}

func PrintWord(word string, asciiArt []string, color string, slice string, startIndexToColor int) {
	// var printStr string
	for i := 1; i <= 8; i++ {
		for indexOfWord, letter := range word{
			index := int(rune(letter)-32)*9
			if startIndexToColor == -1 {
				if slice == "" {
					r,g,b := rgbGenerator(color)
					fmt.Printf("\033[38;2"+r+";"+g+";"+b+"m"+asciiArt[index+i])
				} else {
					fmt.Printf("\033[38;2;255;255;255m"+asciiArt[index+i])
				}

				continue			
			}

			if indexOfWord >= startIndexToColor && indexOfWord < len(slice)+startIndexToColor || startIndexToColor == 0 && slice==""{
				r,g,b := rgbGenerator(color)
				fmt.Printf("\033[38;2;"+r+";"+g+";"+b+"m"+asciiArt[index+i])
			} else {
				fmt.Printf("\033[38;2;255;255;255m"+asciiArt[index+i])
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