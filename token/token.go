package token
import (
    "text_to_tex/templates"
)
const (
    SEC ="SEC" 
    SUB ="SUB"
    PARAGRAPH ="PARAGRAPH"
    SEMICOLON ="SEMICOLON"
    SECNUM = "SECNUM"
    SUBNUM = "SUBNUM"
)
type Token struct {
    Type string 
    Literal string
}
type Tokens []rune
type Ast []Token
func nextChar(input []rune, cur int)( rune){
    if  cur + 1 >= len(input) {
        return rune(0)
    }
    return input[cur + 1]
}

func moveToken(cur int , input []rune)( int, string){
    if cur + 1 >= len(input){ // to do : generar error aqui #* 
        return len(input), ""
    }
    if input[cur + 1 ] == ';'{
        return cur + 1, ""
    }
    i:=0
    for i = cur  + 1; i < len(input) ; i++ {
        if  input[i] == ';'{
            break
        }
    }
    return i, string(input[cur + 1: i])
}    
func moveNext(cur int, input Tokens ) int {
    for i:= cur; i < len(input); i++  {
        if templates.IsSign(input[i]) { //#,_
            return i
        }
    }
    return len(input)
}
func Tokenize(input Tokens) ( Ast){
    cur := moveNext(0, input);
    var Statements Ast
    for cur < len(input) {
        curChar := input[cur]
        switch curChar  {
            case '#':
                nowChar := nextChar(input, cur) 
                if  nowChar == '*'{
                    cur +=1
                    curt, Lit := moveToken(cur, input)
                    cur = curt;
                    Statements = append( Statements,  Token{Type : SECNUM  , Literal : Lit}) 
                } else if  nowChar == '#'{
                    cur +=1
                    nowChar := nextChar(input, cur) 
                    if  nowChar == '*'{
                        cur +=1
                        curt, Lit := moveToken(cur, input)
                        cur = curt;
                        Statements = append( Statements,  Token{Type : SUBNUM  , Literal : Lit}) 
                    }else{
                        curt, Lit := moveToken(cur, input)
                        cur = curt;
                        Statements = append( Statements,  Token{Type : SUB, Literal : Lit}) 
                    }
                } else {
                    curt, Lit := moveToken(cur, input)
                    cur = curt;
                    Statements = append( Statements,  Token{Type : SEC, Literal : Lit}) 
                }
            case ';':
            cur = moveNext(cur, input);
            case '_':
                curt, Lit := moveToken(cur, input)
                cur = curt;
                Statements = append( Statements,  Token{Type : PARAGRAPH, Literal : Lit}) 
                cur = cur + 1
            default:
                cur +=1

        }
    }
    return Statements;
}

