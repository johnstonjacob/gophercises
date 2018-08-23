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

	m := storyhttp.MapPathsToArcs(S)

	if err != nil {
		panic(err)
	}

	tmpl, err := storyhttp.StoryTemplate(S)

	storyhttp.StoryHandler(tmpl, S, m)

	http.ListenAndServe(":8080", nil)

}
