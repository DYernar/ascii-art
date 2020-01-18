package main

import(
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"strings"
	"strconv"
)

func getLength(strToConvert string, symbolsArray []string) int {
	totalLength := 0
	for i:=1; i < 2;i++ {
		for _, letter := range strToConvert {
			index := int(rune(letter)-32)*9
			totalLength += len(symbolsArray[index+i])
		}
	}
	return totalLength
}


func PrintAscii(strToConvert string, symbolsArray []string, whiteSpace int) {
	for i := 0; i <= 8; i++ {
		printed := false
		for _, letter := range strToConvert {
			index := int(rune(letter)-32)*9
			printWhite := whiteSpace
			if !printed {
				for printWhite > 0 {
					fmt.Print(" ")
					printWhite--
				}
				printed = true
			}
			fmt.Print(symbolsArray[index+i])
		}
			fmt.Println()
	}
}

func AlignChoose(terminalWidth int, strLength int, align string, strToConvert string, symbolsArray []string) {
	if align == "center" {
		whiteSpace := (terminalWidth - strLength)/2
		PrintAscii(strToConvert, symbolsArray, whiteSpace)
	} else if align == "left" {
		PrintAscii(strToConvert, symbolsArray, 0)
	} else if align == "right" {
		whiteSpace := terminalWidth - strLength
		PrintAscii(strToConvert, symbolsArray, whiteSpace)
	} else if align == "justify"{
		strArr := strings.Split(strToConvert, " ")
		whiteSpace := ((terminalWidth - strLength)/len(strArr))/7
		var finalString string
		for index, item := range strArr {
			if index == 0 {
				finalString = item
				continue
			}
			if index == len(strArr) {
				finalString = finalString + item
				break	
			}
			for i:= 0; i < whiteSpace; i++ {
				finalString = finalString + " "
			}
			finalString = finalString + item

		}
		PrintAscii(finalString, symbolsArray, whiteSpace )
	}
}


func main() {
	strToConvert := os.Args[1]
	fontTypeFile := os.Args[2]
	var align string
	for index, item := range os.Args[3] {
		if index < 8 {
			continue
		}
		align += string(item)
	}
	var symbolsArray []string
	file, _ :=os.Open(fontTypeFile+".txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		symbolsArray = append(symbolsArray, scanner.Text())
	}

	strLength := getLength(strToConvert, symbolsArray)
	////----------------Terminal length--------------------////
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	termSize := ""
	for index, item := range out {
		if index == len(out)-1 {
			break
		}
		termSize += string(item)
	}
	hAndW := strings.Split(termSize, " ")
	width, _ := strconv.Atoi(hAndW[1])


	///////////////////////////////////////////////////////////
	AlignChoose(width, strLength, align,strToConvert, symbolsArray)

}