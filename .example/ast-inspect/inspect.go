package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"os"
)

func main(){
	expr,err := parser.ParseExpr("fmt.Println(1 + 1 + 1)")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	ast.Inspect(expr,func (n ast.Node)bool{
		if n != nil{
			fmt.Printf("Pos: %d Type:%T\n",n.Pos(),n)
		}
		return true
	})
}

