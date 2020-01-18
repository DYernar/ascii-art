package main

import(
	"os"
	"bufio"
)

func ConvertAndWriteToFile(strToConvert string, asciiSymbols []string, output string) {
	fileOutput, _ :=os.Create(output)
	writer := bufio.NewWriter(fileOutput)
	for i := 1; i<= 8; i++ {
		for _, letter := range strToConvert {
			index := int(rune(letter)-32) * 9
			writer.WriteString(asciiSymbols[index+i])
		}
		writer.WriteString("\n")
		writer.Flush() 
	}
}

func main() {
	fileName := os.Args[2]
	file, _ := os.Open(fileName+".txt")
	scanner := bufio.NewScanner(file)
	var asciiSymbols []string
	for scanner.Scan() {
		asciiSymbols = append(asciiSymbols, scanner.Text())
	}
	strToConvert := os.Args[1]

	var output string

	for index, letter := range os.Args[3] {
		if index < 9{
			continue
		}
		output += string(letter)
	}
	ConvertAndWriteToFile(strToConvert, asciiSymbols, output)

}