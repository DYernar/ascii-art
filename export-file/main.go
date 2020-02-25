package main

import (
    "fmt"
    "html/template"
    "log"
	"net/http"
	"os"
	"bufio"
	"strings"
	"errors"
	
)

var GlobalAsciiArt string

type data struct {
	AsciiArt string
}

func PrintWord(w http.ResponseWriter, word string, asciiArt []string) (string, error) {
	str := ""
	for i := 1; i <= 8; i++ {
		for _, letter := range word{
			index := int(rune(letter)-32)*9
			if index+i >= len(asciiArt) || index+i < 0 {
				return "", errors.New("out of range")
			}
			str = str + asciiArt[index+i]
		}
		str = str + "\n"
	}
	str = str + "\n"

	return str, nil
}


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

func GetAscii(w http.ResponseWriter, input string, banner string ) (string, error) {
	var asciiArt []string

	file, err := os.Open(banner+".txt")

	if err != nil {
		fmt.Println("error opening file : " + err.Error())
		return "", err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
	asciiArt = append(asciiArt, line)
	}

	stringArr := strings.Split(input, "\\n")
	str := ""
	
	for _, word := range stringArr {
		newS, err := PrintWord(w, word, asciiArt)
		if err != nil {
			return "", err
		}
		str += newS
	}


	ConvertAndWriteToFile(input, asciiArt, "output")
	GlobalAsciiArt = str
	return str, nil
}

func getText(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == "POST" {
			w.WriteHeader(200)
			r.ParseForm()
			newStr := r.FormValue("text")
			finalStr := ""
			for index, symb := range newStr {
				if symb == 13 && newStr[index+1] == 10 {
					finalStr = finalStr + string(92) + string(110)
					continue
				}
				if symb==10 {
					continue
				}
				finalStr += string(symb)

			} 

			banner := r.FormValue("banner")
			log.Println(banner)
			str, err := GetAscii(w, finalStr,banner)
			if err != nil {
				w.WriteHeader(500)
				t, _ :=template.ParseFiles("error500.html")
				t.Execute(w, nil)
			} else {
				newData := data {
					AsciiArt: str,
				}
				fmt.Println(newData.AsciiArt)
				fmt.Print(str)

				t, _ := template.ParseFiles("index.html")
				t.Execute(w, newData)
			}
		} else if r.Method == "GET" {
			w.WriteHeader(200)

			w.Header().Set("Content-Type", "text/html")		
			t, _ := template.ParseFiles("index.html")
			t.Execute(w, nil)
		} else {
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "text/html")		
			t, _ :=template.ParseFiles("error400.html")
			t.Execute(w, nil)
		}
	} else {
		w.Header().Set("Content-Type", "text/html")		
		w.WriteHeader(404)
		t, _ := template.ParseFiles("error404.html")
		t.Execute(w, nil)
	}
}

func output(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, GlobalAsciiArt)
}

func main() {
    http.HandleFunc("/", getText)
    http.HandleFunc("/output.txt", output)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}