package storyhttp

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/johnstonjacob/gophercises/ex/cyoa/parsejson"
)

// StoryTemplate is the http handler and html server
func StoryTemplate(s parsejson.Story) (*template.Template, error) {
	t, err := template.ParseFiles("template.html")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tmpl := template.Must(t, err)

	return tmpl, nil
}

// StoryHandler comment
func StoryHandler(tmpl *template.Template, S parsejson.Story, m pathsToArcs) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path, ok := m[r.URL.Path]
		if ok {
			tmpl.Execute(w, S[path])
		} else {
			tmpl.Execute(w, S["intro"])
		}
	})
}

// MapPathsToArcs comment
func MapPathsToArcs(s parsejson.Story) pathsToArcs {
	m := make(pathsToArcs)
	for k := range s {
		m["/"+k] = k
	}
	return m
}

type pathsToArcs map[string]string
