package main

import (
    "fmt"
    "html/template"
    "log"
	"net/http"
	"os"
	"bufio"
	"strings"
	
)

func PrintWord(w http.ResponseWriter, word string, asciiArt []string) string {
	str := ""
	for i := 1; i <= 8; i++ {
		for _, letter := range word{
			index := int(rune(letter)-32)*9
			str = str + asciiArt[index+i]
		}
		str = str + "\n"
	}
	str = str + "\n"

	return str
}

func GetAscii(w http.ResponseWriter, input string ) string {
	var asciiArt []string

	file, err := os.Open("standard.txt")

	if err != nil {
		fmt.Println("error opening file : " + err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
	asciiArt = append(asciiArt, line)
	}

	stringArr := strings.Split(input, "\\n")
	str := ""
	for _, word := range stringArr {
		str += PrintWord(w, word, asciiArt)
	}
	return str
}

func getText(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
		r.ParseForm()
		
	if r.Method == "POST" {
		fmt.Println("text to convert:", r.Form["text"])
		newStr := ""
		for _, str := range r.Form["text"] {
			newStr += str
		}
		str := GetAscii(w, newStr)
		fmt.Fprintf(w, str)
	} else {
		t, _ := template.ParseFiles("index.gtpl")
		t.Execute(w, nil)
	}
}

func Form(w http.ResponseWriter, r *http.Request ) {
	t, _ := template.ParseFiles("index.gtpl")
	t.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", Form)
    http.HandleFunc("/getText", getText)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}