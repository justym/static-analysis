package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"os"
)

func main(){
	fset,err := parser.ParseExpr("(v + 1) + 1")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	ast.Inspect(fset,func(n ast.Node)bool{
		switch n := n.(type){
		case *ast.Ident:
			fmt.Println(n.Name)
		case *ast.BinaryExpr:
			fmt.Println(n.Op)
		case *ast.BasicLit:
			fmt.Println(n.Value)
		}
		return true
	})
}