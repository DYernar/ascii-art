package main

import(
	"fmt"
	"flag"
)

func main() {
	s := flag.String("start", "start", "starts flag parse")
	f := flag.String("name", "YourName", "enter ur name")
	flag.Parse()
	arguments := flag.Args()


	fmt.Printf("flag start: %s\n", *s)
	fmt.Printf("flag name: %s\n", *f)
	fmt.Println(arguments)
}