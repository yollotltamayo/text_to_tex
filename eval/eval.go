package eval
 import (
     "text_to_tex/token"
     "text_to_tex/templates"
 )
type expValue interface {
    eval() string 
}
func Eval( ast token.Ast ) string  {
    ans := ""
    for _,a := range(ast ) {
        if a.Type == "PARAGRAPH"{
            if len(a.Literal )> 0 {
            ans += a.Literal +"\n"
        }
        }else { 
            if len(a.Literal )> 0 {
                ans += templates.Key(a.Type) +"{"+a.Literal +"}"+"\n"
            }
        }
    }
    return ans
}
