package main

import (
	"db_rust/analizador/parser"
	"db_rust/utils"
	"html/template"
	"log"
	"net/http"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var templates = template.Must(template.ParseGlob("template/*"))

type Response struct {
	Source string
	Result string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/analize", analizar)
	// http.HandleFunc("/report", reports)
	http.HandleFunc("/about", about)

	log.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func index(rw http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(rw, "index", nil)
}

func analizar(rw http.ResponseWriter, r *http.Request) {
	sourceCode := r.FormValue("source")

	errorListener := &utils.CustomErrorListener{}
	inputStream := antlr.NewInputStream(sourceCode)

	// CREACION DEL LEXICO
	lexer := parser.Newlexico(inputStream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// CREACION DEL SINTACTICO
	syntax := parser.Newsintactico(stream)
	syntax.RemoveErrorListeners()
	syntax.AddErrorListener(errorListener)
	syntax.BuildParseTrees = true

	// GETTING ROOT
	// ast_root := syntax.Start()
	log.Printf("\nErrores encontrados %v\n", errorListener.Errors)

	// LISTENER TO AST
	// if len(errorListener.Errors) == 0 {
	// 	antlr.ParseTreeWalkerDefault.Walk()
	// }

	result := "Result here"

	res := Response{
		Source: sourceCode,
		Result: result,
	}
	templates.ExecuteTemplate(rw, "index", res)
}

// func reports(rw http.ResponseWriter, r *http.Request) {
// 	templates.ExecuteTemplate(rw, "index")
// }

func about(rw http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(rw, "about", nil)
}
