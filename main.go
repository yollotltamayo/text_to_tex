package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text_to_tex/parser"
)

func main() {
	data := os.Args

	if len(data) >= 2 {

		file, err := ioutil.ReadFile(data[1])

		if err != nil {
			fmt.Printf("Error, can't open file.\n")
			return
		}

		output := []rune(string(file))
		fmt.Printf("%s", parser.Parse(output))
	} else {
		fmt.Printf("Error, source file not provided.\n")
	}
}
