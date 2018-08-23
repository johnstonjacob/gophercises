package main

import (
	"net/http"

	"github.com/johnstonjacob/gophercises/ex/cyoa/parsejson"
	"github.com/johnstonjacob/gophercises/ex/cyoa/storyhttp"
	// "html/template"
	// "http"
)

func main() {
	S, err := parsejson.ParseJSON("gopher.json")

	if err != nil {
		panic(err)
	}

	template, err := storyhttp.Storyhttp(S)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template.Execute(w, S["intro"])
	})

	http.ListenAndServe(":8080", nil)

}
