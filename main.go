package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

// create template container and load all templates into it
var tpl *template.Template

type book struct {
	Title	string
	Author	string
}

// create a FuncMap of functions to pass to the templates
// NOTE it is common to use very short function names, presumably because they
// often get pipelined together like {{.MyString | uc | ft}}
var funcMap = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

// instiate the templates and pass the funcmap up front
// NOTE passing the funcMap after like tpl.Funcs(funcMap) means that the
// template will be instantiated and be looking for any funcs used before they
// are passed so will not work
// this * creates a new template, * adds the func map to the result template
// * runs ParseGlob() on result template
func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	books := []book {
		{ Title: "Harry Potter", Author: "J K Rowling" },
		{ Title: "The Bible", Author: "JHVH" },
		{ Title: "Neuromancer", Author: "William Gibson"},
	}

	// execute to stdout
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", books)
	if err != nil {
		log.Fatalln(err)
	}
}
