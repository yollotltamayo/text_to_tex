package main
import ( 
    "text_to_tex/parser"
    "fmt"
    "os"
)
func main() {
    data := os.Args
    if len(data) <2 {
        fmt.Printf("")
    }else{
        input := []rune(string(data[1]))
        fmt.Printf("%s",parser.Parse(input) )
    }
}



