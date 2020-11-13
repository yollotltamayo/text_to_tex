# Text to Tex
This parser implementation parse my pseudo markdown to current LaTex, currently five commands are working, all comands should be appended at the end with
a semicolon:
- "#", to generate a section  -> "#Hello world;"  = "\section{Hello world}"
- "#*", to generate a no numbered section  -> "#*Hello world;"  = "\section*{Hello world}"
- "##", to generate a subsection  -> "##Hello world;"  = "\subsection{Hello world}"
- "##*", to generate a no numbered subsection  -> "##*Hello world;"  = "\subsection*{Hello world}"
- "_", to mark the beginnning of a empty paragraph "_Hello world;" = "Hello world"
I create this parser to write simple latex fastest and easilier, all aportations are welcome.
# Installation
Is written in Golang so it needs the latest version of GO to compile, then just clone this repo and do:
```
$ git clone "https://github.com/yollotltamayo/text_to_tex.git"
$ cd text_to_tex/
$ go run main.go "here goes you string to parse"
```
