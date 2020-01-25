package main

import(
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
)

func rgbGenerator(color string) (int,int,int) {
	switch color{
		case "black": return 0,0,0
		case "red": return 255,0,0
		case "green": return 0,128,0
		case "yellow": return 255,255,0
		case "blue": return 0,0,255
		case "magenta": return 255,0,255
		case "cyan": return 0,255,255
		case "lime": return 0,255,0
		case "silver": return 192,192,192
		case "gray": return 128,128,128
		case "maroon": return 128,0,0
		case "olive": return 128,128,0
		case "purple": return 128,0,128
		case "teal": return 0,128,128
		case "mint": return 170,255,195
		case "lavender": return 230,190,255
		case "pink": return 250,190,190
		case "brown": return 170, 110, 40
		case "orange": return 245, 130, 48
		case "apricot": return 255, 215, 180
		case "beige": return 255, 250, 200

	}	
	return 255,255,255
}

func PrintWord(word string, asciiArt []string, color string, slice string, startIndexToColor int) {
	// var printStr string
	for i := 1; i <= 8; i++ {
		for indexOfWord, letter := range word{
			index := int(rune(letter)-32)*9
			if startIndexToColor == -1 {
				if slice == "" {
					r,g,b := rgbGenerator(color)
					fmt.Printf("\033[38;2"+strconv.Itoa(r)+";"+strconv.Itoa(g)+";"+strconv.Itoa(b)+"m"+asciiArt[index+i])
				} else {
					fmt.Printf("\033[38;2;255;255;255m"+asciiArt[index+i])
				}
				continue			
			}

			if indexOfWord >= startIndexToColor && indexOfWord < len(slice)+startIndexToColor || startIndexToColor == 0 && slice==""{
				r,g,b := rgbGenerator(color)
				fmt.Printf("\033[38;2;"+strconv.Itoa(r)+";"+strconv.Itoa(g)+";"+strconv.Itoa(b)+"m"+asciiArt[index+i])
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