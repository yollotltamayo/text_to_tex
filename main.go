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
        if  data[1] == "-s"  {
            if len(data) >= 3 {
                file := data[2]
                output := []rune(string(file))
                fmt.Printf("%s", parser.Parse(output))
            }else{
                fmt.Printf("Error, no string to format.\n")

            }
        }else{
            file, err := ioutil.ReadFile(data[1])

            if err != nil {
                fmt.Printf("Error, can't open file.\n")
                return
            }else {
                    output := []rune(string(file))
                    fmt.Printf("%s", parser.Parse(output))
            }

        }
    }else{
        fmt.Print("Error, no source file, try with -s to parse strings.\n")
    }
}
