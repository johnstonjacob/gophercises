package main

import (
	"fmt"

	"github.com/johnstonjacob/gophercises/ex/cyoa/parsejson"
	// "html/template"
	// "http"
)

func main() {
	fmt.Println(parsejson.ParseJSON("gopher.json"))
	fmt.Println("vim-go")
}
