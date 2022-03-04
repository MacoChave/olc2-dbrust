package main

import (
	"html/template"
	"log"
	"net/http"
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

	// Original <=============Funciona
	// temp := template.Must(template.ParseFiles("template/index.html"))

	// http.HandleFunc("/analizar", func(rw http.ResponseWriter, r *http.Request) {
	// 	sourceCode := r.FormValue("source")
	// 	_ = sourceCode

	// 	// TODO: ANALIZAR ENTRADA
	// 	resultado := "Resultado"

	// 	temp.Execute(rw, struct {
	// 		Source string
	// 		Result string
	// 	}{
	// 		Source: sourceCode,
	// 		Result: resultado,
	// 	})
	// })

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	err := temp.Execute(rw, nil)
	// 	fmt.Println(err)
	// })

	log.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func index(rw http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(rw, "index", nil)
}

func analizar(rw http.ResponseWriter, r *http.Request) {
	sourceCode := r.FormValue("source")
	// TODO: Analize source code
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
