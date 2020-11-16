package templates

var keywords = map[string]string{
	"SEC":       "\\section",
	"SUB":       "\\subsection",
	"PARAGRAPH": "",
	"SECNUM":    "\\section*",
	"SUBNUM":    "\\subsection*",
}

var signs = map[rune]bool{
	'#': true,
	'_': true,
}

func Key(a string) string {
	return keywords[a]
}

func IsSign(a rune) bool {
	return signs[a]
}
