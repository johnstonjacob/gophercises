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

	storyHandler, err := storyhttp.Storyhttp(&S)

	http.ListenAndServe(":8080", storyHandler)

}
