package storyhttp

import (
	"fmt"
	"html/template"

	"github.com/johnstonjacob/gophercises/ex/cyoa/parsejson"
)

// Storyhttp is the http handler and html server
func Storyhttp(s parsejson.Story) (*template.Template, error) {

	t, err := storyTemplate()

	if err != nil {
		return nil, err
	}

	return t, nil
}

func storyTemplate() (*template.Template, error) {
	t, err := template.ParseFiles("template.html")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tmpl := template.Must(t, err)

	return tmpl, nil
}
