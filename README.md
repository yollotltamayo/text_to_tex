# Text to Tex
This parser implementation parses pseudo markdown to current LaTeX, currently five commands are working, all commands should be appended at the end with
a semicolon:
- "#", to generate a section  -> "#Hello world;"  = "\section{Hello world}"
- "#*", to generate a no numbered section  -> "#*Hello world;"  = "\section*{Hello world}"
- "##", to generate a subsection  -> "##Hello world;"  = "\subsection{Hello world}"
- "##*", to generate a no numbered subsection  -> "##*Hello world;"  = "\subsection*{Hello world}"
- "_", to mark the beginning of a empty paragraph "_Hello world;" = "Hello world"

I created this parser to write simple latex faster and more easily, all contributions are welcome.


# Installation
It is written in Golang so it needs the latest version of GO to compile, then just clone this repo and do:
```
$ git clone "https://github.com/yollotltamayo/text_to_tex.git"
$ cd text_to_tex/
$ go run main.go {file to parse}
```
