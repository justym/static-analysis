package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main(){
	fset := token.NewFileSet()
	expr,err := parser.ParseExprFrom(fset,"","v+1",0)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ast.Print(fset,expr)
}