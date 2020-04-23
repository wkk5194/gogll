package token

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/goccmack/gogll/ast"
	"github.com/goccmack/gogll/cfg"
	"github.com/goccmack/gogll/im/tokens"
	"github.com/goccmack/goutil/ioutil"
)

type Data struct {
	Types        []*TypeDef
	TypeToString []string
}

type TypeDef struct {
	Name, Comment string
}

func Gen(g *ast.GoGLL, ts *tokens.Tokens) {
	tmpl, err := template.New("Token").Parse(tmplSrc)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, getData(ts))
	if err != nil {
		panic(err)
	}
	if err = ioutil.WriteFile(tokenFile(g.Package.GetString()), buf.Bytes()); err != nil {
		panic(err)
	}
}

func getData(ts *tokens.Tokens) *Data {
	return &Data{
		Types:        getTypes(ts),
		TypeToString: ts.TypeToString,
	}
}

func getTypes(ts *tokens.Tokens) (types []*TypeDef) {
	for i := range ts.TypeToString {
		types = append(types,
			&TypeDef{
				Name:    ts.TypeToString[i],
				Comment: ts.TypeToLiteral[i],
			})
	}
	return
}

// func GetTokenMap(g *ast.GoGLL) map[string]string {
// 	tokmap := map[string]string{
// 		"Error": "Error",
// 		"EOF":   "EOF",
// 	}
// 	for i, tok := range g.Terminals.ElementsSorted() {
// 		tokmap[tok] = fmt.Sprintf("Type%d", i)
// 	}
// 	return tokmap
// }

// func getSortedTokens(g *ast.GoGLL) (tokens []string) {
// 	for _, t := range g.Terminals.Elements() {
// 		tokens = append(tokens, t)
// 	}
// 	sort.Slice(tokens, func(i, j int) bool {
// 		return tokens[i] < tokens[j]
// 	})
// 	return
// }

func tokenFile(pkg string) string {
	return filepath.Join(cfg.BaseDir, "token", "token.go")
}

const tmplSrc = `
/*
Generated by GoGLL. Do not edit
*/

package token

import(
	"fmt"
)

// Token is returned by the lexer for every scanned lexical token
type Token struct {
	Type       Type
	Lext, Rext int
	Literal    []rune
}

// Type is the token type
type Type int
const({{range $i, $typ := .Types}}
	{{$typ.Name}} {{if eq $i 0}} Type = iota {{end}} // {{$typ.Comment}} {{end}}
)

var TypeToString = []string{ {{range $str := .TypeToString}}
	"{{$str}}",{{end}}
}

var StringToType = map[string] Type { {{range $typ := .TypeToString}}
	"{{$typ}}" : {{$typ}}, {{end}}
}

func New(t Type, lext, rext int, lit []rune) *Token {
	return &Token{
		Type: t,
		Lext: lext,
		Rext: rext,
		Literal: lit,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("%s (%d,%d) %s",
		TypeToString[t.Type], t.Lext, t.Rext, string(t.Literal))
}

func (t Type) String() string {
	return TypeToString[t]
}
`
