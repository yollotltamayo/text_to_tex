package parser

import (
	"text_to_tex/eval"
	"text_to_tex/token"
)

type Parser struct {
	Ast token.Ast
}

func Parse(input token.Tokens) string {
	l := token.Tokenize(input)
	val := eval.Eval(l)
	return val
}
