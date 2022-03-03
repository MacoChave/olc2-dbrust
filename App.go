package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	temp := template.Must(template.ParseFiles("template/index.html"))

	http.HandleFunc("/analizar", func(rw http.ResponseWriter, r *http.Request) {
		sourceCode := r.FormValue("source")
		_ = sourceCode

		// TODO: ANALIZAR ENTRADA
		resultado := "Resultado"

		temp.Execute(rw, struct {
			Source string
			Result string
		}{
			Source: sourceCode,
			Result: resultado,
		})
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		err := temp.Execute(rw, nil)
		fmt.Println(err)
	})

	http.ListenAndServe(":8080", nil)
}
