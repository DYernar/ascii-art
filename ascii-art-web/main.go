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

type data struct {
	AsciiArt string
}

func PrintWord(w http.ResponseWriter, word string, asciiArt []string) (string, error) {
	str := ""
	for i := 1; i <= 8; i++ {
		for _, letter := range word{
			index := int(rune(letter)-32)*9
			if index+i >= len(asciiArt) {
				return "", errors.New("out of range")
			}
			str = str + asciiArt[index+i]
		}
		str = str + "\n"
	}
	str = str + "\n"

	return str, nil
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
	return str, nil
}

func getText(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.WriteHeader(200)
		if r.Method == "POST" {
			r.ParseForm()
			newStr := ""
			fmt.Println(r.Form["text"])
			for _, str := range r.Form["text"] {
				newStr += str
			}
			banner := ""
			for _, str := range r.Form["banner"] {
				banner += str
			}
			if banner == "" {
				banner = "standard"
			}

			str, err := GetAscii(w, newStr,banner)
			if err != nil {
				t, _ :=template.ParseFiles("error500.html")
				t.Execute(w, nil)
			} else {
				fmt.Fprintf(w, str)
			}
		} else if r.Method == "GET" {
			w.Header().Set("Content-Type", "text/html")		
			t, _ := template.ParseFiles("index.html")
			t.Execute(w, nil)
		} else {
			w.Header().Set("Content-Type", "text/html")		
			w.WriteHeader(http.StatusNotImplemented)
			t, _ :=template.ParseFiles("error400.html")
			t.Execute(w, nil)
		}
	} else {
		w.Header().Set("Content-Type", "text/html")		
		w.WriteHeader(http.StatusNotFound)
		t, _ := template.ParseFiles("error404.html")
		t.Execute(w, nil)
	}
}

func main() {
    http.HandleFunc("/", getText)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}