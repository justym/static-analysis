package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

const src = `package main
func main(){
v := 100 // v is int value
fmt.Println(v+1)
}`

func main(){
	fset := token.NewFileSet()
	file,err := parser.ParseFile(fset,"main.go",src,parser.ParseComments)
	if err != nil{
		fmt.Errorf("%v",err)
		os.Exit(1)
	}
	cmap := ast.NewCommentMap(fset, file, file.Comments)
	ast.Inspect(file, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.AssignStmt:
			for _, cg := range cmap[n] {
				fmt.Println(cg.Text())
			}
		}
		return true
	})
}